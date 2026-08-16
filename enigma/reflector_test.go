package enigma_test

import (
	"github.com/ccaroon/enigmachine/enigma"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Reflector", func() {

	It("Can get the ID", func() {
		refA, err := enigma.GetReflector('A')
		Expect(err).To(BeNil())
		Expect(refA.Id()).To(Equal('A'))

		refB, err := enigma.GetReflector('B')
		Expect(err).To(BeNil())
		Expect(refB.Id()).To(Equal('B'))

		refC, err := enigma.GetReflector('C')
		Expect(err).To(BeNil())
		Expect(refC.Id()).To(Equal('C'))
	})

	// YRUHQSLDPXNGOKMIEBFZCWVJAT
	It("Should properly map all valid inputs", func() {
		refB, err := enigma.GetReflector('B')

		Expect(err).To(BeNil())
		Expect(refB.Reflect('A')).To(Equal('Y'))
		Expect(refB.Reflect('Z')).To(Equal('T'))

		for _, letter := range enigma.ALPHABET {
			idx := enigma.LetterToIdx(letter)
			outLetter := refB.Index(idx)

			Expect(refB.Reflect(letter)).To(Equal(outLetter))
		}
	})

	It("Should be symmetrical: A->F <=> F->A", func() {
		refC, err := enigma.GetReflector('C')

		Expect(err).To(BeNil())
		Expect(refC.Reflect('A')).To(Equal('F'))
		Expect(refC.Reflect('F')).To(Equal('A'))
	})

	It("Cannot Get an unknown reflector", func() {
		ref, err := enigma.GetReflector('X')

		Expect(ref).To(BeNil())
		Expect(err).To(MatchError("Unsupported Reflector: 'X'"))
	})

})
