package bdd

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "test functions suite")
}

var _ = Describe("Add method",
	func() {
		Context("1과 2를 받으면", func() {
			It("3을 반환한다.", func() {
				Expect(3).To(Equal(Add(1, 2)))
			})
		})
		Context("-1과 1을 받으면", func() {
			It("0을 반환한다.", func() {
				Expect(0).To(Equal(Add(-1, 1)))
			})
		})
	},
)

var _ = Describe("Sub method",
	func() {
		Context("1과 2를 받으면", func() {
			It("-1을 반환한다.", func() {
				Expect(-1).To(Equal(Sub(1, 2)))
			})
		})
		Context("-1과 1을 받으면", func() {
			It("-2를 반환한다.", func() {
				Expect(-23).To(Equal(Sub(-1, 1)))
			})
		})
	},
)
