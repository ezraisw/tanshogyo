package usecaseimpl_test

import (
	"context"
	"errors"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	repositorymock "github.com/pwnedgod/tanshogyo/pkg/common/repository/mock"
	"github.com/pwnedgod/tanshogyo/pkg/common/util/hasher"
	"github.com/pwnedgod/tanshogyo/pkg/common/util/helper"
	"github.com/pwnedgod/tanshogyo/pkg/common/util/idgen"
	"github.com/pwnedgod/tanshogyo/pkg/common/util/timehelper"
	"github.com/pwnedgod/tanshogyo/services/user/internal/app/user/model"
	"github.com/pwnedgod/tanshogyo/services/user/internal/app/user/usecase"
	usecaseimpl "github.com/pwnedgod/tanshogyo/services/user/internal/app/user/usecase/impl"
	usecasemock "github.com/pwnedgod/tanshogyo/services/user/internal/app/user/usecase/mock"
)

var _ = Describe("UserRegisterer", func() {
	var (
		mockUserFormValidator *usecasemock.MockUserFormValidator
		mockUserRepository    *repositorymock.MockRepository[model.User]
		userRegisterer        *usecaseimpl.UserRegisterer

		now time.Time
		id  string
	)
	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		mockUserFormValidator = usecasemock.NewMockUserFormValidator(ctrl)
		mockUserRepository = repositorymock.NewMockRepository[model.User](ctrl)
		userRegisterer = usecaseimpl.NewUserRegisterer(usecaseimpl.UserRegistererOptions{
			UserFormValidator: mockUserFormValidator,
			UserRepository:    mockUserRepository,
			Nower:             timehelper.NowerFunc(func() time.Time { return now }),
			IDGen:             idgen.IDGenFunc(func() string { return id }),
			Hasher:            hasher.NoHasher,
		})

		now = dummyNow
		id = dummyId

	})

	Context("Register", func() {
		var form usecase.UserForm
		BeforeEach(func() {
			form = dummyForm
		})

		When("UserFormValidator.Validate returns an error", func() {
			var returnedErr error
			BeforeEach(func() {
				returnedErr = errors.New("mock error")
				mockUserFormValidator.EXPECT().Validate(gomock.Any(), gomock.Any()).Return(returnedErr)
			})

			It("should return the error", func() {
				err := userRegisterer.Register(context.Background(), form)
				Expect(err).To(MatchError(returnedErr))
			})
		})

		When("UserFormValidator.Validate returns no error", func() {
			var validateCall *gomock.Call
			BeforeEach(func() {
				validateCall = mockUserFormValidator.EXPECT().Validate(gomock.Any(), gomock.Any()).Return(nil)
			})

			JustBeforeEach(func() {
				validateCall.Do(func(_ context.Context, argForm usecase.UserForm) {
					Expect(argForm).To(Equal(form))
				})
			})

			When("UserRepository.Create returns an error", func() {
				var returnedErr error
				BeforeEach(func() {
					returnedErr = errors.New("mock error")
					mockUserRepository.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, returnedErr)
				})

				It("should return the error", func() {
					err := userRegisterer.Register(context.Background(), form)
					Expect(err).To(MatchError(returnedErr))
				})
			})

			When("UserRepository.Create returns no error", func() {
				var (
					expectedUser *model.User
					returnedUser *model.User
					createCall   *gomock.Call
				)
				BeforeEach(func() {
					expectedUser = helper.Ptr(dummyUser)
					returnedUser = expectedUser
					createCall = mockUserRepository.EXPECT().Create(gomock.Any(), gomock.Any()).Return(returnedUser, nil)
				})

				JustBeforeEach(func() {
					createCall.Do(func(_ context.Context, user *model.User) {
						Expect(user).To(Equal(expectedUser))
					})
				})

				It("should return no error", func() {
					err := userRegisterer.Register(context.Background(), form)
					Expect(err).To(Succeed())
				})
			})
		})
	})
})
