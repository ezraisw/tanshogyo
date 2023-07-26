package usecaseimpl_test

import (
	"context"
	"errors"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	repositorymock "github.com/pwnedgod/tanshogyo/pkg/common/repository/mock"
	"github.com/pwnedgod/tanshogyo/pkg/common/util/helper"
	"github.com/pwnedgod/tanshogyo/pkg/common/util/timehelper"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/model"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/usecase"
	usecaseimpl "github.com/pwnedgod/tanshogyo/services/product/internal/app/product/usecase/impl"
	usecasemock "github.com/pwnedgod/tanshogyo/services/product/internal/app/product/usecase/mock"
)

var _ = Describe("ProductUpdater", func() {
	var (
		mockProductFormValidator *usecasemock.MockProductFormValidator
		mockProductRepository    *repositorymock.MockRepository[model.Product]
		productUpdater           *usecaseimpl.ProductUpdater

		now time.Time
	)
	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		mockProductFormValidator = usecasemock.NewMockProductFormValidator(ctrl)
		mockProductRepository = repositorymock.NewMockRepository[model.Product](ctrl)
		productUpdater = usecaseimpl.NewProductUpdater(usecaseimpl.ProductUpdaterOptions{
			ProductFormValidator: mockProductFormValidator,
			ProductRepository:    mockProductRepository,
			Nower:                timehelper.NowerFunc(func() time.Time { return now }),
		})

		now = dummyNow
	})

	Context("Update", func() {
		var (
			id   string
			form usecase.ProductForm
		)
		BeforeEach(func() {
			id = dummyId
			form = dummyForm
		})

		When("ProductFormValidator.Validate returns an error", func() {
			var returnedErr error
			BeforeEach(func() {
				returnedErr = errors.New("mock error")
				mockProductFormValidator.EXPECT().Validate(gomock.Any(), gomock.Any()).Return(returnedErr)
			})

			It("should return the error", func() {
				_, err := productUpdater.Update(context.Background(), id, form)
				Expect(err).To(MatchError(returnedErr))
			})
		})

		When("ProductFormValidator.Validate returns no error", func() {
			var validateCall *gomock.Call
			BeforeEach(func() {
				validateCall = mockProductFormValidator.EXPECT().Validate(gomock.Any(), gomock.Any()).Return(nil)
			})

			JustBeforeEach(func() {
				By("passing correct arguments to Validate")

				validateCall.Do(func(_ context.Context, argForm usecase.ProductForm) {
					Expect(argForm).To(Equal(form))
				})
			})

			When("ProductRepository.Find returns an error", func() {
				var returnedErr error
				BeforeEach(func() {
					returnedErr = errors.New("mock error")
					mockProductRepository.EXPECT().Find(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, returnedErr)
				})

				It("should return the error", func() {
					_, err := productUpdater.Update(context.Background(), id, form)
					Expect(err).To(MatchError(returnedErr))
				})
			})

			When("ProductRepository.Find returns no error", func() {
				var returnedProduct *model.Product
				BeforeEach(func() {
					now = dummyNow
					earlier := dummyNow.Add(-1 * time.Hour)

					returnedProduct = helper.Ptr(dummyProduct)
					returnedProduct.CreatedAt = earlier.UTC()
					returnedProduct.UpdatedAt = earlier.UTC()
					mockProductRepository.EXPECT().Find(gomock.Any(), gomock.Any(), gomock.Any()).Return(returnedProduct, nil)
				})

				When("ProductRepository.Update returns an error", func() {
					var returnedErr error
					BeforeEach(func() {
						returnedErr = errors.New("mock error")
						mockProductRepository.EXPECT().Update(gomock.Any(), gomock.Any()).Return(returnedErr)
					})

					It("should return the error", func() {
						_, err := productUpdater.Update(context.Background(), id, form)
						Expect(err).To(MatchError(returnedErr))
					})
				})

				When("ProductRepository.Update returns no error", func() {
					var (
						expectedProduct *model.Product
						updateCall      *gomock.Call
					)
					BeforeEach(func() {
						form = usecase.ProductForm{
							SellerID:    "bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb",
							Name:        "Forbidden Item",
							Description: "Sacred",
							Price:       1000000,
							Quantity:    1,
						}
						expectedProduct = &model.Product{
							ID:          id,
							SellerID:    form.SellerID,
							Name:        form.Name,
							Description: form.Description,
							Price:       form.Price,
							Quantity:    form.Quantity,
							CreatedAt:   returnedProduct.CreatedAt,
							UpdatedAt:   now.UTC(),
						}
						updateCall = mockProductRepository.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
					})

					JustBeforeEach(func() {
						By("passing correct arguments to Update")

						updateCall.Do(func(_ context.Context, product *model.Product) {
							Expect(product).To(Equal(expectedProduct))
						})
					})

					It("should return correct dto", func() {
						expectedDto := usecase.Product{
							ID:          expectedProduct.ID,
							SellerID:    expectedProduct.SellerID,
							Name:        expectedProduct.Name,
							Description: expectedProduct.Description,
							Price:       expectedProduct.Price,
							Quantity:    expectedProduct.Quantity,
							CreatedAt:   expectedProduct.CreatedAt,
							UpdatedAt:   expectedProduct.UpdatedAt,
						}

						dto, err := productUpdater.Update(context.Background(), id, form)
						Expect(err).To(Succeed())
						Expect(dto).To(Equal(expectedDto))
					})
				})
			})
		})
	})
})
