package models_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ccaroon/enigma/models"
)

var _ = Describe("Rotor", func() {

	It("Should Invert the Wiring On Init", func() {
		rotor := models.NewRotor("I", "ZYXWVUTSRQPONMLKJIHGFEDCBA")

		Expect(rotor.Inverse).To(Equal("ABCDEFGHIJKLMNOPQRSTUVWXYZ"))
	})

})
