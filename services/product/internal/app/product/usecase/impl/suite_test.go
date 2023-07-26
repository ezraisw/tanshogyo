package usecaseimpl_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestProductUsecases(t *testing.T) {
	RegisterFailHandler(Fail)
	suiteConfig, reporterConfig := GinkgoConfiguration()
	reporterConfig.NoColor = true // Prevents unsupported ANSI color characters.
	RunSpecs(t, "Product Usecase Suite", suiteConfig, reporterConfig)
}
