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
	"github.com/pwnedgod/tanshogyo/pkg/common/util/idgen"
	"github.com/pwnedgod/tanshogyo/pkg/common/util/timehelper"
	sellererrors "github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/errors"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/model"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/usecase"
	usecaseimpl "github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/usecase/impl"
	usecasemock "github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/usecase/mock"
)

var _ = Describe("SellerRegisterer", func() {
	var (
		mockSellerFormValidator *usecasemock.MockSellerFormValidator
		mockSellerRepository    *repositorymock.MockRepository[model.Seller]
		sellerRegisterer        *usecaseimpl.SellerRegisterer

		now time.Time
		id  string
	)
	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		mockSellerFormValidator = usecasemock.NewMockSellerFormValidator(ctrl)
		mockSellerRepository = repositorymock.NewMockRepository[model.Seller](ctrl)
		sellerRegisterer = usecaseimpl.NewSellerRegisterer(usecaseimpl.SellerRegistererOptions{
			SellerFormValidator: mockSellerFormValidator,
			SellerRepository:    mockSellerRepository,
			Nower:               timehelper.NowerFunc(func() time.Time { return now }),
			IDGen:               idgen.IDGenFunc(func() string { return id }),
		})

		now = dummyNow
		id = dummyId

	})

	Context("Register", func() {
		var (
			userId string
			form   usecase.SellerForm
		)
		BeforeEach(func() {
			userId = dummyUserId
			form = dummyForm
		})

		When("SellerFormValidator.Validate returns an error", func() {
			var returnedErr error
			BeforeEach(func() {
				returnedErr = errors.New("mock error")
				mockSellerFormValidator.EXPECT().Validate(gomock.Any(), gomock.Any()).Return(returnedErr)
			})

			It("should return the error", func() {
				err := sellerRegisterer.Register(context.Background(), userId, form)
				Expect(err).To(MatchError(returnedErr))
			})
		})

		When("SellerFormValidator.Validate returns no error", func() {
			var validateCall *gomock.Call
			BeforeEach(func() {
				validateCall = mockSellerFormValidator.EXPECT().Validate(gomock.Any(), gomock.Any()).Return(nil)
			})

			JustBeforeEach(func() {
				validateCall.Do(func(_ context.Context, argForm usecase.SellerForm) {
					Expect(argForm).To(Equal(form))
				})
			})

			When("SellerRepository.Exusts returns an error", func() {
				var returnedErr error
				BeforeEach(func() {
					returnedErr = errors.New("mock error")
					mockSellerRepository.EXPECT().Exists(gomock.Any(), gomock.Any()).Return(false, returnedErr)
				})

				It("should return the error", func() {
					err := sellerRegisterer.Register(context.Background(), userId, form)
					Expect(err).To(MatchError(returnedErr))
				})
			})

			When("SellerRepository.Exists returns true and no error", func() {
				BeforeEach(func() {
					mockSellerRepository.EXPECT().Exists(gomock.Any(), gomock.Any()).Return(true, nil)
				})

				It("should return already has seller account error", func() {
					err := sellerRegisterer.Register(context.Background(), userId, form)
					Expect(err).To(MatchError(sellererrors.ErrAlreadyHasSellerAccount))
				})
			})

			When("SellerRepostiryo.Exists returns false and no error", func() {
				BeforeEach(func() {
					mockSellerRepository.EXPECT().Exists(gomock.Any(), gomock.Any()).Return(false, nil)
				})

				When("SellerRepository.Create returns an error", func() {
					var returnedErr error
					BeforeEach(func() {
						returnedErr = errors.New("mock error")
						mockSellerRepository.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, returnedErr)
					})

					It("should return the error", func() {
						err := sellerRegisterer.Register(context.Background(), userId, form)
						Expect(err).To(MatchError(returnedErr))
					})
				})

				When("SellerRepository.Create returns no error", func() {
					var (
						expectedSeller *model.Seller
						returnedSeller *model.Seller
						createCall     *gomock.Call
					)
					BeforeEach(func() {
						expectedSeller = helper.Ptr(dummySeller)
						returnedSeller = expectedSeller
						createCall = mockSellerRepository.EXPECT().Create(gomock.Any(), gomock.Any()).Return(returnedSeller, nil)
					})

					JustBeforeEach(func() {
						createCall.Do(func(_ context.Context, seller *model.Seller) {
							Expect(seller).To(Equal(expectedSeller))
						})
					})

					It("should return no error", func() {
						err := sellerRegisterer.Register(context.Background(), userId, form)
						Expect(err).To(Succeed())
					})
				})
			})
		})
	})
})
