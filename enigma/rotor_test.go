package enigma_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ccaroon/enigmachine/enigma"
)

// func inverseWiring(wiring string) string {
// 	inverse := make([]byte, 26)
// 	for idx, letter := range wiring {
// 		newIdx := enigma.LetterToIdx(byte(letter))
// 		newLtr := enigma.IdxToLetter(byte(idx))
// 		inverse[newIdx] = newLtr
// 	}

// 	return string(inverse)
// }

func fullTest(rotor *enigma.Rotor) {
	var offset byte = enigma.LetterToIdx(rotor.GetTopLetter())

	expectedFwd := rotor.Wiring()
	for idx, letter := range enigma.ALPHABET {
		outLetter := byte(expectedFwd[idx]) - offset
		Expect(rotor.Forward(byte(letter))).To(Equal(outLetter))
	}

	expectedRev := rotor.InverseWiring()
	for idx, letter := range enigma.ALPHABET {
		outLetter := byte(expectedRev[idx])
		Expect(rotor.Reverse(byte(letter))).To(Equal(outLetter))
	}
}

func testForward(rotor *enigma.Rotor, testCases []string) {
	offset := enigma.LetterToIdx(rotor.GetTopLetter())
	for _, test := range testCases {
		input := test[0]
		expected := test[1] - offset
		Expect(rotor.Forward(input)).To(Equal(expected))
	}
}

