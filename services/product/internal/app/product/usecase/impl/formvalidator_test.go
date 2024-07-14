package usecaseimpl_test

import (
	"context"

	"github.com/ezraisw/tanshogyo/pkg/common/preseterrors"
	"github.com/ezraisw/tanshogyo/pkg/common/rules"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/product/usecase"
	usecaseimpl "github.com/ezraisw/tanshogyo/services/product/internal/app/product/usecase/impl"
	factorymock "github.com/ezraisw/tanshogyo/services/product/internal/app/seller/factory/mock"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProductFormValidator", func() {
	var (
		mockSellerExistsRuleFactory *factorymock.MockSellerExistsRuleFactory
		productFormValidator        *usecaseimpl.ProductFormValidator
	)
	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		mockSellerExistsRuleFactory = factorymock.NewMockSellerExistsRuleFactory(ctrl)
		productFormValidator = usecaseimpl.NewProductFormValidator(usecaseimpl.ProductFormValidatorOptions{
			SellerExistsRuleFactory: mockSellerExistsRuleFactory,
		})

		mockSellerExistsRuleFactory.EXPECT().
			Make(gomock.Any()).
			Return(validation.By(func(value interface{}) error { return nil })).
			AnyTimes()
	})

	DescribeTable("Validate", func(form usecase.ProductForm, expectedFieldErr *preseterrors.FieldError) {
		err, _ := productFormValidator.Validate(context.Background(), form).(*preseterrors.ValidationError)
		if expectedFieldErr == nil {
			Expect(err).ToNot(HaveOccurred())
		} else {
			Expect(err).To(HaveOccurred())
			Expect(err.FieldErrors).To(ContainElement(expectedFieldErr))
		}
	},
		Entry("When given empty SellerID", usecase.ProductForm{SellerID: ""}, rules.ToFieldError("sellerId", validation.ErrRequired)),
		Entry("When given empty Name", usecase.ProductForm{Name: ""}, rules.ToFieldError("name", validation.ErrRequired)),
		Entry("When given empty Description", usecase.ProductForm{Description: ""}, rules.ToFieldError("description", validation.ErrRequired)),
		Entry("When given empty Price", usecase.ProductForm{Price: 0}, rules.ToFieldError("price", validation.ErrRequired)),
		Entry("When given Price less than 0", usecase.ProductForm{Price: -100}, rules.ToFieldErrorWithParams("price", validation.ErrMinGreaterThanRequired, map[string]any{
			"threshold": 0,
		})),
		Entry("When given Quantity less than 0", usecase.ProductForm{Quantity: -100}, rules.ToFieldErrorWithParams("quantity", validation.ErrMinGreaterEqualThanRequired, map[string]any{
			"threshold": 0,
		})),
	)
})
