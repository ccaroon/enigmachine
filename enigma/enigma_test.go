package enigma_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ccaroon/enigmachine/enigma"
)

func testRotorConfig(machine *enigma.Enigma, setting string) {
	Expect(machine.GetRotorByIdx(0).GetTopLetter()).To(Equal(byte(setting[0])))
	Expect(machine.GetRotorByIdx(1).GetTopLetter()).To(Equal(byte(setting[1])))
	Expect(machine.GetRotorByIdx(2).GetTopLetter()).To(Equal(byte(setting[2])))
}

var _ = Describe("Enigma", func() {

	It("Can get a Rotor by Id", func() {
		machine := enigma.NewEnigma(
			"B",
			[]string{"I", "II", "III"},
			[]byte{},
		)

		Expect(machine).ToNot(BeNil())

		rotor := machine.GetRotor("I")
		Expect(rotor.Id()).To(Equal("I"))

		rotor = machine.GetRotor("II")
		Expect(rotor.Id()).To(Equal("II"))

		rotor = machine.GetRotor("III")
		Expect(rotor.Id()).To(Equal("III"))

		rotor = machine.GetRotor("X")
		Expect(rotor).To(BeNil())
	})

	It("Can Get Rotor by Index", func() {
		machine := enigma.NewEnigma(
			"B",
			[]string{"I", "II", "III"},
			[]byte{},
		)

		Expect(machine).ToNot(BeNil())

		Expect(machine.GetRotorByIdx(0).Id()).To(Equal("I"))
		Expect(machine.GetRotorByIdx(1).Id()).To(Equal("II"))
		Expect(machine.GetRotorByIdx(2).Id()).To(Equal("III"))
	})

	It("Can Pass the Identity Test", func() {
		machine := enigma.NewEnigma(
			"I", // Identity Reflector
			[]string{},
			[]byte{},
		)
		Expect(machine).ToNot(BeNil())

		outLetter := machine.EncipherLetter('A')
		Expect(outLetter).To(Equal(byte('A')))

		outLetter = machine.EncipherLetter('M')
		Expect(outLetter).To(Equal(byte('M')))

		outLetter = machine.EncipherLetter('X')
		Expect(outLetter).To(Equal(byte('X')))

		outLetter = machine.EncipherLetter('Z')
		Expect(outLetter).To(Equal(byte('Z')))
	})

	It("Can set starting config for Rotors", func() {
		machine := enigma.NewEnigma(
			"B",
			[]string{"III", "II", "I"},
			[]byte{},
		)

		machine.ConfigureRotors("ABC")
		testRotorConfig(machine, "ABC")

		machine.ConfigureRotors("ZYX")
		testRotorConfig(machine, "ZYX")
	})

	It("Can properly step multiple rotors", func() {
		machine := enigma.NewEnigma(
			"B",
			[]string{"III", "II", "I"},
			[]byte{},
		)

		machine.ConfigureRotors("KDO")
		expected := []string{"KDO", "KDP", "KDQ", "KER", "LFS", "LFT", "LFU"}
		for _, seq := range expected {
			testRotorConfig(machine, seq)
			machine.Step()
		}

		machine.ConfigureRotors("VDP")
		expected = []string{"VDP", "VDQ", "VER", "WFS", "WFT"}
		for _, seq := range expected {
			testRotorConfig(machine, seq)
			machine.Step()
		}
	})

	It("Can encipher a single letter", func() {
		machine := enigma.NewEnigma(
			"B",
			[]string{"I", "II", "III"},
			[]byte{},
		)

		Expect(machine).ToNot(BeNil())

		// A -> U
		outLetter := machine.EncipherLetter('A')
		Expect(outLetter).To(Equal(byte('U')))
		// reciprocal
		outLetter = machine.EncipherLetter('U')
		Expect(outLetter).To(Equal(byte('A')))

		// G -> P
		outLetter = machine.EncipherLetter('G')
		Expect(outLetter).To(Equal(byte('P')))
		// reciprocal
		outLetter = machine.EncipherLetter('P')
		Expect(outLetter).To(Equal(byte('G')))
	})

	It("Can encipher a single letter with stepping", func() {
		machine := enigma.NewEnigma(
			"B",
			[]string{"I", "II", "III"},
			[]byte{},
		)

		Expect(machine).ToNot(BeNil())

		// Step/Rotate III
		rotorIII := machine.GetRotor("III")
		Expect(rotorIII.Id()).To(Equal("III"))
		rotorIII.Step()

		outLetter := machine.EncipherLetter('A')
		Expect(outLetter).To(Equal(byte('B')))
		// reciprocal
		outLetter = machine.EncipherLetter('B')
		Expect(outLetter).To(Equal(byte('A')))

		outLetter = machine.EncipherLetter('G')
		Expect(outLetter).To(Equal(byte('X')))
		// reciprocal
		outLetter = machine.EncipherLetter('X')
		Expect(outLetter).To(Equal(byte('G')))
	})

	It("Can encipher a string", func() {
		machine := enigma.NewEnigma(
			"B",
			[]string{"I", "II", "III"},
			[]byte{'A', 'Z'},
		)
		machine.ConfigureRotors("FUN")

		Expect(machine.EncipherString("YNGXQ")).To(Equal("OCAML"))
	})

})