var _ = Describe("Rotor", func() {

	It("Rotor I: Input 'A' // Top Letter A-Z", func() {
		var offset byte
		var input byte = 'A'

		rotor := enigma.GetRotor("I")
		wiring := rotor.Wiring()

		for idx, letter := range enigma.ALPHABET {
			rotor.SetTopLetter(byte(letter))
			offset = enigma.LetterToIdx(rotor.GetTopLetter())

			output := wiring[idx] - offset
			Expect(rotor.Forward(input)).To(Equal(output))
		}
	})

	It("Rotor III: Top Letter 'B' // Forward", func() {
		rotor := enigma.GetRotor("III")
		rotor.SetTopLetter('B')

		// "In->Out"
		testCases := []string{
			"AD", "BF", "CH",
			"KV", "LZ", "MN",
			"XQ", "YO", "ZB",
		}
		testForward(rotor, testCases)
	})

	It("Rotor III: Top Letter 'Z' // Forward", func() {
		rotor := enigma.GetRotor("III")
		rotor.SetTopLetter('Z')

		// "In->Out"
		testCases := []string{
			"AO", "BB", "CD",
			"KT", "LX", "MV",
			"XU", "YS", "ZQ",
		}

		testForward(rotor, testCases)
	})

	// It("Should pass the Identity Test", func() {
	// 	rotor := enigma.GetRotor("0")
	// 	fullTest(rotor)
	// })

	// It("Can Set Top Letter", func() {
	// 	rotor := enigma.GetRotor("0")
	// 	var offset byte = 0

	// 	rotor.SetTopLetter('A')
	// 	Expect(rotor.Forward('A')).To(Equal(byte('A')))

	// 	rotor.SetTopLetter('B')
	// 	offset = enigma.LetterToIdx(rotor.GetTopLetter())
	// 	Expect(rotor.Forward('A')).To(Equal(byte('B') - offset))

	// 	rotor.SetTopLetter('C')
	// 	offset = enigma.LetterToIdx(rotor.GetTopLetter())
	// 	Expect(rotor.Forward('A')).To(Equal(byte('C') - offset))

	// 	rotor.SetTopLetter('J')
	// 	offset = enigma.LetterToIdx(rotor.GetTopLetter())
	// 	Expect(rotor.Forward('A')).To(Equal(byte('J') - offset))

	// 	rotor.SetTopLetter('Q')
	// 	offset = enigma.LetterToIdx(rotor.GetTopLetter())
	// 	Expect(rotor.Forward('A')).To(Equal(byte('Q') - offset))

	// 	rotor.SetTopLetter('X')
	// 	offset = enigma.LetterToIdx(rotor.GetTopLetter())
	// 	Expect(rotor.Forward('A')).To(Equal(byte('X') - offset))
	// })

	// It("Rotor I: Top Letter 'A'", func() {
	// 	rotor := enigma.GetRotor("I")
	// 	fullTest(rotor)
	// })

	// It("Rotor II: Top Letter 'A'", func() {
	// 	rotor := enigma.GetRotor("II")
	// 	fullTest(rotor)
	// })

	// It("Rotor III: Top Letter 'A'", func() {
	// 	rotor := enigma.GetRotor("III")
	// 	fullTest(rotor)
	// })

	// It("Rotor IV: Top Letter 'A'", func() {
	// 	rotor := enigma.GetRotor("IV")
	// 	fullTest(rotor)
	// })

	// It("Rotor V: Top Letter 'A'", func() {
	// 	rotor := enigma.GetRotor("V")
	// 	fullTest(rotor)
	// })

	// It("Rotor I: Top Letter 'B'", func() {
	// 	rotor := enigma.GetRotor("I")

	// 	rotor.SetTopLetter('B')
	// 	fullTest(rotor)
	// })

	// I:    ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// O:  BACDEFGHIJKLMNOPQRSTUVWXYZ
	// It("Top Letter Test", func() {
	// 	var offset byte

	// 	rotor := enigma.NewRotor("X", "BACDEFGHIJKLMNOPQRSTUVWXYZ", 'A', 'Z')

	// 	Expect(rotor.Forward('A')).To(Equal(byte('B')))
	// 	Expect(rotor.Reverse('A')).To(Equal(byte('B')))

	// 	rotor.SetTopLetter('B')
	// 	offset = enigma.LetterToIdx(rotor.GetTopLetter())
	// 	Expect(rotor.Forward('A')).To(Equal(byte('A') - offset))
	// 	Expect(rotor.Reverse('A')).To(Equal(byte('A')))

	// 	rotor.SetTopLetter('C')
	// 	offset = enigma.LetterToIdx(rotor.GetTopLetter())
	// 	Expect(rotor.Forward('A')).To(Equal(byte('C') - offset))
	// 	Expect(rotor.Reverse('A')).To(Equal(byte('Z')))
	// })

	// map_r_to_l rotor_III 'O' `14` should be `17`.
	// It("Rotor III: Top Letter 'O' // Forward // Pos 14 -> 17", func() {
	// 	// var offset byte
	// 	// ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// 	//               |
	// 	// YEIWGAKMUSQOBDFHJLCPRTXVZN
	// 	rotor := enigma.GetRotor("III")
	// 	rotor.SetTopLetter('O')
	// 	// offset = enigma.LetterToIdx(rotor.GetTopLetter())

	// 	input := byte('O')  // enigma.IdxToLetter(14)  // 'O'
	// 	output := byte('R') //enigma.IdxToLetter(17) // - offset //byte('F') - offset

	// 	Expect(rotor.Forward(input)).To(Equal(output))
	// })

	// `map_l_to_r rotor_I 'F' 10` should be `14`.
	// It("Rotor I: Top Letter 'F' // Reverse // Pos 10 -> 14", func() {
	// 	//          ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// 	//    EKMFLGDQVZNTOWYHXUSPAIBRCJ
	// 	rotor := enigma.GetRotor("I")
	// 	rotor.SetTopLetter('F')

	// 	input := enigma.IdxToLetter(10) // 'K'
	// 	output := byte('W')
	// 	Expect(rotor.Reverse(input)).To(Equal(output))
	// })

	// It("Rotor II: Top Letter 'K' // Wrap Test // F&R", func() {
	// 	rotor := enigma.GetRotor("II")
	// 	rotor.SetTopLetter('K')

	// 	fullTest(rotor)
	// 	// // Forward
	// 	// expectedFwd := rotor.Wiring()
	// 	// for idx, letter := range enigma.ALPHABET {
	// 	// 	outLetter := byte(expectedFwd[idx])
	// 	// 	Expect(rotor.Forward(byte(letter))).To(Equal(outLetter))
	// 	// }

	// 	// // Reverse - NO WRAP
	// 	// inLettersss := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// 	// expectedRev := "QZFSPMHBVRTAEJOKGWUDXNCYLI"
	// 	// // inLettersss := "AX"
	// 	// // expectedRev := "QY"
	// 	// for idx, letter := range inLettersss {
	// 	// 	if letter == '-' {
	// 	// 		continue
	// 	// 	}
	// 	// 	outLetter := byte(expectedRev[idx])
	// 	// 	Expect(rotor.Reverse(byte(letter))).To(Equal(outLetter))
	// 	// }
	// })

	// It("Can Step", func() {
	// 	rotor := enigma.GetRotor("I")

	// 	Expect(rotor.GetTopLetter()).To(Equal(byte('A')))

	// 	rotor.Step()
	// 	Expect(rotor.GetTopLetter()).To(Equal(byte('B')))

	// 	rotor.Step()
	// 	Expect(rotor.GetTopLetter()).To(Equal(byte('C')))

	// 	rotor.Step()
	// 	rotor.Step()
	// 	rotor.Step()
	// 	rotor.Step()
	// 	rotor.Step()
	// 	Expect(rotor.GetTopLetter()).To(Equal(byte('H')))
	// })

	// It("Can detect when it's at the notch", func() {
	// 	// Notch for II is E
	// 	rotor := enigma.GetRotor("II")

	// 	Expect(rotor.AtNotch()).To(BeFalse())

	// 	rotor.Step()
	// 	rotor.Step()
	// 	rotor.Step()
	// 	rotor.Step()
	// 	Expect(rotor.AtNotch()).To(BeTrue())

	// 	rotor.Step()
	// 	Expect(rotor.AtNotch()).To(BeFalse())
	// })

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
		rotorIII.Step()
		Expect(rotorIII.GetTopLetter()).To(Equal(byte('B')))
		offset := enigma.LetterToIdx(rotorIII.GetTopLetter())

		output = rotorIII.Forward(input)
		Expect(output).To(Equal(byte('D') - offset))

		output = rotorII.Forward(output)
		Expect(output).To(Equal(byte('D')))

		output = reflB.Reflect(output)
		Expect(output).To(Equal(byte('H')))

		output = rotorII.Reverse(output)
		Expect(output).To(Equal(byte('M') - offset))

		output = rotorIII.Reverse(output)
		Expect(output).To(Equal(byte('U')))

	})

})
