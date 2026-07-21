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

	// [1 3 5 7 9 11 2 15 17 19 23 21 25 13 24 4 8 22 6 0 10 12 20 18 16 14]
	It("Rotor III // Forward // Top Letter 'A'", func() {
		rotor := enigma.GetRotor("III")

		Expect(rotor.GetTopLetter()).To(Equal('A'))

		// by position
		Expect(rotor.Forward(0)).To(Equal(1))

		// by converted position
		inPos := enigma.LetterToIdx('A')
		outPos := enigma.LetterToIdx('B')
		Expect(rotor.Forward(inPos)).To(Equal(outPos))

		// by letter
		Expect(rotor.RightToLeft('A')).To(Equal('B'))
	})

	// [1 3 5 7 9 11 2 15 17 19 23 21 25 13 24 4 8 22 6 0 10 12 20 18 16 14]
	It("Rotor III // Forward // Top Letter 'B'", func() {
		rotor := enigma.GetRotor("III")
		rotor.SetTopLetter('B')
		Expect(rotor.GetTopLetter()).To(Equal('B'))

		// by position
		Expect(rotor.Forward(0)).To(Equal(3))

		// by converted position
		inPos := enigma.LetterToIdx('A')
		outPos := enigma.LetterToIdx('D')
		Expect(rotor.Forward(inPos)).To(Equal(outPos))

		// by letter
		Expect(rotor.RightToLeft('A')).To(Equal('D'))
	})

})
