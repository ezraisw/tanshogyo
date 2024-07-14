package usecaseimpl_test

import (
	"context"
	"errors"
	"time"

	"github.com/ezraisw/tanshogyo/pkg/common/preseterrors"
	repositorymock "github.com/ezraisw/tanshogyo/pkg/common/repository/mock"
	"github.com/ezraisw/tanshogyo/pkg/common/util/hasher"
	"github.com/ezraisw/tanshogyo/pkg/common/util/helper"
	"github.com/ezraisw/tanshogyo/pkg/common/util/idgen"
	"github.com/ezraisw/tanshogyo/pkg/common/util/timehelper"
	"github.com/ezraisw/tanshogyo/services/user/internal/app/user/model"
	"github.com/ezraisw/tanshogyo/services/user/internal/app/user/usecase"
	usecaseimpl "github.com/ezraisw/tanshogyo/services/user/internal/app/user/usecase/impl"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserLoginer", func() {
	var (
		mockUserRepository           *repositorymock.MockRepository[model.User]
		mockAuthenticationRepository *repositorymock.MockRepository[model.Authentication]
		userLoginer                  *usecaseimpl.UserLoginer

		now   time.Time
		token string
	)
	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		mockUserRepository = repositorymock.NewMockRepository[model.User](ctrl)
		mockAuthenticationRepository = repositorymock.NewMockRepository[model.Authentication](ctrl)
		userLoginer = usecaseimpl.NewUserLoginer(usecaseimpl.UserLoginerOptions{
			UserRepository:           mockUserRepository,
			AuthenticationRepository: mockAuthenticationRepository,
			Nower:                    timehelper.NowerFunc(func() time.Time { return now }),
			IDGen:                    idgen.IDGenFunc(func() string { return token }),
			Hasher:                   hasher.NoHasher,
		})

		now = dummyNow
		token = dummyId
	})

	Context("Login", func() {
		var form usecase.LoginForm
		BeforeEach(func() {
			form = dummyLoginForm
		})

		When("UserRepository.Find returns an error", func() {
			var returnedErr error
			BeforeEach(func() {
				returnedErr = errors.New("mock error")
				mockUserRepository.EXPECT().Find(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, returnedErr)
			})

			It("should return the error", func() {
				_, err := userLoginer.Login(context.Background(), form)
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
				_, err := userLoginer.Login(context.Background(), form)
				Expect(err).To(MatchError(preseterrors.ErrUnauthorized))
			})
		})

		When("UserRepository.Find returns no error", func() {
			var returnedUser *model.User
			BeforeEach(func() {
				returnedUser = helper.Ptr(dummyUser)
				mockUserRepository.EXPECT().Find(gomock.Any(), gomock.Any(), gomock.Any()).Return(returnedUser, nil)
			})

			When("password does not match", func() {
				BeforeEach(func() {
					returnedUser.Password = form.Password + "something"
				})

				It("should return unauthorized error", func() {
					_, err := userLoginer.Login(context.Background(), form)
					Expect(err).To(MatchError(preseterrors.ErrUnauthorized))
				})
			})

			When("password matches", func() {
				BeforeEach(func() {
					returnedUser.Password = form.Password
				})

				When("AuthenticationRepository.Create returns an error", func() {
					var returnedErr error
					BeforeEach(func() {
						returnedErr = errors.New("mock error")
						mockAuthenticationRepository.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, returnedErr)
					})

					It("should return the error", func() {
						_, err := userLoginer.Login(context.Background(), form)
						Expect(err).To(MatchError(returnedErr))
					})
				})

				When("AuthenticationRepository.Create returns no error", func() {
					var (
						expectedAuthentication *model.Authentication
						createCall             *gomock.Call
					)
					BeforeEach(func() {
						expectedAuthentication = &model.Authentication{
							Token:     token,
							UserID:    returnedUser.ID,
							CreatedAt: now.UTC(),
							UpdatedAt: now.UTC(),
							ExpiredAt: now.Add(usecaseimpl.DurationAuthenticationExpired).UTC(),
						}
						createCall = mockAuthenticationRepository.EXPECT().Create(gomock.Any(), gomock.Any()).Return(expectedAuthentication, nil)
					})

					JustBeforeEach(func() {
						createCall.Do(func(_ context.Context, authentication *model.Authentication) {
							Expect(authentication).To(Equal(expectedAuthentication))
						})
					})

					It("should return the correct token", func() {
						expectedResult := usecase.AuthenticationResult{
							Token: token,
						}

						result, err := userLoginer.Login(context.Background(), form)
						Expect(err).To(Succeed())
						Expect(result).To(Equal(expectedResult))
					})
				})
			})
		})
	})
})
