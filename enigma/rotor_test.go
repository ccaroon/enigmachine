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
				Expect(outIdx).To(Equal(expIdx), "Idx(%s): %c -> [%d|%c] [%d|%c]", test, inLtr, expIdx, expLtr, outIdx, outLtr)
				Expect(outLtr).To(Equal(expLtr), "Ltr(%s): %c -> [%d|%c] [%d|%c]", test, inLtr, expIdx, expLtr, outIdx, outLtr)
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
				if expIdx < 0 {
					expIdx = 26 + expIdx
				}

				inIdx := enigma.LetterToIdx(inLtr)

				outIdx, outLtr := rotor.Forward(inIdx)
				Expect(outIdx).To(Equal(expIdx), "Idx(%s): %c -> [%d|%c] [%d|%c]", test, inLtr, expIdx, expLtr, outIdx, outLtr)
				Expect(outLtr).To(Equal(expLtr), "Ltr(%s): %c -> [%d|%c] [%d|%c]", test, inLtr, expIdx, expLtr, outIdx, outLtr)
			}
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

		Specify("Top Letter 'B'", func() {
			rotor.SetTopLetter('B')
			offset := enigma.LetterToIdx(rotor.GetTopLetter())
			Expect(rotor.GetTopLetter()).To(Equal('B'))

			for _, test := range []string{"AS", "BZ", "CF", "XJ", "YN", "ZL"} {
				inLtr := rune(test[0])
				expLtr := rune(test[1])
				expIdx := enigma.LetterToIdx(expLtr)

				inIdx := enigma.LetterToIdx(inLtr) - offset

				outIdx, outLtr := rotor.Reverse(inIdx)
				Expect(outLtr).To(Equal(expLtr), "Ltr(%s|%d): %c -> [%d|%c] [%c]", test, inIdx, inLtr, expIdx, expLtr, outLtr)
				Expect(outIdx).To(Equal(expIdx), "Idx(%s|%d): %c -> [%d|%c] [%c]", test, inIdx, inLtr, expIdx, expLtr, outLtr)
			}

			// Wrap
			outIdx, outLtr := rotor.Reverse(0)
			Expect(outLtr).To(Equal('Z'), "Letter: 0 -> 25|Z")
			Expect(outIdx).To(Equal(25), "Index: 0 -> 25|Z")

			outIdx, outLtr = rotor.Reverse(25)
			Expect(outLtr).To(Equal('S'), "Letter: 0 -> 18|S")
			Expect(outIdx).To(Equal(18), "Index: 0 -> 18|S")
		})

		Specify("Top Letter 'V'", func() {
			rotor.SetTopLetter('V')
			offset := enigma.LetterToIdx(rotor.GetTopLetter())
			Expect(rotor.GetTopLetter()).To(Equal('V'))

			for _, test := range []string{"AY", "BF", "CL", "XP", "YT", "ZR"} {
				inLtr := rune(test[0])
				expLtr := rune(test[1])
				expIdx := enigma.LetterToIdx(expLtr)

				inIdx := enigma.LetterToIdx(inLtr) - offset

				outIdx, outLtr := rotor.Reverse(inIdx)
				Expect(outLtr).To(Equal(expLtr), "Ltr(%s|%d): %c -> [%d|%c] [%c]", test, inIdx, inLtr, expIdx, expLtr, outLtr)
				Expect(outIdx).To(Equal(expIdx), "Idx(%s|%d): %c -> [%d|%c] [%c]", test, inIdx, inLtr, expIdx, expLtr, outLtr)
			}

			outIdx, outLtr := rotor.Reverse(0)
			Expect(outLtr).To(Equal('Q'))
			Expect(outIdx).To(Equal(16))

			outIdx, outLtr = rotor.Reverse(13)
			Expect(outLtr).To(Equal('V'))
			Expect(outIdx).To(Equal(21))

		})

		Context("Rotor Chaining", func() {
			rotorI := enigma.GetRotor("I")
			rotorII := enigma.GetRotor("II")
			rotorIII := enigma.GetRotor("III")
			Specify("Forward(AAA): III -> II -> I", func() {
				Expect(rotorIII.GetTopLetter()).To(Equal('A'))
				Expect(rotorII.GetTopLetter()).To(Equal('A'))
				Expect(rotorI.GetTopLetter()).To(Equal('A'))

				for _, path := range []string{"ABJZ", "BDKN", "XSZJ", "QIXR"} {
					inLtr := rune(path[0])
					inIdx := enigma.LetterToIdx(inLtr)
					expOut := rune(path[1])

					outIdx, outLtr := rotorIII.Forward(inIdx)
					Expect(outLtr).To(Equal(expOut), "RotorIII(A): %c -> %c", inLtr, expOut)

					inIdx = outIdx
					expOut = rune(path[2])
					outIdx, outLtr = rotorII.Forward(inIdx)
					Expect(outLtr).To(Equal(expOut), "RotorII(A): %c -> %c", inLtr, expOut)

					inIdx = outIdx
					expOut = rune(path[3])
					outIdx, outLtr = rotorI.Forward(inIdx)
					Expect(outLtr).To(Equal(expOut), "RotorI(A): %c -> %c", inLtr, expOut)
				}
			})

			Specify("Forward(BAA): III -> II -> I", func() {
				rotorIII.SetTopLetter('B')

				Expect(rotorIII.GetTopLetter()).To(Equal('B'))
				Expect(rotorII.GetTopLetter()).To(Equal('A'))
				Expect(rotorI.GetTopLetter()).To(Equal('A'))

				for _, path := range []string{"ADDF", "BFSS", "XQCM", "MNWB", "ZBAE"} {
					inLtr := rune(path[0])
					inIdx := enigma.LetterToIdx(inLtr)
					expOut := rune(path[1])

					outIdx, outLtr := rotorIII.Forward(inIdx)
					Expect(outLtr).To(Equal(expOut), "RotorIII(B): %c -> %c", inLtr, expOut)

					inIdx = outIdx
					expOut = rune(path[2])
					outIdx, outLtr = rotorII.Forward(inIdx)
					Expect(outLtr).To(Equal(expOut), "RotorII(A): %c -> %c", inLtr, expOut)

					inIdx = outIdx
					expOut = rune(path[3])
					outIdx, outLtr = rotorI.Forward(inIdx)
					Expect(outLtr).To(Equal(expOut), "RotorI(A): %c -> %c", inLtr, expOut)
				}
			})
		})
	})
})
