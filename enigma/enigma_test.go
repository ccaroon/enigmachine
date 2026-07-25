package enigma_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ccaroon/enigmachine/enigma"
)

func testRotorConfig(machine *enigma.Enigma, setting string) {
	Expect(machine.GetRotorByIdx(0).GetTopLetter()).To(Equal(rune(setting[0])))
	Expect(machine.GetRotorByIdx(1).GetTopLetter()).To(Equal(rune(setting[1])))
	Expect(machine.GetRotorByIdx(2).GetTopLetter()).To(Equal(rune(setting[2])))
}

var _ = Describe("Enigma", func() {

	It("Can get a Rotor by Id", func() {
		machine := enigma.NewEnigma(
			"B",
			[]string{"I", "II", "III"},
			[]rune{},
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
			[]rune{},
		)

		Expect(machine).ToNot(BeNil())

		Expect(machine.GetRotorByIdx(0).Id()).To(Equal("I"))
		Expect(machine.GetRotorByIdx(1).Id()).To(Equal("II"))
		Expect(machine.GetRotorByIdx(2).Id()).To(Equal("III"))
	})

	It("Can set starting config for Rotors", func() {
		machine := enigma.NewEnigma(
			"B",
			[]string{"III", "II", "I"},
			[]rune{},
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
			[]rune{},
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

	Context("Enciphering", func() {
		Specify("Single Letter: B | I,II,III | AAA | []", func() {
			machine := enigma.NewEnigma(
				"B",
				[]string{"I", "II", "III"},
				[]rune{},
			)

			Expect(machine).ToNot(BeNil())

			// A -> U
			outLetter := machine.EncipherLetter('A')
			Expect(outLetter).To(Equal('U'))
			// reciprocal
			outLetter = machine.EncipherLetter('U')
			Expect(outLetter).To(Equal('A'))

			// G -> P
			outLetter = machine.EncipherLetter('G')
			Expect(outLetter).To(Equal('P'))
			// reciprocal
			outLetter = machine.EncipherLetter('P')
			Expect(outLetter).To(Equal('G'))
		})

		Specify("Single Letter: B | I,II,III | AAB | []", func() {
			machine := enigma.NewEnigma(
				"B",
				[]string{"I", "II", "III"},
				[]rune{},
			)

			Expect(machine).ToNot(BeNil())

			machine.ConfigureRotors("AAB")

			// A -> B
			outLetter := machine.EncipherLetter('A')
			Expect(outLetter).To(Equal('B'), "%c -> %c", 'A', outLetter)
			// // reciprocal
			outLetter = machine.EncipherLetter('B')
			Expect(outLetter).To(Equal('A'))

			// G -> P
			outLetter = machine.EncipherLetter('G')
			Expect(outLetter).To(Equal('X'))
			// reciprocal
			outLetter = machine.EncipherLetter('X')
			Expect(outLetter).To(Equal('G'))
		})

		// It("Can encipher a single letter with stepping", func() {
		// 	machine := enigma.NewEnigma(
		// 		"B",
		// 		[]string{"I", "II", "III"},
		// 		[]rune{},
		// 	)

		// 	Expect(machine).ToNot(BeNil())

		// 	// Step/Rotate III
		// 	rotorIII := machine.GetRotor("III")
		// 	Expect(rotorIII.Id()).To(Equal("III"))
		// 	rotorIII.Step()

		// 	outLetter := machine.EncipherLetter('A')
		// 	Expect(outLetter).To(Equal(byte('B')))
		// 	// reciprocal
		// 	outLetter = machine.EncipherLetter('B')
		// 	Expect(outLetter).To(Equal(byte('A')))

		// 	outLetter = machine.EncipherLetter('G')
		// 	Expect(outLetter).To(Equal(byte('X')))
		// 	// reciprocal
		// 	outLetter = machine.EncipherLetter('X')
		// 	Expect(outLetter).To(Equal(byte('G')))
		// })

		// It("Can encipher FUN letters", func() {
		// 	machine := enigma.NewEnigma(
		// 		"B",
		// 		[]string{"I", "II", "III"},
		// 		[]rune{},
		// 	)
		// 	machine.ConfigureRotors("FUM")

		// 	rotorI := machine.GetRotor("I")
		// 	Expect(rotorI.GetTopLetter()).To(Equal('F'))

		// 	rotorII := machine.GetRotor("II")
		// 	Expect(rotorII.GetTopLetter()).To(Equal('U'))

		// 	rotorIII := machine.GetRotor("III")
		// 	Expect(rotorIII.GetTopLetter()).To(Equal('M'))

		// 	machine.Step()
		// 	outLetter := machine.EncipherLetter('A')
		// 	Expect(outLetter).To(Equal('Y'), "FUN(Y): Got [%c]", outLetter)

		// })

		It("Can encipher a string - AAA", func() {
			var input string
			var expOutput string

			machine := enigma.NewEnigma(
				"B",
				[]string{"I", "II", "III"},
				[]rune{},
			)

			machine.ConfigureRotors("AAA")
			input = "AAAAA"
			expOutput = "BDZGO"
			Expect(machine.EncipherString(input)).To(Equal(expOutput), input)

			machine.ConfigureRotors("AAA")
			input = "CRAIG"
			expOutput = "QCZQF"
			Expect(machine.EncipherString(input)).To(Equal(expOutput), input)
			// continue with existing rotor positions
			input = "CATE"
			expOutput = "MCRW"
			Expect(machine.EncipherString(input)).To(Equal(expOutput), input)

			machine.ConfigureRotors("AAA")
			input = "CATE"
			expOutput = "QDHW"
			Expect(machine.EncipherString(input)).To(Equal(expOutput), input)

			machine.ConfigureRotors("AAA")
			input = "HELLO WORLD"
			expOutput = "ILBDA AMTAZ"
			Expect(machine.EncipherString(input)).To(Equal(expOutput), input)

			machine.ConfigureRotors("AAA")
			input = "craig, cate"
			expOutput = "QCZQF, MCRW"
			Expect(machine.EncipherString(input)).To(Equal(expOutput), input)

			machine.ConfigureRotors("AAA")
			input = "heLLo 42 World"
			expOutput = "ILBDA 42 AMTAZ"
			Expect(machine.EncipherString(input)).To(Equal(expOutput), input)

		})

		It("Can encipher a string - AAB", func() {
			var input string
			var expOutput string

			machine := enigma.NewEnigma(
				"B",
				[]string{"I", "II", "III"},
				[]rune{},
			)

			machine.ConfigureRotors("AAB")
			input = "PYTHON"
			expOutput = "HWFMKR"
			Expect(machine.EncipherString(input)).To(Equal(expOutput), input)

			machine.ConfigureRotors("AAB")
			input = "this is the way the world ends"
			expOutput = "ZTQB LV RTC CSZ ABN HRSFF PZMN"
			Expect(machine.EncipherString(input)).To(Equal(expOutput), input)
		})

		It("Can encipher a string - Various", func() {
			var input string

			machine := enigma.NewEnigma(
				"B",
				[]string{"I", "II", "III"},
				[]rune{},
			)

			fmt.Println()
			input = "PYTHON"
			tests := map[string]string{
				"AAA": "LMHKAE",
				"AAB": "HWFMKR",
				"ABB": "CQJQTH",
				"AUN": "WJVRYC",
				"BBB": "AWGLLT",
				"BAA": "IALIMQ",
				"FUN": "EGDBTJ",
			}
			for cfg, expOut := range tests {
				machine.ConfigureRotors(cfg)
				Expect(machine.EncipherString(input)).To(Equal(expOut), "%s) %s -> %s", cfg, input, expOut)

			}
		})

		It("Ocaml FUN", func() {
			machine := enigma.NewEnigma(
				"B",
				[]string{"I", "II", "III"},
				[]rune{'A', 'Z'},
			)
			// machine.ToggleTrace()
			machine.ConfigureRotors("FUN")
			output := machine.EncipherString("YNGXQ")
			Expect(output).To(Equal("OCAML"))
		})
	})
})
