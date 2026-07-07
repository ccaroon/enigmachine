package enigma_test

import (
	"github.com/ccaroon/enigmachine/enigma"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Reflector", func() {
	// YRUHQSLDPXNGOKMIEBFZCWVJAT
	It("Should properly map all valid inputs", func() {
		refB := enigma.GetReflector("B")

		Expect(refB.Reflect(byte('A'))).To(Equal(byte('Y')))
		Expect(refB.Reflect(byte('Z'))).To(Equal(byte('T')))

		for _, letter := range enigma.ALPHABET {
			idx := enigma.LetterToIdx(byte(letter))
			outLetter := refB.Index(idx)

			Expect(refB.Reflect(byte(letter))).To(Equal(outLetter))
		}
	})

	It("Should be symmetrical: A->F <=> F->A", func() {
		refC := enigma.GetReflector("C")

		Expect(refC.Reflect(byte('A'))).To(Equal(byte('F')))
		Expect(refC.Reflect(byte('F'))).To(Equal(byte('A')))
	})

})
