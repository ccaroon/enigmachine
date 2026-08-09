package keysheet_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ccaroon/enigmachine/keysheet"
)

var _ = Describe("Keysheet", func() {
	Context("Generates", func() {
		It("Valid Key Sheet", func() {
			ks := keysheet.GenerateKeySheet("caroon.org", 8, 2026)

			Expect(ks).To(Not(BeNil()))

			Expect(ks.Network).To(Equal("caroon.org"))
			Expect(ks.Month).To(Equal(8))
			Expect(ks.Year).To(Equal(2026))
			Expect(len(ks.Entries)).To(Equal(31))
		})
	})
})
