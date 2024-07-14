package usecaseimpl_test

import (
	"context"
	"errors"

	repositorymock "github.com/ezraisw/tanshogyo/pkg/common/repository/mock"
	"github.com/ezraisw/tanshogyo/services/user/internal/app/user/model"
	"github.com/ezraisw/tanshogyo/services/user/internal/app/user/usecase"
	usecaseimpl "github.com/ezraisw/tanshogyo/services/user/internal/app/user/usecase/impl"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserUniqueChecker", func() {
	var (
		mockUserRepository *repositorymock.MockRepository[model.User]
		userUniqueChecker  *usecaseimpl.UserUniqueChecker
	)
	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		mockUserRepository = repositorymock.NewMockRepository[model.User](ctrl)
		userUniqueChecker = usecaseimpl.NewUserUniqueChecker(usecaseimpl.UserUniqueCheckerOptions{
			UserRepository: mockUserRepository,
		})
	})

	Context("CheckUnique", func() {
		var (
			excludedId string
			fields     []usecase.Field
		)
		BeforeEach(func() {
			excludedId = dummyId
			fields = dummyFields
		})

		When("UserRepository.Exists returns an error", func() {
			var returnedErr error
			BeforeEach(func() {
				returnedErr = errors.New("mock error")
				mockUserRepository.EXPECT().Exists(gomock.Any(), gomock.Any()).Return(false, returnedErr)
			})

			It("should return the error", func() {
				_, err := userUniqueChecker.CheckUnique(context.Background(), excludedId, fields)
				Expect(err).To(MatchError(returnedErr))
			})
		})

		When("UserRepository.Exists returns no error", func() {
			var returnedExists bool
			BeforeEach(func() {
				returnedExists = true
				mockUserRepository.EXPECT().Exists(gomock.Any(), gomock.Any()).Return(returnedExists, nil)
			})

			When("excludedId is empty", func() {
				excludedId = ""

				It("should return unique as opposite of exists and no error", func() {
					expectedUnique := !returnedExists

					unique, err := userUniqueChecker.CheckUnique(context.Background(), excludedId, fields)
					Expect(err).To(Succeed())
					Expect(unique).To(Equal(expectedUnique))
				})
			})

			When("excludedId is not empty", func() {
				excludedId = dummyId

				It("should return unique as opposite of exists and no error", func() {
					expectedUnique := !returnedExists

					unique, err := userUniqueChecker.CheckUnique(context.Background(), excludedId, fields)
					Expect(err).To(Succeed())
					Expect(unique).To(Equal(expectedUnique))
				})
			})
		})
	})
})
