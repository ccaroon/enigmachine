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

	It("Should pass the Identity Test", func() {
		rotor := enigma.NewRotor("0", enigma.ALPHABET, 'A')

		for _, letter := range enigma.ALPHABET {
			Expect(rotor.Forward(byte(letter))).To(Equal(byte(letter)))
		}
	})

	It("Rotor I: Starting Setting 'A'", func() {
		rotor := enigma.GetRotor("I")

		Expect(rotor.Forward('A')).To(Equal(byte('E')))
		Expect(rotor.Forward('B')).To(Equal(byte('K')))
		// ...
		Expect(rotor.Forward('Y')).To(Equal(byte('C')))
		Expect(rotor.Forward('Z')).To(Equal(byte('J')))
	})

	It("Rotor I: Starting Setting 'B'", func() {
		rotor := enigma.GetRotor("I")
		rotor.SetTopLetter('B')

		Expect(rotor.Forward('A')).To(Equal(byte('K')))
		Expect(rotor.Forward('B')).To(Equal(byte('M')))
		// ...
		Expect(rotor.Forward('Y')).To(Equal(byte('J')))
		Expect(rotor.Forward('Z')).To(Equal(byte('E')))
	})

})
