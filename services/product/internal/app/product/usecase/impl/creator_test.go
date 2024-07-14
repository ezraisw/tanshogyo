package usecaseimpl_test

import (
	"context"
	"errors"
	"time"

	repositorymock "github.com/ezraisw/tanshogyo/pkg/common/repository/mock"
	"github.com/ezraisw/tanshogyo/pkg/common/util/helper"
	"github.com/ezraisw/tanshogyo/pkg/common/util/idgen"
	"github.com/ezraisw/tanshogyo/pkg/common/util/timehelper"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/product/model"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/product/usecase"
	usecaseimpl "github.com/ezraisw/tanshogyo/services/product/internal/app/product/usecase/impl"
	usecasemock "github.com/ezraisw/tanshogyo/services/product/internal/app/product/usecase/mock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProductCreator", func() {
	var (
		mockProductFormValidator *usecasemock.MockProductFormValidator
		mockProductRepository    *repositorymock.MockRepository[model.Product]
		productCreator           *usecaseimpl.ProductCreator

		now time.Time
		id  string
	)
	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		mockProductFormValidator = usecasemock.NewMockProductFormValidator(ctrl)
		mockProductRepository = repositorymock.NewMockRepository[model.Product](ctrl)
		productCreator = usecaseimpl.NewProductCreator(usecaseimpl.ProductCreatorOptions{
			ProductFormValidator: mockProductFormValidator,
			ProductRepository:    mockProductRepository,
			Nower:                timehelper.NowerFunc(func() time.Time { return now }),
			IDGen:                idgen.IDGenFunc(func() string { return id }),
		})

		now = dummyNow
		id = dummyId
	})

	Context("Create", func() {
		var form usecase.ProductForm
		BeforeEach(func() {
			form = dummyForm
		})

		When("ProductFormValidator.Validate returns an error", func() {
			var returnedErr error
			BeforeEach(func() {
				returnedErr = errors.New("mock error")
				mockProductFormValidator.EXPECT().Validate(gomock.Any(), gomock.Any()).Return(returnedErr)
			})

			It("should return the error", func() {
				_, err := productCreator.Create(context.Background(), form)
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

			When("ProductRepository.Create returns an error", func() {
				var returnedErr error
				BeforeEach(func() {
					returnedErr = errors.New("mock error")
					mockProductRepository.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, returnedErr)
				})

				It("should return the error", func() {
					_, err := productCreator.Create(context.Background(), form)
					Expect(err).To(MatchError(returnedErr))
				})
			})

			When("ProductRepository.Create returns no error", func() {
				var (
					expectedProduct *model.Product
					returnedProduct *model.Product
					createCall      *gomock.Call
				)
				BeforeEach(func() {
					expectedProduct = helper.Ptr(dummyProduct)
					returnedProduct = expectedProduct
					createCall = mockProductRepository.EXPECT().Create(gomock.Any(), gomock.Any()).Return(returnedProduct, nil)
				})

				JustBeforeEach(func() {
					By("passing correct arguments to Create")

					createCall.Do(func(_ context.Context, product *model.Product) {
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

					dto, err := productCreator.Create(context.Background(), form)
					Expect(err).To(Succeed())
					Expect(dto).To(Equal(expectedDto))
				})
			})
		})
	})
})
