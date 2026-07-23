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
			Expect(wiring[inPos]).To(Equal(25 - inPos))
		}

		inPos := enigma.LetterToIdx('A')
		Expect(wiring[inPos]).To(Equal(enigma.LetterToIdx('Z')))

		inPos = enigma.LetterToIdx('B')
		Expect(wiring[inPos]).To(Equal(enigma.LetterToIdx('Y')))

		inPos = enigma.LetterToIdx('P')
		Expect(wiring[inPos]).To(Equal(enigma.LetterToIdx('K')))
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

		// Specify("Top Letter 'A'", func() {
		// 	Expect(rotor.GetTopLetter()).To(Equal('A'))

		// 	// By("Position")
		// 	Expect(rotor.Forward(0)).To(Equal(1))

		// 	// By("Converted Position (LetterToIdx)")
		// 	inPos := enigma.LetterToIdx('A')
		// 	outPos := enigma.LetterToIdx('B')
		// 	Expect(rotor.Forward(inPos)).To(Equal(outPos))

		// 	// By("Letter")
		// 	Expect(rotor.RightToLeft('A')).To(Equal('B'))
		// })

		// Specify("Top Letter 'B'", func() {
		// 	rotor.SetTopLetter('B')
		// 	Expect(rotor.GetTopLetter()).To(Equal('B'))

		// 	// By Position
		// 	Expect(rotor.Forward(0)).To(Equal(3))

		// 	// By Converted Letter
		// 	inPos := enigma.LetterToIdx('A')
		// 	outPos := enigma.LetterToIdx('D')
		// 	Expect(rotor.Forward(inPos)).To(Equal(outPos))

		// 	// By Letter
		// 	Expect(rotor.RightToLeft('A')).To(Equal('D'))
		// 	Expect(rotor.RightToLeft('B')).To(Equal('F'))
		// })

		// Specify("Top Letter 'X'", func() {
		// 	rotor.SetTopLetter('X')
		// 	Expect(rotor.GetTopLetter()).To(Equal('X'))

		// 	Expect(rotor.RightToLeft('A')).To(Equal('S'))
		// 	Expect(rotor.RightToLeft('X')).To(Equal('K'))
		// })

		Specify("CRAIG", func() {
			rotor.SetTopLetter('B')
			Expect(rotor.GetTopLetter()).To(Equal('B'))

			// By Position
			outIdx, outLtr := rotor.Fwd(0)
			Expect(outIdx).To(Equal(2), "CRAIG: 0 -> 2")
			Expect(outLtr).To(Equal('D'), "CRAIG: 0 -> D")

			// outPos := rotor.Fwd2(0)
			// Expect(outPos).To(Equal(enigma.RotorPosition{2, 'D'}))

			// // By Converted Letter
			// inIdx := enigma.LetterToIdx('A')
			// // outIdx := enigma.LetterToIdx('D')
			// outPos = rotor.Fwd2(inIdx)
			// Expect(outPos).To(Equal(enigma.RotorPosition{2, 'D'}))

			// // By Letter
			// Expect(rotor.RightToLeft('A')).To(Equal('D'))
			// Expect(rotor.RightToLeft('B')).To(Equal('F'))
		})

	})

	// Context("Rotor III // Reverse", func() {
	// 	rotor := enigma.GetRotor("III")

	// 	Specify("Top Letter 'A'", func() {
	// 		Expect(rotor.GetTopLetter()).To(Equal('A'))

	// 		// By Position
	// 		Expect(rotor.Reverse(1)).To(Equal(0), "Position: 1 -> 0")

	// 		// By Converted Letter
	// 		inPos := enigma.LetterToIdx('B')
	// 		outPos := enigma.LetterToIdx('A')
	// 		Expect(rotor.Reverse(inPos)).To(Equal(outPos), "LetterToIdx: B -> A")

	// 		// By Letter
	// 		for _, set := range []string{"BA", "AT", "CG", "ZM"} {
	// 			inLetter := rune(set[0])
	// 			outLetter := rune(set[1])
	// 			Expect(rotor.LeftToRight(inLetter)).To(Equal(outLetter), "Letter: %c -> %c", inLetter, outLetter)
	// 		}
	// 	})

	// 	Specify("Top Letter 'B'", func() {
	// 		rotor.SetTopLetter('B')
	// 		Expect(rotor.GetTopLetter()).To(Equal('B'))

	// 		// By Position
	// 		Expect(rotor.Reverse(1)).To(Equal(25), "Position: 1 -> 25")

	// 		// By Converted Letter
	// 		inPos := enigma.LetterToIdx('B')
	// 		outPos := enigma.LetterToIdx('Z')
	// 		Expect(rotor.Reverse(inPos)).To(Equal(outPos), "LetterToIdx: B -> Z")

	// 		// By Letter
	// 		for _, set := range []string{"AS", "BZ", "CF", "OY"} {
	// 			inLetter := rune(set[0])
	// 			outLetter := rune(set[1])
	// 			Expect(rotor.LeftToRight(inLetter)).To(Equal(outLetter), "Letter: %c -> %c", inLetter, outLetter)
	// 		}
	// 	})

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
	// })
})
