package usecaseimpl_test

import (
	"context"
	"errors"

	repositorymock "github.com/ezraisw/tanshogyo/pkg/common/repository/mock"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/seller/model"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/seller/usecase"
	usecaseimpl "github.com/ezraisw/tanshogyo/services/product/internal/app/seller/usecase/impl"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SellerExistsChecker", func() {
	var (
		mockSellerRepository *repositorymock.MockRepository[model.Seller]
		sellerExistsChecker  *usecaseimpl.SellerExistsChecker
	)
	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		mockSellerRepository = repositorymock.NewMockRepository[model.Seller](ctrl)
		sellerExistsChecker = usecaseimpl.NewSellerExistsChecker(usecaseimpl.SellerExistsCheckerOptions{
			SellerRepository: mockSellerRepository,
		})
	})

	Context("CheckExists", func() {
		var fields []usecase.Field
		BeforeEach(func() {
			fields = dummyFields
		})

		When("SellerRepository.Exists returns an error", func() {
			var returnedErr error
			BeforeEach(func() {
				returnedErr = errors.New("mock error")
				mockSellerRepository.EXPECT().Exists(gomock.Any(), gomock.Any()).Return(false, returnedErr)
			})

			It("should return the error", func() {
				_, err := sellerExistsChecker.CheckExists(context.Background(), fields)
				Expect(err).To(MatchError(returnedErr))
			})
		})

		When("SellerRepository.Exists returns no error", func() {
			var returnedExists bool
			BeforeEach(func() {
				returnedExists = true
				mockSellerRepository.EXPECT().Exists(gomock.Any(), gomock.Any()).Return(returnedExists, nil)
			})

			It("should return exists and no error", func() {
				expectedExists := returnedExists

				exists, err := sellerExistsChecker.CheckExists(context.Background(), fields)
				Expect(err).To(Succeed())
				Expect(exists).To(Equal(expectedExists))
			})
		})
	})
})
