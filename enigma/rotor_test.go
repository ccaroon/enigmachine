package enigma_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ccaroon/enigmachine/enigma"
)

func inverseWiring(wiring string) string {
	inverse := make([]byte, 26)
	for idx, letter := range wiring {
		newIdx := enigma.LetterToIdx(byte(letter))
		newLtr := enigma.IdxToLetter(byte(idx))
		inverse[newIdx] = newLtr
	}

	return string(inverse)
}

func fullTest(rotor *enigma.Rotor) {
	expectedFwd := rotor.Wiring()
	for idx, letter := range enigma.ALPHABET {
		outLetter := byte(expectedFwd[idx])
		Expect(rotor.Forward(byte(letter))).To(Equal(outLetter))
	}

	expectedRev := inverseWiring(expectedFwd)
	for idx, letter := range enigma.ALPHABET {
		outLetter := byte(expectedRev[idx])
		Expect(rotor.Reverse(byte(letter))).To(Equal(outLetter))
	}
}

var _ = Describe("Rotor", func() {

	It("Should pass the Identity Test", func() {
		rotor := enigma.GetRotor("0")
		fullTest(rotor)
	})

	It("Can Set Top Letter", func() {
		rotor := enigma.GetRotor("0")

		rotor.SetTopLetter('A')
		Expect(rotor.Forward('A')).To(Equal(byte('A')))

		rotor.SetTopLetter('B')
		Expect(rotor.Forward('A')).To(Equal(byte('B')))

		rotor.SetTopLetter('C')
		Expect(rotor.Forward('A')).To(Equal(byte('C')))

		rotor.SetTopLetter('J')
		Expect(rotor.Forward('A')).To(Equal(byte('J')))

		rotor.SetTopLetter('Q')
		Expect(rotor.Forward('A')).To(Equal(byte('Q')))

		rotor.SetTopLetter('X')
		Expect(rotor.Forward('A')).To(Equal(byte('X')))
	})

	It("Rotor I: Top Letter 'A'", func() {
		rotor := enigma.GetRotor("I")
		fullTest(rotor)
	})

	It("Rotor II: Top Letter 'A'", func() {
		rotor := enigma.GetRotor("II")
		fullTest(rotor)
	})

	It("Rotor III: Top Letter 'A'", func() {
		rotor := enigma.GetRotor("III")
		fullTest(rotor)
	})

	It("Rotor IV: Top Letter 'A'", func() {
		rotor := enigma.GetRotor("IV")
		fullTest(rotor)
	})

	It("Rotor V: Top Letter 'A'", func() {
		rotor := enigma.GetRotor("V")
		fullTest(rotor)
	})

	It("Rotor I: Top Letter 'B'", func() {
		rotor := enigma.GetRotor("I")
		rotor.SetTopLetter('B')

		fullTest(rotor)
	})

	It("Top Letter Test", func() {
		// I:    ABCDEFGHIJKLMNOPQRSTUVWXYZ
		// O:  BACDEFGHIJKLMNOPQRSTUVWXYZ
		rotor := enigma.NewRotor("X", "BACDEFGHIJKLMNOPQRSTUVWXYZ", 'A')

		Expect(rotor.Forward('A')).To(Equal(byte('B')))
		Expect(rotor.Reverse('A')).To(Equal(byte('B')))

		rotor.SetTopLetter('B')
		Expect(rotor.Forward('A')).To(Equal(byte('A')))
		Expect(rotor.Reverse('A')).To(Equal(byte('A')))

		rotor.SetTopLetter('C')
		Expect(rotor.Forward('A')).To(Equal(byte('C')))
		Expect(rotor.Reverse('A')).To(Equal(byte('Z')))
	})

	// map_r_to_l rotor_III 'O' `14` should be `17`.
	It("Rotor III: Top Letter 'O' // Forward // Pos 14 -> 17", func() {
		// ABCDEFGHIJKLMNOPQRSTUVWXYZ
		//               |
		// YEIWGAKMUSQOBDFHJLCPRTXVZN
		rotor := enigma.GetRotor("III")
		rotor.SetTopLetter('O')

		input := enigma.IdxToLetter(14) // 'O'
		output := byte('F')

		Expect(rotor.Forward(input)).To(Equal(output))
	})

	// `map_l_to_r rotor_I 'F' 10` should be `14`.
	It("Rotor I: Top Letter 'F' // Reverse // Pos 10 -> 14", func() {
		//          ABCDEFGHIJKLMNOPQRSTUVWXYZ
		//    EKMFLGDQVZNTOWYHXUSPAIBRCJ
		rotor := enigma.GetRotor("I")
		rotor.SetTopLetter('F')

		input := enigma.IdxToLetter(10) // 'K'
		output := byte('W')
		Expect(rotor.Reverse(input)).To(Equal(output))
	})

	It("Rotor II: Top Letter 'K' // Wrap Test // F&R", func() {
		rotor := enigma.GetRotor("II")
		rotor.SetTopLetter('K')

		fullTest(rotor)
		// // Forward
		// expectedFwd := rotor.Wiring()
		// for idx, letter := range enigma.ALPHABET {
		// 	outLetter := byte(expectedFwd[idx])
		// 	Expect(rotor.Forward(byte(letter))).To(Equal(outLetter))
		// }

		// // Reverse - NO WRAP
		// inLettersss := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		// expectedRev := "QZFSPMHBVRTAEJOKGWUDXNCYLI"
		// // inLettersss := "AX"
		// // expectedRev := "QY"
		// for idx, letter := range inLettersss {
		// 	if letter == '-' {
		// 		continue
		// 	}
		// 	outLetter := byte(expectedRev[idx])
		// 	Expect(rotor.Reverse(byte(letter))).To(Equal(outLetter))
		// }
	})

	It("Chain I/O -> III -> II -> ReflB", func() {
		rotorII := enigma.GetRotor("II")
		rotorIII := enigma.GetRotor("III")
		reflB := enigma.GetReflector("B")

		input := byte('A')

		// Baseline
		output := rotorIII.Forward(input)
		Expect(output).To(Equal(byte('B')))

		output = rotorII.Forward(output)
		Expect(output).To(Equal(byte('J')))

		output = reflB.Reflect(output)
		Expect(output).To(Equal(byte('X')))

		// Step III
		// rotorIII.SetTopLetter('B')

		// output = rotorIII.Forward(input)
		// Expect(output).To(Equal(byte('D')))

		// output = rotorII.Forward(output)
		// Expect(output).To(Equal(byte('D')))

		// output = reflB.Reflect(output)
		// Expect(output).To(Equal(byte('H')))
	})

})
