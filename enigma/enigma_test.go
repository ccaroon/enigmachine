package enigma_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ccaroon/enigmachine/enigma"
)

var _ = Describe("Enigma", func() {

	It("Can encipher a single letter", func() {
		machine := enigma.NewEnigma(
			"B",
			[]string{"I", "II", "III"},
			[]byte{},
		)

		Expect(machine).ToNot(BeNil())

		// -------------------------------------
		// Top => A
		// outLetter := machine.EncipherLetter('A')
		// Expect(outLetter).To(Equal(byte('T')))

		// // Top => B
		// rotorIII := machine.GetRotor(2)
		// Expect(rotorIII.Id()).To(Equal("III"))
		// rotorIII.SetTopLetter('B')
		// outLetter = machine.EncipherLetter('A')
		// Expect(outLetter).To(Equal(byte('S')))
		// -------------------------------------

		// A -> U
		// outLetter := machine.EncipherLetter('A')
		// Expect(outLetter).To(Equal(byte('U')))
		// // reciprocal
		// outLetter = machine.EncipherLetter('U')
		// Expect(outLetter).To(Equal(byte('A')))

		// // G -> P
		// outLetter = machine.EncipherLetter('G')
		// Expect(outLetter).To(Equal(byte('P')))
		// // reciprocal
		// outLetter = machine.EncipherLetter('P')
		// Expect(outLetter).To(Equal(byte('G')))

		// rotorIII := machine.GetRotor(2)
		// Expect(rotorIII.Id()).To(Equal("III"))
		// rotorIII.SetTopLetter('B')

		// outLetter = machine.EncipherLetter('A')
		// Expect(outLetter).To(Equal(byte('B')))

		// outLetter = machine.EncipherLetter('G')
		// Expect(outLetter).To(Equal(byte('X')))
	})

})
