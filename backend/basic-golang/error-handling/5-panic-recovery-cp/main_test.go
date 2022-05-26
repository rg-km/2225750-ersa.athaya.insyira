package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Wrapping Error", func() {
	Describe("Panic revocery in printBook", func() {
		It("returns wrapped error", func() {
			printBook(10)
			Expect("0. Books  is: The Eye of the World").To(Equal("0. Books  is: The Eye of the World"))
		})
	})
})
