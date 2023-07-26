package usecaseimpl_test

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pwnedgod/tanshogyo/pkg/common/preseterrors"
	repositorymock "github.com/pwnedgod/tanshogyo/pkg/common/repository/mock"
	"github.com/pwnedgod/tanshogyo/pkg/common/util/helper"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/model"
	usecaseimpl "github.com/pwnedgod/tanshogyo/services/product/internal/app/product/usecase/impl"
	sellerusecase "github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/usecase"
	sellerusecasemock "github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/usecase/mock"
)

var _ = Describe("ProductAuthedCreator", func() {
	var (
		mockSellerGetter      *sellerusecasemock.MockSellerGetter
		mockProductRepository *repositorymock.MockRepository[model.Product]
		productSellerChecker  *usecaseimpl.ProductSellerChecker
	)
	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		mockSellerGetter = sellerusecasemock.NewMockSellerGetter(ctrl)
		mockProductRepository = repositorymock.NewMockRepository[model.Product](ctrl)
		productSellerChecker = usecaseimpl.NewProductSellerChecker(usecaseimpl.ProductSellerCheckerOptions{
			SellerGetter:      mockSellerGetter,
			ProductRepository: mockProductRepository,
		})
	})

	Context("CheckSeller", func() {
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
				err := productSellerChecker.CheckSeller(context.Background(), userId, id)
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
				err := productSellerChecker.CheckSeller(context.Background(), userId, id)
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

			When("ProductRepository.Find returns an error", func() {
				var returnedErr error
				BeforeEach(func() {
					returnedErr = errors.New("mock error")
					mockProductRepository.EXPECT().Find(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, returnedErr)
				})

				It("should return the error", func() {
					err := productSellerChecker.CheckSeller(context.Background(), userId, id)
					Expect(err).To(MatchError(returnedErr))
				})
			})

			When("ProductRepository.Find returns product with different seller id and no error", func() {
				var (
					returnedProduct *model.Product
				)
				BeforeEach(func() {
					returnedProduct = helper.Ptr(dummyProduct)
					returnedProduct.SellerID = returnedSeller.ID + "different"
					mockProductRepository.EXPECT().Find(gomock.Any(), gomock.Any(), gomock.Any()).Return(returnedProduct, nil)
				})

				It("should return forbidden error", func() {
					err := productSellerChecker.CheckSeller(context.Background(), userId, id)
					Expect(err).To(MatchError(preseterrors.ErrForbidden))
				})
			})

			When("ProductRepository.Find returns no error", func() {
				var (
					returnedProduct *model.Product
				)
				BeforeEach(func() {
					returnedProduct = helper.Ptr(dummyProduct)
					mockProductRepository.EXPECT().Find(gomock.Any(), gomock.Any(), gomock.Any()).Return(returnedProduct, nil)
				})

				It("should return no error", func() {
					err := productSellerChecker.CheckSeller(context.Background(), userId, id)
					Expect(err).To(Succeed())
				})
			})
		})
	})
})
