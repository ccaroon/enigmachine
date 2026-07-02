package models_test

import (
	"github.com/ccaroon/enigma/models"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Alphabet", func() {

	It("Can convert a letter (A-Z) to an index (0-25)", func() {
		for idx, letter := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
			Expect(models.LetterToIdx(byte(letter))).To(Equal(int(idx)))
		}
	})

})
