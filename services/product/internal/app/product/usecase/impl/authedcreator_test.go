package usecaseimpl_test

import (
	"context"
	"errors"

	"github.com/ezraisw/tanshogyo/pkg/common/preseterrors"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/product/usecase"
	usecaseimpl "github.com/ezraisw/tanshogyo/services/product/internal/app/product/usecase/impl"
	usecasemock "github.com/ezraisw/tanshogyo/services/product/internal/app/product/usecase/mock"
	sellerusecase "github.com/ezraisw/tanshogyo/services/product/internal/app/seller/usecase"
	sellerusecasemock "github.com/ezraisw/tanshogyo/services/product/internal/app/seller/usecase/mock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProductAuthedCreator", func() {
	var (
		mockSellerGetter     *sellerusecasemock.MockSellerGetter
		mockProductCreator   *usecasemock.MockProductCreator
		productAuthedCreator *usecaseimpl.ProductAuthedCreator
	)
	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		mockSellerGetter = sellerusecasemock.NewMockSellerGetter(ctrl)
		mockProductCreator = usecasemock.NewMockProductCreator(ctrl)
		productAuthedCreator = usecaseimpl.NewProductAuthedCreator(usecaseimpl.ProductAuthedCreatorOptions{
			SellerGetter:   mockSellerGetter,
			ProductCreator: mockProductCreator,
		})
	})

	Context("AuthedCreate", func() {
		var (
			userId string
			form   usecase.AuthedProductForm
		)
		BeforeEach(func() {
			userId = dummyUserId
			form = dummyAuthedForm
		})

		When("SellerGetter.GetByUserID returns an error", func() {
			var returnedErr error
			BeforeEach(func() {
				returnedErr = errors.New("mock error")
				mockSellerGetter.EXPECT().GetByUserID(gomock.Any(), gomock.Any()).Return(sellerusecase.Seller{}, returnedErr)
			})

			It("should return the error", func() {
				_, err := productAuthedCreator.AuthedCreate(context.Background(), userId, form)
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
				_, err := productAuthedCreator.AuthedCreate(context.Background(), userId, form)
				Expect(err).To(MatchError(preseterrors.ErrUnauthorized))
			})
		})

		When("SellerGetter.GetByUserID returns no error", func() {
			var (
				returnedSeller  sellerusecase.Seller
				getByUserIdCall *gomock.Call
			)
			BeforeEach(func() {
				returnedSeller = dummySeller
				getByUserIdCall = mockSellerGetter.EXPECT().GetByUserID(gomock.Any(), gomock.Any()).Return(returnedSeller, nil)
			})

			JustBeforeEach(func() {
				By("passing correct arguments to GetByUserID")

				getByUserIdCall.Do(func(_ context.Context, argUserId string) {
					Expect(argUserId).To(Equal(userId))
				})
			})

			When("ProductCreator.Create returns an error", func() {
				var returnedErr error
				BeforeEach(func() {
					returnedErr = errors.New("mock error")
					mockProductCreator.EXPECT().Create(gomock.Any(), gomock.Any()).Return(usecase.Product{}, returnedErr)
				})

				It("should return the error", func() {
					_, err := productAuthedCreator.AuthedCreate(context.Background(), userId, form)
					Expect(err).To(MatchError(returnedErr))
				})
			})

			When("ProductCreator.Create returns no error", func() {
				var (
					returnedProduct usecase.Product
					createCall      *gomock.Call
				)
				BeforeEach(func() {
					returnedProduct = dummyProductDto
					createCall = mockProductCreator.EXPECT().Create(gomock.Any(), gomock.Any()).Return(returnedProduct, nil)
				})

				JustBeforeEach(func() {
					By("passing correct arguments to Create")

					createCall.Do(func(_ context.Context, argForm usecase.ProductForm) {
						expectedForm := usecase.ProductForm{
							SellerID:    returnedSeller.ID,
							Name:        form.Name,
							Description: form.Description,
							Price:       form.Price,
							Quantity:    form.Quantity,
						}
						Expect(argForm).To(Equal(expectedForm))
					})
				})

				It("should return the error", func() {
					product, err := productAuthedCreator.AuthedCreate(context.Background(), userId, form)
					Expect(err).To(Succeed())
					Expect(product).To(Equal(returnedProduct))
				})
			})
		})
	})
})
