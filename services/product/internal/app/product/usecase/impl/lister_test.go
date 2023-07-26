package usecaseimpl_test

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pwnedgod/tanshogyo/pkg/common/entity"
	"github.com/pwnedgod/tanshogyo/pkg/common/repository"
	repositorymock "github.com/pwnedgod/tanshogyo/pkg/common/repository/mock"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/model"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/usecase"
	usecaseimpl "github.com/pwnedgod/tanshogyo/services/product/internal/app/product/usecase/impl"
)

var _ = Describe("ProductLister", func() {
	var (
		mockProductRepository *repositorymock.MockRepository[model.Product]
		productLister         *usecaseimpl.ProductLister
	)
	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		mockProductRepository = repositorymock.NewMockRepository[model.Product](ctrl)
		productLister = usecaseimpl.NewProductLister(usecaseimpl.ProductListerOptions{
			ProductRepository: mockProductRepository,
		})
	})

	Context("List", func() {
		var (
			limit  int
			offset int
		)
		BeforeEach(func() {
			limit = dummyLimit
			offset = dummyOffset
		})

		When("ProductRepository.Count returns an error", func() {
			var returnedErr error
			BeforeEach(func() {
				returnedErr = errors.New("mock error")
				mockProductRepository.EXPECT().Count(gomock.Any(), gomock.Any()).Return(0, returnedErr)
			})

			It("should return the error", func() {
				_, err := productLister.List(context.Background(), limit, offset)
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

					list, err := productLister.List(context.Background(), limit, offset)
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
						_, err := productLister.List(context.Background(), limit, offset)
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

						list, err := productLister.List(context.Background(), limit, offset)
						Expect(err).To(Succeed())
						Expect(list).To(Equal(expectedList))
					})
				})
			})
		})
	})
})
