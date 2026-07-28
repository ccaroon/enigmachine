package enigma_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ccaroon/enigmachine/enigma"
)

var _ = Describe("Key", func() {

	It("Can Parse a Key Specification String", func() {
		key := enigma.ParseKeySpec("B:I,II,III:AZ,XY,EM")

		Expect(key.ReflId).To(Equal("B"))
		Expect(key.RotorIds).To(Equal([]string{"I", "II", "III"}))
		Expect(key.RotorCfg).To(Equal("AAA"))
		Expect(key.PbSpec).To(Equal([]string{"AZ", "XY", "EM"}))
	})
})
