package main_test

import (
	. "github.com/smousa/chatty"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Urlreader", func() {
	var (
		r     URLReader
		input string
	)

	BeforeEach(func() {
		r = &DefaultURLReader{}
	})

	Context("http://www.google.com", func() {
		BeforeEach(func() {
			input = "http://www.google.com"
		})

		Context("Page title", func() {
			var (
				out string
				err error
			)

			JustBeforeEach(func() {
				out, err = r.PageTitle(input)
			})

			It("should display 'Google' as the page title", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(out).To(Equal("Google"))
			})
		})
	})
})
