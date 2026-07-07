package enigma_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ccaroon/enigmachine/enigma"
)

var _ = Describe("Rotor", func() {

	It("Should pass the Identity Test", func() {
		rotor := enigma.NewRotor("0", enigma.ALPHABET, 'A')

		for _, letter := range enigma.ALPHABET {
			Expect(rotor.Forward(byte(letter))).To(Equal(byte(letter)))
		}
	})

	It("Rotor I: Forward -> Top Letter 'A'", func() {
		rotor := enigma.GetRotor("I")

		Expect(rotor.Forward('A')).To(Equal(byte('E')))
		Expect(rotor.Forward('B')).To(Equal(byte('K')))
		// ...
		Expect(rotor.Forward('Y')).To(Equal(byte('C')))
		Expect(rotor.Forward('Z')).To(Equal(byte('J')))
	})

	It("Rotor I: Forward -> Top Letter 'B'", func() {
		rotor := enigma.GetRotor("I")
		rotor.SetTopLetter('B')

		Expect(rotor.Forward('A')).To(Equal(byte('J')))
		Expect(rotor.Forward('B')).To(Equal(byte('E')))
		// ...
		Expect(rotor.Forward('Y')).To(Equal(byte('R')))
		Expect(rotor.Forward('Z')).To(Equal(byte('C')))
	})

	It("Roto I: Reverse -> Top Letter 'A'", func() {
		rotor := enigma.GetRotor("I")

		Expect(rotor.Reverse('A')).To(Equal(byte('U')))
		Expect(rotor.Reverse('F')).To(Equal(byte('D')))
		// ...
		Expect(rotor.Reverse('Y')).To(Equal(byte('O')))
		Expect(rotor.Reverse('Z')).To(Equal(byte('J')))
	})

	It("Roto I: Reverse -> Top Letter 'B'", func() {
		rotor := enigma.GetRotor("I")
		rotor.SetTopLetter('B')

		Expect(rotor.Reverse('A')).To(Equal(byte('V')))
		Expect(rotor.Reverse('F')).To(Equal(byte('E')))
		// ...
		Expect(rotor.Reverse('Y')).To(Equal(byte('P')))
		Expect(rotor.Reverse('Z')).To(Equal(byte('K')))
	})

	It("Top Letter Test", func() {
		//   ABCDEFGHIJKLMNOPQRSTUVWXYZ
		//     BACDEFGHIJKLMNOPQRSTUVWXYZ
		rotor := enigma.NewRotor("X", "BACDEFGHIJKLMNOPQRSTUVWXYZ", 'A')

		Expect(rotor.Forward('A')).To(Equal(byte('B')))
		Expect(rotor.Reverse('A')).To(Equal(byte('B')))

		rotor.SetTopLetter('B')
		Expect(rotor.Forward('A')).To(Equal(byte('Z')))
		Expect(rotor.Reverse('A')).To(Equal(byte('C')))

		rotor.SetTopLetter('C')
		Expect(rotor.Forward('A')).To(Equal(byte('Y')))
		Expect(rotor.Reverse('A')).To(Equal(byte('D')))
	})

	// map_r_to_l rotor_III 'O' `14` should be `17`.
	It("Rotor III: Forward 'O' 14 -> 17", func() {
		// ABCDEFGHIJKLMNOPQRSTUVWXYZ
		//                         |
		//               BDFHJLCPRTXVZNYEIWGAKMUSQO
		rotor := enigma.GetRotor("III")
		rotor.SetTopLetter('O')

		input := rotor.Index(14) // <- 'Y'
		output := byte('X')

		Expect(rotor.Forward(input)).To(Equal(output))
	})

	// `map_l_to_r rotor_I 'F' 10` should be `14`.
	It("Rotor I: Reverse 'F' 10 -> 14", func() {
		// ABCDEFGHIJKLMNOPQRSTUVWXYZ
		//                |
		//      EKMFLGDQVZNTOWYHXUSPAIBRCJ
		rotor := enigma.GetRotor("I")
		rotor.SetTopLetter('F')

		input := rotor.Index(10)
		output := byte('P')
		Expect(rotor.Reverse(input)).To(Equal(output))
	})

	It("Rotor II: Starting K -- Misc", func() {
		rotor := enigma.GetRotor("II")
		rotor.SetTopLetter('K')

		Expect(rotor.Forward(byte('A'))).To(Equal(byte('Q')))
		Expect(rotor.Forward(byte('Q'))).To(Equal(byte('R')))
		Expect(rotor.Forward(byte('R'))).To(Equal(byte('U')))
		Expect(rotor.Forward(byte('U'))).To(Equal(byte('L')))
		Expect(rotor.Forward(byte('Z'))).To(Equal(byte('C')))
		// ...
		Expect(rotor.Reverse(byte('A'))).To(Equal(byte('K')))
		Expect(rotor.Reverse(byte('Z'))).To(Equal(byte('C')))
		Expect(rotor.Reverse(byte('H'))).To(Equal(byte('V')))
		Expect(rotor.Reverse(byte('T'))).To(Equal(byte('X')))
		Expect(rotor.Reverse(byte('Q'))).To(Equal(byte('A')))

	})

})
