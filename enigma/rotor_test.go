package enigma_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ccaroon/enigmachine/enigma"
)

var _ = Describe("Rotor", func() {

	It("Should Invert the Wiring On Init", func() {
		rotor := enigma.GetRotor("I")
		Expect((rotor.Inverse)).To(Equal("UWYGADFPVZBECKMTHXSLRINQOJ"))
	})

})
