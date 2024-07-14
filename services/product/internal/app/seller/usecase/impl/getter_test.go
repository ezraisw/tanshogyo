package usecaseimpl_test

import (
	"context"
	"errors"

	repositorymock "github.com/ezraisw/tanshogyo/pkg/common/repository/mock"
	"github.com/ezraisw/tanshogyo/pkg/common/util/helper"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/seller/model"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/seller/usecase"
	usecaseimpl "github.com/ezraisw/tanshogyo/services/product/internal/app/seller/usecase/impl"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SellerGetter", func() {
	var (
		mockSellerRepository *repositorymock.MockRepository[model.Seller]
		productGetter        *usecaseimpl.SellerGetter
	)
	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		mockSellerRepository = repositorymock.NewMockRepository[model.Seller](ctrl)
		productGetter = usecaseimpl.NewSellerGetter(usecaseimpl.SellerGetterOptions{
			SellerRepository: mockSellerRepository,
		})
	})

	Context("GetByUserID", func() {
		var userId string
		BeforeEach(func() {
			userId = dummyUserId
		})

		When("SellerRepository.Find returns an error", func() {
			var returnedErr error
			BeforeEach(func() {
				returnedErr = errors.New("mock error")
				mockSellerRepository.EXPECT().Find(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, returnedErr)
			})

			It("should return the error", func() {
				_, err := productGetter.GetByUserID(context.Background(), userId)
				Expect(err).To(MatchError(returnedErr))
			})
		})

		When("SellerRepository.Find returns no error", func() {
			var returnedSeller *model.Seller
			BeforeEach(func() {
				returnedSeller = helper.Ptr(dummySeller)
				mockSellerRepository.EXPECT().Find(gomock.Any(), gomock.Any(), gomock.Any()).Return(returnedSeller, nil)
			})

			It("should return correct dto", func() {
				expectedDto := usecase.Seller{
					ID:        returnedSeller.ID,
					UserID:    returnedSeller.UserID,
					ShopName:  returnedSeller.ShopName,
					CreatedAt: returnedSeller.CreatedAt,
					UpdatedAt: returnedSeller.UpdatedAt,
				}

				dto, err := productGetter.GetByUserID(context.Background(), userId)
				Expect(err).To(Succeed())
				Expect(dto).To(Equal(expectedDto))
			})
		})
	})
})
