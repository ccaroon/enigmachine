package enigma_test

import (
	"github.com/ccaroon/enigmachine/enigma"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Reflector", func() {

	It("Can get the ID", func() {
		refB := enigma.GetReflector("B")
		Expect(refB.Id()).To(Equal("B"))
	})

	// YRUHQSLDPXNGOKMIEBFZCWVJAT
	It("Should properly map all valid inputs", func() {
		refB := enigma.GetReflector("B")

		Expect(refB.Reflect('A')).To(Equal('Y'))
		Expect(refB.Reflect('Z')).To(Equal('T'))

		for _, letter := range enigma.ALPHABET {
			idx := enigma.LetterToIdx(letter)
			outLetter := refB.Index(idx)

			Expect(refB.Reflect(letter)).To(Equal(outLetter))
		}
	})

	It("Should be symmetrical: A->F <=> F->A", func() {
		refC := enigma.GetReflector("C")

		Expect(refC.Reflect('A')).To(Equal('F'))
		Expect(refC.Reflect('F')).To(Equal('A'))
	})
})
