package usecaseimpl_test

import (
	"context"
	"errors"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pwnedgod/tanshogyo/pkg/common/preseterrors"
	repositorymock "github.com/pwnedgod/tanshogyo/pkg/common/repository/mock"
	"github.com/pwnedgod/tanshogyo/pkg/common/util/helper"
	"github.com/pwnedgod/tanshogyo/pkg/common/util/timehelper"
	"github.com/pwnedgod/tanshogyo/services/user/internal/app/user/model"
	"github.com/pwnedgod/tanshogyo/services/user/internal/app/user/usecase"
	usecaseimpl "github.com/pwnedgod/tanshogyo/services/user/internal/app/user/usecase/impl"
)

var _ = Describe("UserAuthenticator", func() {
	var (
		mockAuthenticationRepository *repositorymock.MockRepository[model.Authentication]
		mockUserRepository           *repositorymock.MockRepository[model.User]
		userAuthenticator            *usecaseimpl.UserAuthenticator

		now time.Time
	)
	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		mockUserRepository = repositorymock.NewMockRepository[model.User](ctrl)
		mockAuthenticationRepository = repositorymock.NewMockRepository[model.Authentication](ctrl)
		userAuthenticator = usecaseimpl.NewUserAuthenticator(usecaseimpl.UserAuthenticatorOptions{
			AuthenticationRepository: mockAuthenticationRepository,
			UserRepository:           mockUserRepository,
			Nower:                    timehelper.NowerFunc(func() time.Time { return now }),
		})

		now = dummyNow
	})

	Context("Authenticate", func() {
		var token string
		BeforeEach(func() {
			token = dummyId
		})

		When("AuthenticationRepository.Find returns an error", func() {
			var returnedErr error
			BeforeEach(func() {
				returnedErr = errors.New("mock error")
				mockAuthenticationRepository.EXPECT().Find(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, returnedErr)
			})

			It("should return the error", func() {
				_, err := userAuthenticator.Authenticate(context.Background(), token)
				Expect(err).To(MatchError(returnedErr))
			})
		})

		When("AuthenticationRepository.Find returns not found error", func() {
			var returnedErr error
			BeforeEach(func() {
				returnedErr = preseterrors.ErrNotFound
				mockAuthenticationRepository.EXPECT().Find(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, returnedErr)
			})

			It("should return unauthorized error", func() {
				_, err := userAuthenticator.Authenticate(context.Background(), token)
				Expect(err).To(MatchError(preseterrors.ErrUnauthorized))
			})
		})

		When("AuthenticationRepository.Find returns no error", func() {
			var returnedAuthentication *model.Authentication
			BeforeEach(func() {
				returnedAuthentication = helper.Ptr(dummyAuthentication)
				mockAuthenticationRepository.EXPECT().Find(gomock.Any(), gomock.Any(), gomock.Any()).Return(returnedAuthentication, nil)
			})

			When("UserRepository.Find returns an error", func() {
				var returnedErr error
				BeforeEach(func() {
					returnedErr = errors.New("mock error")
					mockUserRepository.EXPECT().Find(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, returnedErr)
				})

				It("should return the error", func() {
					_, err := userAuthenticator.Authenticate(context.Background(), token)
					Expect(err).To(MatchError(returnedErr))
				})
			})

			When("UserRepository.Find returns not found error", func() {
				var returnedErr error
				BeforeEach(func() {
					returnedErr = preseterrors.ErrNotFound
					mockUserRepository.EXPECT().Find(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, returnedErr)
				})

				It("should return unauthorized error", func() {
					_, err := userAuthenticator.Authenticate(context.Background(), token)
					Expect(err).To(MatchError(preseterrors.ErrUnauthorized))
				})
			})

			When("UserRepository.Find returns no error", func() {
				var returnedUser *model.User
				BeforeEach(func() {
					returnedUser = helper.Ptr(dummyUser)
					mockUserRepository.EXPECT().Find(gomock.Any(), gomock.Any(), gomock.Any()).Return(returnedUser, nil)
				})

				It("should return user and no error", func() {
					expectedDto := usecase.User{
						ID:        returnedUser.ID,
						Username:  returnedUser.Username,
						Email:     returnedUser.Email,
						Name:      returnedUser.Name,
						CreatedAt: returnedUser.CreatedAt,
						UpdatedAt: returnedUser.UpdatedAt,
					}

					dto, err := userAuthenticator.Authenticate(context.Background(), token)
					Expect(err).To(Succeed())
					Expect(dto).To(Equal(expectedDto))
				})
			})
		})
	})
})
