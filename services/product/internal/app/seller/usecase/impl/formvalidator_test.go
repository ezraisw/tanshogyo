package usecaseimpl_test

import (
	"context"

	"github.com/ezraisw/tanshogyo/pkg/common/preseterrors"
	"github.com/ezraisw/tanshogyo/pkg/common/rules"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/seller/usecase"
	usecaseimpl "github.com/ezraisw/tanshogyo/services/product/internal/app/seller/usecase/impl"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SellerFormValidator", func() {
	var (
		userFormValidator *usecaseimpl.SellerFormValidator
	)
	BeforeEach(func() {
		userFormValidator = usecaseimpl.NewSellerFormValidator()
	})

	DescribeTable("Validate", func(form usecase.SellerForm, expectedFieldErr *preseterrors.FieldError) {
		err, _ := userFormValidator.Validate(context.Background(), form).(*preseterrors.ValidationError)
		if expectedFieldErr == nil {
			Expect(err).ToNot(HaveOccurred())
		} else {
			Expect(err).To(HaveOccurred())
			Expect(err.FieldErrors).To(ContainElement(expectedFieldErr))
		}
	},
		Entry("When ShopName is empty", usecase.SellerForm{ShopName: ""}, rules.ToFieldError("shopName", validation.ErrRequired)),
	)
})
