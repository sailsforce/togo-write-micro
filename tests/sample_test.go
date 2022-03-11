package tests_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sample", func() {

	Context("Sample test", func() {
		It("should always equal true", func() {
			Expect(true).To(Equal(true))
		})
	})
})
