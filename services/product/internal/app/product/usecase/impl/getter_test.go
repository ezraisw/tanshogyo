package usecaseimpl_test

import (
	"context"
	"errors"

	repositorymock "github.com/ezraisw/tanshogyo/pkg/common/repository/mock"
	"github.com/ezraisw/tanshogyo/pkg/common/util/helper"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/product/model"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/product/usecase"
	usecaseimpl "github.com/ezraisw/tanshogyo/services/product/internal/app/product/usecase/impl"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProductGetter", func() {
	var (
		mockProductRepository *repositorymock.MockRepository[model.Product]
		productGetter         *usecaseimpl.ProductGetter
	)
	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		mockProductRepository = repositorymock.NewMockRepository[model.Product](ctrl)
		productGetter = usecaseimpl.NewProductGetter(usecaseimpl.ProductGetterOptions{
			ProductRepository: mockProductRepository,
		})
	})

	Context("Get", func() {
		var id string
		BeforeEach(func() {
			id = dummyId
		})

		When("ProductRepository.Find returns an error", func() {
			var returnedErr error
			BeforeEach(func() {
				returnedErr = errors.New("mock error")
				mockProductRepository.EXPECT().Find(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, returnedErr)
			})

			It("should return the error", func() {
				_, err := productGetter.Get(context.Background(), id)
				Expect(err).To(MatchError(returnedErr))
			})
		})

		When("ProductRepository.Find returns no error", func() {
			var returnedProduct *model.Product
			BeforeEach(func() {
				returnedProduct = helper.Ptr(dummyProduct)
				mockProductRepository.EXPECT().Find(gomock.Any(), gomock.Any(), gomock.Any()).Return(returnedProduct, nil)
			})

			It("should return correct dto", func() {
				expectedDto := usecase.Product{
					ID:          returnedProduct.ID,
					SellerID:    returnedProduct.SellerID,
					Name:        returnedProduct.Name,
					Description: returnedProduct.Description,
					Price:       returnedProduct.Price,
					Quantity:    returnedProduct.Quantity,
					CreatedAt:   returnedProduct.CreatedAt,
					UpdatedAt:   returnedProduct.UpdatedAt,
				}

				dto, err := productGetter.Get(context.Background(), id)
				Expect(err).To(Succeed())
				Expect(dto).To(Equal(expectedDto))
			})
		})
	})
})
