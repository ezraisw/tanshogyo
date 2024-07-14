package usecaseimpl_test

import (
	"context"
	"errors"

	"github.com/ezraisw/tanshogyo/pkg/common/preseterrors"
	usecaseimpl "github.com/ezraisw/tanshogyo/services/product/internal/app/product/usecase/impl"
	usecasemock "github.com/ezraisw/tanshogyo/services/product/internal/app/product/usecase/mock"
	sellerusecase "github.com/ezraisw/tanshogyo/services/product/internal/app/seller/usecase"
	sellerusecasemock "github.com/ezraisw/tanshogyo/services/product/internal/app/seller/usecase/mock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProductAuthedDeleter", func() {
	var (
		mockSellerGetter     *sellerusecasemock.MockSellerGetter
		mockProductDeleter   *usecasemock.MockProductDeleter
		productAuthedDeleter *usecaseimpl.ProductAuthedDeleter
	)
	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		mockSellerGetter = sellerusecasemock.NewMockSellerGetter(ctrl)
		mockProductDeleter = usecasemock.NewMockProductDeleter(ctrl)
		productAuthedDeleter = usecaseimpl.NewProductAuthedDeleter(usecaseimpl.ProductAuthedDeleterOptions{
			SellerGetter:   mockSellerGetter,
			ProductDeleter: mockProductDeleter,
		})
	})

	Context("AuthedDelete", func() {
		var (
			userId string
			id     string
		)
		BeforeEach(func() {
			userId = dummyUserId
			id = dummyId
		})

		When("SellerGetter.GetByUserID returns an error", func() {
			var returnedErr error
			BeforeEach(func() {
				returnedErr = errors.New("mock error")
				mockSellerGetter.EXPECT().GetByUserID(gomock.Any(), gomock.Any()).Return(sellerusecase.Seller{}, returnedErr)
			})

			It("should return the error", func() {
				err := productAuthedDeleter.AuthedDelete(context.Background(), userId, id)
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
				err := productAuthedDeleter.AuthedDelete(context.Background(), userId, id)
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

			When("ProductDeleter.Delete returns an error", func() {
				var returnedErr error
				BeforeEach(func() {
					returnedErr = errors.New("mock error")
					mockProductDeleter.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(returnedErr)
				})

				It("should return the error", func() {
					err := productAuthedDeleter.AuthedDelete(context.Background(), userId, id)
					Expect(err).To(MatchError(returnedErr))
				})
			})

			When("ProductDeleter.Delete returns no error", func() {
				var deleteCall *gomock.Call
				BeforeEach(func() {
					deleteCall = mockProductDeleter.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)
				})

				JustBeforeEach(func() {
					By("passing correct arguments to Delete")

					deleteCall.Do(func(_ context.Context, argId string) {
						Expect(argId).To(Equal(id))
					})
				})

				It("should return the error", func() {
					err := productAuthedDeleter.AuthedDelete(context.Background(), userId, id)
					Expect(err).To(Succeed())
				})
			})
		})
	})
})
