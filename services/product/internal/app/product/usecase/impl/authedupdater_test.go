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

var _ = Describe("ProductAuthedUpdater", func() {
	var (
		mockSellerGetter     *sellerusecasemock.MockSellerGetter
		mockProductUpdater   *usecasemock.MockProductUpdater
		productAuthedUpdater *usecaseimpl.ProductAuthedUpdater
	)
	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		mockSellerGetter = sellerusecasemock.NewMockSellerGetter(ctrl)
		mockProductUpdater = usecasemock.NewMockProductUpdater(ctrl)
		productAuthedUpdater = usecaseimpl.NewProductAuthedUpdater(usecaseimpl.ProductAuthedUpdaterOptions{
			SellerGetter:   mockSellerGetter,
			ProductUpdater: mockProductUpdater,
		})
	})

	Context("AuthedUpdate", func() {
		var (
			userId string
			id     string
			form   usecase.AuthedProductForm
		)
		BeforeEach(func() {
			userId = dummyUserId
			id = dummyId
			form = dummyAuthedForm
		})

		When("SellerGetter.GetByUserID returns an error", func() {
			var returnedErr error
			BeforeEach(func() {
				returnedErr = errors.New("mock error")
				mockSellerGetter.EXPECT().GetByUserID(gomock.Any(), gomock.Any()).Return(sellerusecase.Seller{}, returnedErr)
			})

			It("should return the error", func() {
				_, err := productAuthedUpdater.AuthedUpdate(context.Background(), userId, id, form)
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
				_, err := productAuthedUpdater.AuthedUpdate(context.Background(), userId, id, form)
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

			When("ProductUpdater.Update returns an error", func() {
				var returnedErr error
				BeforeEach(func() {
					returnedErr = errors.New("mock error")
					mockProductUpdater.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any()).Return(usecase.Product{}, returnedErr)
				})

				It("should return the error", func() {
					_, err := productAuthedUpdater.AuthedUpdate(context.Background(), userId, id, form)
					Expect(err).To(MatchError(returnedErr))
				})
			})

			When("ProductUpdater.Update returns no error", func() {
				var (
					returnedProduct usecase.Product
					updateCall      *gomock.Call
				)
				BeforeEach(func() {
					returnedProduct = dummyProductDto
					updateCall = mockProductUpdater.EXPECT().Update(gomock.Any(), gomock.Any(), gomock.Any()).Return(returnedProduct, nil)
				})

				JustBeforeEach(func() {
					By("passing correct arguments to Update")

					updateCall.Do(func(_ context.Context, argId string, argForm usecase.ProductForm) {
						expectedForm := usecase.ProductForm{
							SellerID:    returnedSeller.ID,
							Name:        form.Name,
							Description: form.Description,
							Price:       form.Price,
							Quantity:    form.Quantity,
						}
						Expect(argId).To(Equal(id))
						Expect(argForm).To(Equal(expectedForm))
					})
				})

				It("should return the error", func() {
					product, err := productAuthedUpdater.AuthedUpdate(context.Background(), userId, id, form)
					Expect(err).To(Succeed())
					Expect(product).To(Equal(returnedProduct))
				})
			})
		})
	})
})
