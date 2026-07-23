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

		Specify("Top Letter 'A'", func() {
			Expect(rotor.GetTopLetter()).To(Equal('A'))

			// By("Position")
			Expect(rotor.Forward(0)).To(Equal(1))

			// By("Converted Position (LetterToIdx)")
			inPos := enigma.LetterToIdx('A')
			outPos := enigma.LetterToIdx('B')
			Expect(rotor.Forward(inPos)).To(Equal(outPos))

			// By("Letter")
			Expect(rotor.RightToLeft('A')).To(Equal('B'))
		})

		Specify("Top Letter 'B'", func() {
			rotor.SetTopLetter('B')
			Expect(rotor.GetTopLetter()).To(Equal('B'))

			// By Position
			Expect(rotor.Forward(0)).To(Equal(3))

			// By Converted Letter
			inPos := enigma.LetterToIdx('A')
			outPos := enigma.LetterToIdx('D')
			Expect(rotor.Forward(inPos)).To(Equal(outPos))

			// By Letter
			Expect(rotor.RightToLeft('A')).To(Equal('D'))
			Expect(rotor.RightToLeft('B')).To(Equal('F'))
		})

		Specify("Top Letter 'X'", func() {
			rotor.SetTopLetter('X')
			Expect(rotor.GetTopLetter()).To(Equal('X'))

			Expect(rotor.RightToLeft('A')).To(Equal('S'))
			Expect(rotor.RightToLeft('X')).To(Equal('K'))
		})
	})

	Context("Rotor III // Reverse", func() {
		rotor := enigma.GetRotor("III")

		Specify("Top Letter 'A'", func() {
			Expect(rotor.GetTopLetter()).To(Equal('A'))

			// By Position
			Expect(rotor.Reverse(1)).To(Equal(0), "Position: 1 -> 0")

			// By Converted Letter
			inPos := enigma.LetterToIdx('B')
			outPos := enigma.LetterToIdx('A')
			Expect(rotor.Reverse(inPos)).To(Equal(outPos), "LetterToIdx: B -> A")

			// By Letter
			for _, set := range []string{"BA", "AT", "CG", "ZM"} {
				inLetter := rune(set[0])
				outLetter := rune(set[1])
				Expect(rotor.LeftToRight(inLetter)).To(Equal(outLetter), "Letter: %c -> %c", inLetter, outLetter)
			}
		})

		Specify("Top Letter 'B'", func() {
			rotor.SetTopLetter('B')
			Expect(rotor.GetTopLetter()).To(Equal('B'))

			// By Position
			Expect(rotor.Reverse(1)).To(Equal(25), "Position: 1 -> 25")

			// By Converted Letter
			inPos := enigma.LetterToIdx('B')
			outPos := enigma.LetterToIdx('Z')
			Expect(rotor.Reverse(inPos)).To(Equal(outPos), "LetterToIdx: B -> Z")

			// By Letter
			for _, set := range []string{"AS", "BZ", "CF", "OY"} {
				inLetter := rune(set[0])
				outLetter := rune(set[1])
				Expect(rotor.LeftToRight(inLetter)).To(Equal(outLetter), "Letter: %c -> %c", inLetter, outLetter)
			}
		})

		Specify("Top Letter 'V'", func() {
			rotor.SetTopLetter('V')
			Expect(rotor.GetTopLetter()).To(Equal('V'))

			// By Letter
			for _, set := range []string{"AY", "BF", "SC", "VQ", "ZR"} {
				inLetter := rune(set[0])
				outLetter := rune(set[1])
				Expect(rotor.LeftToRight(inLetter)).To(Equal(outLetter), "Letter: %c -> %c", inLetter, outLetter)
			}
		})
	})
})
