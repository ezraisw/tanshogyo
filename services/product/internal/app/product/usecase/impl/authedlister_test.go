package usecaseimpl_test

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pwnedgod/tanshogyo/pkg/common/entity"
	"github.com/pwnedgod/tanshogyo/pkg/common/preseterrors"
	"github.com/pwnedgod/tanshogyo/pkg/common/repository"
	repositorymock "github.com/pwnedgod/tanshogyo/pkg/common/repository/mock"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/model"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/usecase"
	usecaseimpl "github.com/pwnedgod/tanshogyo/services/product/internal/app/product/usecase/impl"
	sellerusecase "github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/usecase"
	sellerusecasemock "github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/usecase/mock"
)

var _ = Describe("ProductAuthedLister", func() {
	var (
		mockSellerGetter      *sellerusecasemock.MockSellerGetter
		mockProductRepository *repositorymock.MockRepository[model.Product]
		productAuthedLister   *usecaseimpl.ProductAuthedLister
	)
	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		mockSellerGetter = sellerusecasemock.NewMockSellerGetter(ctrl)
		mockProductRepository = repositorymock.NewMockRepository[model.Product](ctrl)
		productAuthedLister = usecaseimpl.NewProductAuthedLister(usecaseimpl.ProductAuthedListerOptions{
			SellerGetter:      mockSellerGetter,
			ProductRepository: mockProductRepository,
		})
	})

	Context("List", func() {
		var (
			userId string
			limit  int
			offset int
		)
		BeforeEach(func() {
			userId = dummyUserId
			limit = dummyLimit
			offset = dummyOffset
		})

		When("SellerGetter.GetByUserID returns an error", func() {
			var returnedErr error
			BeforeEach(func() {
				returnedErr = errors.New("mock error")
				mockSellerGetter.EXPECT().GetByUserID(gomock.Any(), gomock.Any()).Return(sellerusecase.Seller{}, returnedErr)
			})

			It("should return the error", func() {
				_, err := productAuthedLister.AuthedList(context.Background(), userId, limit, offset)
				Expect(err).To(MatchError(returnedErr))
			})
		})

		When("SellerGetter.GetByUserID returns not found error", func() {
			var returnedErr error
			BeforeEach(func() {
				returnedErr = preseterrors.ErrNotFound
				mockSellerGetter.EXPECT().GetByUserID(gomock.Any(), gomock.Any()).Return(sellerusecase.Seller{}, returnedErr)
			})

			It("should return the error", func() {
				_, err := productAuthedLister.AuthedList(context.Background(), userId, limit, offset)
				Expect(err).To(MatchError(preseterrors.ErrUnauthorized))
			})
		})

		When("SellerGetter.GetByUserID returns no error", func() {
			BeforeEach(func() {
				mockSellerGetter.EXPECT().GetByUserID(gomock.Any(), gomock.Any()).Return(dummySeller, nil)
			})

			When("ProductRepository.Count returns an error", func() {
				var returnedErr error
				BeforeEach(func() {
					returnedErr = errors.New("mock error")
					mockProductRepository.EXPECT().Count(gomock.Any(), gomock.Any()).Return(0, returnedErr)
				})

				It("should return the error", func() {
					_, err := productAuthedLister.AuthedList(context.Background(), userId, limit, offset)
					Expect(err).To(MatchError(returnedErr))
				})
			})

			When("ProductRepository.Count returns no error", func() {
				BeforeEach(func() {
					mockProductRepository.EXPECT().Count(gomock.Any(), gomock.Any()).Return(len(dummyProducts), nil)
				})

				When("count is less than offset", func() {
					BeforeEach(func() {
						offset = len(dummyProducts) + 2
					})

					It("should return empty data with no error", func() {
						expectedList := usecase.ProductList{
							Count:  len(dummyProducts),
							Limit:  limit,
							Offset: offset,
							Data:   []usecase.Product{},
						}

						list, err := productAuthedLister.AuthedList(context.Background(), userId, limit, offset)
						Expect(err).To(Succeed())
						Expect(list).To(Equal(expectedList))
					})
				})

				When("count is more than offset", func() {
					When("ProductRepository.FindMany returns an error", func() {
						var returnedErr error
						BeforeEach(func() {
							returnedErr = errors.New("mock error")
							mockProductRepository.EXPECT().FindMany(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, returnedErr)
						})

						It("should return the error", func() {
							_, err := productAuthedLister.AuthedList(context.Background(), userId, limit, offset)
							Expect(err).To(MatchError(returnedErr))
						})
					})

					When("ProductRepository.FindMany returns no error", func() {
						var (
							findManyCall *gomock.Call
						)
						BeforeEach(func() {
							findManyCall = mockProductRepository.EXPECT().FindMany(gomock.Any(), gomock.Any(), gomock.Any()).Return(dummyProducts, nil)
						})

						JustBeforeEach(func() {
							By("passing correct arguments to FindMany")

							findManyCall.Do(func(_ context.Context, _ entity.Clause, options repository.FindManyOptions) {
								expectedOptions := repository.FindManyOptions{
									Limit:  limit,
									Offset: offset,
								}
								Expect(options).To(Equal(expectedOptions))
							})
						})

						It("should return list and no error", func() {
							expectedList := usecase.ProductList{
								Count:  len(dummyProducts),
								Limit:  limit,
								Offset: offset,
								Data:   []usecase.Product{},
							}

							for _, product := range dummyProducts {
								expectedList.Data = append(expectedList.Data, usecase.Product{
									ID:          product.ID,
									SellerID:    product.SellerID,
									Name:        product.Name,
									Description: product.Description,
									Price:       product.Price,
									Quantity:    product.Quantity,
									CreatedAt:   product.CreatedAt,
									UpdatedAt:   product.UpdatedAt,
								})
							}

							list, err := productAuthedLister.AuthedList(context.Background(), userId, limit, offset)
							Expect(err).To(Succeed())
							Expect(list).To(Equal(expectedList))
						})
					})
				})
			})
		})
	})
})
