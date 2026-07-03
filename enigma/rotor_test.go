package enigma_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ccaroon/enigmachine/enigma"
)

var _ = Describe("Rotor", func() {

	// It("Should Invert the Wiring On Init", func() {
	// 	rotor := enigma.GetRotor("I")
	// 	Expect((rotor.Inverse)).To(Equal("UWYGADFPVZBECKMTHXSLRINQOJ"))
	// })

	// It("Should pass the Identity Test", func() {
	// 	rotor := enigma.NewRotor("0", enigma.ALPHABET, 'A')

	// 	for _, letter := range enigma.ALPHABET {
	// 		Expect(rotor.Forward(byte(letter))).To(Equal(byte(letter)))
	// 	}
	// })

	// It("Rotor I: Forward -> Starting Setting 'A'", func() {
	// 	rotor := enigma.GetRotor("I")

	// 	Expect(rotor.Forward('A')).To(Equal(byte('E')))
	// 	Expect(rotor.Forward('B')).To(Equal(byte('K')))
	// 	// ...
	// 	Expect(rotor.Forward('Y')).To(Equal(byte('C')))
	// 	Expect(rotor.Forward('Z')).To(Equal(byte('J')))
	// })

	// It("Rotor I: Forward -> Starting Setting 'B'", func() {
	// 	rotor := enigma.GetRotor("I")
	// 	rotor.SetTopLetter('B')

	// 	Expect(rotor.Forward('A')).To(Equal(byte('K')))
	// 	Expect(rotor.Forward('B')).To(Equal(byte('M')))
	// 	// ...
	// 	Expect(rotor.Forward('Y')).To(Equal(byte('J')))
	// 	Expect(rotor.Forward('Z')).To(Equal(byte('E')))
	// })

	// It("Roto I: Reverse -> Starting Setting 'A'", func() {
	// 	rotor := enigma.GetRotor("I")

	// 	Expect(rotor.Reverse('A')).To(Equal(byte('U')))
	// 	Expect(rotor.Reverse('F')).To(Equal(byte('D')))
	// 	// ...
	// 	Expect(rotor.Reverse('Y')).To(Equal(byte('O')))
	// 	Expect(rotor.Reverse('Z')).To(Equal(byte('J')))
	// })

	// It("Roto I: Reverse -> Starting Setting 'B'", func() {
	// 	rotor := enigma.GetRotor("I")
	// 	rotor.SetTopLetter('B')

	// 	Expect(rotor.Reverse('A')).To(Equal(byte('W')))
	// 	Expect(rotor.Reverse('F')).To(Equal(byte('F')))
	// 	// ...
	// 	Expect(rotor.Reverse('Y')).To(Equal(byte('J')))
	// 	Expect(rotor.Reverse('Z')).To(Equal(byte('U')))
	// })

	// It("Top Letter Test", func() {
	// 	//          ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// 	//        BACDEFGHIJKLMNOPQRSTUVWXYZ
	// 	rotor := enigma.NewRotor("X", "BACDEFGHIJKLMNOPQRSTUVWXYZ", 'A')

	// 	Expect(rotor.Forward('A')).To(Equal(byte('B')))
	// 	Expect(rotor.Reverse('A')).To(Equal(byte('B')))

	// 	rotor.SetTopLetter('B')
	// 	Expect(rotor.Forward('A')).To(Equal(byte('A')))
	// 	Expect(rotor.Reverse('A')).To(Equal(byte('A')))

	// 	rotor.SetTopLetter('C')
	// 	Expect(rotor.Forward('A')).To(Equal(byte('C')))
	// 	Expect(rotor.Reverse('A')).To(Equal(byte('C')))
	// })

	// map_r_to_l rotor_III 'O' `14` should be `17`.
	It("Rotor III: Forward 'O' 14 -> 17", func() {
		// ABCDEFGHIJKLMNOPQRSTUVWXYZ
		// BDFHJLCPRTXVZNYEIWGAKMUSQO
		rotor := enigma.GetRotor("III")
		rotor.SetTopLetter('O')

		input := enigma.IdxToLetter(14)  // O
		output := enigma.IdxToLetter(17) // R
		fmt.Printf("%c -> %c\n", input, output)
		Expect(rotor.Forward(input)).To(Equal(output))
	})

	// If `rotor_I = "EKMFLGDQVZNTOWYHXUSPAIBRCJ"`, then
	// `map_l_to_r rotor_I 'F' 10` should be `14`.

})
