package enigma_test

import (
	"regexp"

	"github.com/ccaroon/enigmachine/enigma"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Alphabet", func() {

	It("Can convert a letter (A-Z) to an index (0-25)", func() {
		for idx, letter := range enigma.ALPHABET {
			Expect(enigma.LetterToIdx(letter)).To(Equal(idx))
		}
	})

	It("Can convert an index(0-25) to a letter (A-Z)", func() {
		for idx := range 26 {
			expectedLetter := rune(enigma.ALPHABET[idx])
			Expect(enigma.IdxToLetter(idx)).To(Equal(expectedLetter))
		}
	})

	It("Can get a string of random letters", func() {
		rndString := enigma.RandomLetters(7)

		Expect(len(rndString)).To(Equal(7))

		// should only contain upper-case A-Z
		ok, err := regexp.MatchString("^[A-Z]+$", rndString)
		Expect(err).To(BeNil())
		Expect(ok).To(BeTrue())
	})
})
