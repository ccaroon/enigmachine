package enigma_test

import (
	"github.com/ccaroon/enigmachine/enigma"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rotor", func() {

	It("Can create a new Rotor", func() {
		rotor := enigma.NewRotor(
			"X",
			"ZYXWVUTSRQPONMLKJIHGFEDCBA",
			'A',
			'A',
		)

		Expect(rotor.Id()).To(Equal("X"))
		Expect(rotor.GetTopLetter()).To(Equal('A'))

		wiring := rotor.Wiring()
		for _, letter := range enigma.ALPHABET {
			inPos := enigma.LetterToIdx(letter)
			// 25 => 25 - inPos
			Expect(rune(wiring[inPos])).To(Equal(enigma.IdxToLetter(25 - inPos)))
		}

		inPos := enigma.LetterToIdx('A')
		Expect(rune(wiring[inPos])).To(Equal('Z'))

		inPos = enigma.LetterToIdx('B')
		Expect(rune(wiring[inPos])).To(Equal('Y'))

		inPos = enigma.LetterToIdx('P')
		Expect(rune(wiring[inPos])).To(Equal('K'))
	})

	It("Can get a rotor by name", func() {
		for _, id := range []string{"I", "II", "III", "IV", "V"} {
			rotor := enigma.GetRotor(id)

			Expect(rotor.Id()).To(Equal(id))
			Expect(rotor.GetTopLetter()).To(Equal('A'))
		}

		// Non-existent rotor
		rotor := enigma.GetRotor("42")
		Expect(rotor).To(BeNil())
	})

	It("Can step", func() {
		rotor := enigma.GetRotor("I")

		Expect(rotor.GetTopLetter()).To(Equal('A'))

		rotor.Step()
		Expect(rotor.GetTopLetter()).To(Equal('B'))

		rotor.Step()
		Expect(rotor.GetTopLetter()).To(Equal('C'))

		rotor.Step()
		rotor.Step()
		rotor.Step()
		Expect(rotor.GetTopLetter()).To(Equal('F'))
	})

	It("Can determine when it's at the turnover notch", func() {
		rotor := enigma.GetRotor("II") // Notch at 'E'

		Expect(rotor.GetTopLetter()).To(Equal('A'))
		Expect(rotor.AtNotch()).To(BeFalse())

		rotor.Step()
		Expect(rotor.GetTopLetter()).To(Equal('B'))
		Expect(rotor.AtNotch()).To(BeFalse())

		rotor.Step()
		Expect(rotor.GetTopLetter()).To(Equal('C'))
		Expect(rotor.AtNotch()).To(BeFalse())

		rotor.Step()
		Expect(rotor.GetTopLetter()).To(Equal('D'))
		Expect(rotor.AtNotch()).To(BeFalse())

		rotor.Step()
		Expect(rotor.GetTopLetter()).To(Equal('E'))
		Expect(rotor.AtNotch()).To(BeTrue())
	})

	Context("Rotor III // Forward", func() {
		rotor := enigma.GetRotor("III")

		Specify("Top Letter 'A'", func() {
			offset := enigma.LetterToIdx(rotor.GetTopLetter())
			Expect(rotor.GetTopLetter()).To(Equal('A'))

			for _, test := range []string{"AB", "BD", "CF", "XS", "YQ", "ZO"} {
				inLtr := rune(test[0])
				expLtr := rune(test[1])
				expIdx := enigma.LetterToIdx(expLtr) - offset

				inIdx := enigma.LetterToIdx(inLtr)

				outIdx, outLtr := rotor.Forward(inIdx)
				Expect(outLtr).To(Equal(expLtr), "Ltr(%s): %c -> %c / %d", test, inLtr, expLtr, expIdx)
				Expect(outIdx).To(Equal(expIdx), "Idx(%s): %c -> %c / %d", test, inLtr, expLtr, expIdx)
			}
		})

		Specify("Top Letter 'B'", func() {
			rotor.SetTopLetter('B')
			offset := enigma.LetterToIdx(rotor.GetTopLetter())
			Expect(rotor.GetTopLetter()).To(Equal('B'))

			for _, test := range []string{"AD", "BF", "CH", "XQ", "YO", "ZB"} {
				inLtr := rune(test[0])
				expLtr := rune(test[1])
				expIdx := enigma.LetterToIdx(expLtr) - offset

				inIdx := enigma.LetterToIdx(inLtr)

				outIdx, outLtr := rotor.Forward(inIdx)
				Expect(outLtr).To(Equal(expLtr), "Ltr(%s): %c -> %c / %d", test, inLtr, expLtr, expIdx)
				Expect(outIdx).To(Equal(expIdx), "Idx(%s): %c -> %c / %d", test, inLtr, expLtr, expIdx)
			}
		})

		Specify("Top Letter 'X'", func() {
			rotor.SetTopLetter('X')
			offset := enigma.LetterToIdx(rotor.GetTopLetter())
			Expect(rotor.GetTopLetter()).To(Equal('X'))

			for _, test := range []string{"AS", "BQ", "CO", "XK", "YM", "ZU"} {
				inLtr := rune(test[0])
				expLtr := rune(test[1])
				expIdx := enigma.LetterToIdx(expLtr) - offset

				inIdx := enigma.LetterToIdx(inLtr)

				outIdx, outLtr := rotor.Forward(inIdx)
				Expect(outLtr).To(Equal(expLtr), "Ltr(%s): %c -> %c / %d", test, inLtr, expLtr, expIdx)
				Expect(outIdx).To(Equal(expIdx), "Idx(%s): %c -> %c / %d", test, inLtr, expLtr, expIdx)
			}
		})

		Specify("XRAIG", func() {
			rotor.SetTopLetter('B')
			Expect(rotor.GetTopLetter()).To(Equal('B'))

			outIdx, outLtr := rotor.Forward(enigma.LetterToIdx('A'))
			Expect(outIdx).To(Equal(2))
			Expect(outLtr).To(Equal('D'))

			outIdx, outLtr = rotor.Forward(enigma.LetterToIdx('B'))
			Expect(outIdx).To(Equal(4))
			Expect(outLtr).To(Equal('F'))
		})

	})

	Context("Rotor III // Reverse", func() {
		rotor := enigma.GetRotor("III")

		Specify("Top Letter 'A'", func() {
			offset := enigma.LetterToIdx(rotor.GetTopLetter())
			Expect(rotor.GetTopLetter()).To(Equal('A'))

			for _, test := range []string{"AT", "BA", "CG", "XK", "YO", "ZM"} {
				inLtr := rune(test[0])
				expLtr := rune(test[1])
				expIdx := enigma.LetterToIdx(expLtr) - offset

				inIdx := enigma.LetterToIdx(inLtr)

				outIdx, outLtr := rotor.Reverse(inIdx)
				Expect(outLtr).To(Equal(expLtr), "Ltr(%s): %c -> %c / %d", test, inLtr, expLtr, expIdx)
				Expect(outIdx).To(Equal(expIdx), "Idx(%s): %c -> %c / %d", test, inLtr, expLtr, expIdx)
			}
		})

		Specify("XRAIG Top Letter 'B'", func() {
			rotor.SetTopLetter('B')
			offset := enigma.LetterToIdx(rotor.GetTopLetter())
			Expect(rotor.GetTopLetter()).To(Equal('B'))

			// "AS", "BZ", "CF", "XJ", "YN", "ZL"
			for _, test := range []string{"AS", "BZ", "CF", "XJ", "YN", "ZL"} {
				inLtr := rune(test[0])
				expLtr := rune(test[1])
				expIdx := enigma.LetterToIdx(expLtr)

				inIdx := enigma.LetterToIdx(inLtr) - offset
				// fmt.Printf("\ninIdx[%d]\n", inIdx)

				outIdx, outLtr := rotor.Reverse(inIdx)
				Expect(outLtr).To(Equal(expLtr), "Ltr(%s|%d): %c -> [%d|%c] [%c]", test, inIdx, inLtr, expIdx, expLtr, outLtr)
				Expect(outIdx).To(Equal(expIdx), "Idx(%s|%d): %c -> [%d|%c] [%c]", test, inIdx, inLtr, expIdx, expLtr, outLtr)
			}

			// Wrap
			// outIdx, outLtr := rotor.Reverse(0)
			// Expect(outLtr).To(Equal('Z'), "Letter: 0 -> 25|Z")
			// Expect(outIdx).To(Equal(25), "Index: 0 -> 25|Z")

			outIdx, outLtr := rotor.Reverse(25)
			Expect(outLtr).To(Equal('S'), "Letter: 0 -> 18|S")
			Expect(outIdx).To(Equal(18), "Index: 0 -> 18|S")

			// // No Wrap
			// outIdx, outLtr = rotor.Reverse(1)
			// Expect(outLtr).To(Equal('F'), "Letter: 1 -> 5|F")
			// Expect(outIdx).To(Equal(5), "Index: 1 -> 5|F")

			// outIdx, outLtr = rotor.Reverse(2)
			// Expect(outLtr).To(Equal('A'), "Letter: 2 -> 0|A")
			// Expect(outIdx).To(Equal(0), "Index: 2 -> 0|A")

			// outIdx, outLtr = rotor.Reverse(3)
			// Expect(outLtr).To(Equal('O'), "Letter: 3 -> 14|O")
			// Expect(outIdx).To(Equal(14), "Index: 3 -> 14|O")
		})

		// 	Specify("Top Letter 'V'", func() {
		// 		rotor.SetTopLetter('V')
		// 		Expect(rotor.GetTopLetter()).To(Equal('V'))

		// 		// By Letter
		// 		for _, set := range []string{"AY", "BF", "SC", "VQ", "ZR"} {
		// 			inLetter := rune(set[0])
		// 			outLetter := rune(set[1])
		// 			Expect(rotor.LeftToRight(inLetter)).To(Equal(outLetter), "Letter: %c -> %c", inLetter, outLetter)
		// 		}
		// 	})
		// })

		// Context("Rotor Chaining", func() {
		// 	rotorI := enigma.GetRotor("I")
		// 	rotorII := enigma.GetRotor("II")
		// 	rotorIII := enigma.GetRotor("III")
		// 	Specify("Forward(AAA): III -> II -> I", func() {
		// 		Expect(rotorIII.GetTopLetter()).To(Equal('A'))
		// 		Expect(rotorII.GetTopLetter()).To(Equal('A'))
		// 		Expect(rotorI.GetTopLetter()).To(Equal('A'))

		// 		for _, path := range []string{"ABJZ", "BDKN", "XSZJ", "QIXR"} {
		// 			inLtr := rune(path[0])
		// 			expOut := rune(path[1])

		// 			outLtr := rotorIII.RightToLeft(inLtr)
		// 			Expect(outLtr).To(Equal(expOut), "RotorIII(A): %c -> %c", inLtr, expOut)

		// 			inLtr = outLtr
		// 			expOut = rune(path[2])
		// 			outLtr = rotorII.RightToLeft(inLtr)
		// 			Expect(outLtr).To(Equal(expOut), "RotorII(A): %c -> %c", inLtr, expOut)

		// 			inLtr = outLtr
		// 			expOut = rune(path[3])
		// 			outLtr = rotorI.RightToLeft(inLtr)
		// 			Expect(outLtr).To(Equal(expOut), "RotorI(A): %c -> %c", inLtr, expOut)
		// 		}
		// 	})

		// 	Specify("Forward(BAA): III -> II -> I", func() {
		// 		rotorIII.SetTopLetter('B')

		// 		Expect(rotorIII.GetTopLetter()).To(Equal('B'))
		// 		Expect(rotorII.GetTopLetter()).To(Equal('A'))
		// 		Expect(rotorI.GetTopLetter()).To(Equal('A'))

		// 		for _, path := range []string{"ADDF"} {
		// 			inLtr := rune(path[0])
		// 			expOut := rune(path[1])

		// 			outLtr := rotorIII.RightToLeft(inLtr)
		// 			Expect(outLtr).To(Equal(expOut), "RotorIII(B): %c -> %c", inLtr, expOut)

		// 			inLtr = outLtr - 1
		// 			expOut = rune(path[2])
		// 			outLtr = rotorII.RightToLeft(inLtr)
		// 			Expect(outLtr).To(Equal(expOut), "RotorII(A): %c -> %c", inLtr, expOut)
		// 			fmt.Printf("\nRotorII(A): %c -> %c\n", inLtr, expOut)

		// 			inLtr = outLtr
		// 			expOut = rune(path[3])
		// 			outLtr = rotorI.RightToLeft(inLtr)
		// 			Expect(outLtr).To(Equal(expOut), "RotorI(A): %c -> %c", inLtr, expOut)
		// 		}
		// 	})
	})
})
