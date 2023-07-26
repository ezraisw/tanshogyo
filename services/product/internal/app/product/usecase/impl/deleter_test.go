package usecaseimpl_test

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	repositorymock "github.com/pwnedgod/tanshogyo/pkg/common/repository/mock"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/model"
	usecaseimpl "github.com/pwnedgod/tanshogyo/services/product/internal/app/product/usecase/impl"
)

var _ = Describe("ProductDeleter", func() {
	var (
		mockProductRepository *repositorymock.MockRepository[model.Product]
		productDeleter        *usecaseimpl.ProductDeleter
	)
	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		mockProductRepository = repositorymock.NewMockRepository[model.Product](ctrl)
		productDeleter = usecaseimpl.NewProductDeleter(usecaseimpl.ProductDeleterOptions{
			ProductRepository: mockProductRepository,
		})
	})

	Context("Delete", func() {
		var id string
		BeforeEach(func() {
			id = dummyId
		})

		When("ProductRepository.Delete returns an error", func() {
			var returnedErr error
			BeforeEach(func() {
				returnedErr = errors.New("mock error")
				mockProductRepository.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(returnedErr)
			})

			It("should return the error", func() {
				err := productDeleter.Delete(context.Background(), id)
				Expect(err).To(MatchError(returnedErr))
			})
		})

		When("ProductRepository.Delete returns no error", func() {
			BeforeEach(func() {
				mockProductRepository.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)
			})

			It("should return no error", func() {
				err := productDeleter.Delete(context.Background(), id)
				Expect(err).To(Succeed())
			})
		})
	})
})
