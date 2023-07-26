package usecaseimpl_test

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pwnedgod/tanshogyo/pkg/common/preseterrors"
	"github.com/pwnedgod/tanshogyo/pkg/common/rules"
	factorymock "github.com/pwnedgod/tanshogyo/services/user/internal/app/user/factory/mock"
	"github.com/pwnedgod/tanshogyo/services/user/internal/app/user/usecase"
	usecaseimpl "github.com/pwnedgod/tanshogyo/services/user/internal/app/user/usecase/impl"
)

var _ = Describe("UserFormValidator", func() {
	var (
		mockUserUniqueRuleFactory *factorymock.MockUserUniqueRuleFactory
		userFormValidator         *usecaseimpl.UserFormValidator
	)
	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		mockUserUniqueRuleFactory = factorymock.NewMockUserUniqueRuleFactory(ctrl)
		userFormValidator = usecaseimpl.NewUserFormValidator(usecaseimpl.UserFormValidatorOptions{
			UserUniqueRuleFactory: mockUserUniqueRuleFactory,
		})

		mockUserUniqueRuleFactory.EXPECT().
			Make(gomock.Any(), gomock.Any()).
			Return(validation.By(func(any) error { return nil })).
			AnyTimes()
	})

	DescribeTable("Validate", func(form usecase.UserForm, expectedFieldErr *preseterrors.FieldError) {
		err, _ := userFormValidator.Validate(context.Background(), form).(*preseterrors.ValidationError)
		if expectedFieldErr == nil {
			Expect(err).ToNot(HaveOccurred())
		} else {
			Expect(err).To(HaveOccurred())
			Expect(err.FieldErrors).To(ContainElement(expectedFieldErr))
		}
	},
		Entry("When Username is empty", usecase.UserForm{Username: ""}, rules.ToFieldError("username", validation.ErrRequired)),
		Entry("When Username length is less than 4", usecase.UserForm{Username: "lol"}, rules.ToFieldErrorWithParams("username", validation.ErrLengthTooShort, map[string]any{
			"min": 4,
			"max": 0,
		})),
		Entry("When Password is empty", usecase.UserForm{Password: ""}, rules.ToFieldError("password", validation.ErrRequired)),
		Entry("When Password length is less than 8", usecase.UserForm{Password: "P4ss"}, rules.ToFieldErrorWithParams("password", validation.ErrLengthTooShort, map[string]any{
			"min": 8,
			"max": 0,
		})),
		Entry("When Password does not contain upper case, lower case, or digit", usecase.UserForm{Password: "p4ssw0rd"}, rules.ToFieldError("password", rules.ErrLowerCaseUpperCaseAndDigits)),
		Entry("When Email is empty", usecase.UserForm{Email: ""}, rules.ToFieldError("email", validation.ErrRequired)),
		Entry("When Email is not an email", usecase.UserForm{Email: "invalid"}, rules.ToFieldError("email", is.ErrEmail)),
		Entry("When Name is empty", usecase.UserForm{Name: ""}, rules.ToFieldError("name", validation.ErrRequired)),
	)
})
