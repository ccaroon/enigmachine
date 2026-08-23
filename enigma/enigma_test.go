package enigma_test

import (
	"strings"

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
	var machine *enigma.Enigma

	BeforeEach(func() {
		var err error

		machine, err = enigma.NewEnigma(
			'B',
			[]string{"I", "II", "III"},
			[]string{},
		)

		Expect(err).To(BeNil())
		Expect(machine).ToNot(BeNil())
	})

	It("Can get a Rotor by Id", func() {
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
		Expect(machine.GetRotorByIdx(0).Id()).To(Equal("I"))
		Expect(machine.GetRotorByIdx(1).Id()).To(Equal("II"))
		Expect(machine.GetRotorByIdx(2).Id()).To(Equal("III"))
	})

	It("Can set starting config for Rotors", func() {
		machine, err := enigma.NewEnigma(
			'B',
			[]string{"III", "II", "I"},
			[]string{},
		)
		Expect(err).To(BeNil())
		Expect(machine).ToNot(BeNil())

		machine.ConfigureRotors("ABC")
		testRotorConfig(machine, "ABC")

		machine.ConfigureRotors("ZYX")
		testRotorConfig(machine, "ZYX")
	})

	It("Can properly step multiple rotors", func() {
		machine, err := enigma.NewEnigma(
			'B',
			[]string{"III", "II", "I"},
			[]string{},
		)
		Expect(err).To(BeNil())
		Expect(machine).ToNot(BeNil())

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

		It("Can encipher a string - AAA", func() {
			var input string
			var expOutput string

			machine.ConfigureRotors("AAA")
			input = "AAAAA"
			expOutput = "BDZGO"
			output := machine.EncipherString(input, true)
			Expect(output.String()).To(Equal(expOutput), input)

			machine.ConfigureRotors("AAA")
			input = "CRAIG"
			expOutput = "QCZQF"
			output = machine.EncipherString(input, true)
			Expect(output.String()).To(Equal(expOutput), input)
			// continue with existing rotor positions
			input = "CATE"
			expOutput = "MCRW"
			output = machine.EncipherString(input, true)
			Expect(output.String()).To(Equal(expOutput), input)

			machine.ConfigureRotors("AAA")
			input = "CATE"
			expOutput = "QDHW"
			output = machine.EncipherString(input, true)
			Expect(output.String()).To(Equal(expOutput), input)

			machine.ConfigureRotors("AAA")
			input = "HELLO WORLD"
			expOutput = "ILBDAAMTAZ"
			output = machine.EncipherString(input, true)
			Expect(output.String()).To(Equal(expOutput), input)

			machine.ConfigureRotors("AAA")
			input = "craig, cate"
			expOutput = "QCZQFMCRW"
			output = machine.EncipherString(input, true)
			Expect(output.String()).To(Equal(expOutput), input)

			machine.ConfigureRotors("AAA")
			input = "heLLo 42 World"
			expOutput = "ILBDAAMTAZ"
			output = machine.EncipherString(input, true)
			Expect(output.String()).To(Equal(expOutput), input)

		})

		It("Can encipher a string - AAB", func() {
			var input string
			var expOutput string

			machine.ConfigureRotors("AAB")
			input = "PYTHON"
			expOutput = "HWFMKR"
			output := machine.EncipherString(input, true)
			Expect(output.String()).To(Equal(expOutput), input)

			machine.ConfigureRotors("AAB")
			input = "this is the way the world ends"
			expOutput = "ZTQBLVRTCCSZABNHRSFFPZMN"
			output = machine.EncipherString(input, true)
			Expect(output.String()).To(Equal(expOutput), input)
		})

		It("Can encipher a string - Various", func() {
			var input string

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
				output := machine.EncipherString(input, true)
				Expect(output.String()).To(Equal(expOut), "%s) %s -> %s", cfg, input, expOut)

			}
		})

		It("Ocaml FUN", func() {
			machine, err := enigma.NewEnigma(
				'B',
				[]string{"I", "II", "III"},
				[]string{"AZ"},
			)
			Expect(err).To(BeNil())
			Expect(machine).ToNot(BeNil())

			// machine.ToggleTrace()
			machine.ConfigureRotors("FUN")
			output := machine.EncipherString("YNGXQ", true)
			Expect(output.String()).To(Equal("OCAML"))
		})
	})

	Context("Output Formatting", func() {

		It("Can keep original input formatting", func() {
			options := enigma.NewFormatOptions(0, 0, true)
			machine.ConfigureRotors("FOO")

			input := `This is the way the world ends
This is the way the world ends
This is the way the world ends
Not with a bang but a whimper.`

			// HVCE PX EMI PGK WDU STFSI LIUU
			// ZUYG MN KLY REP PRW QXSNV YJZO
			// EMJJ JT ELA UHV ACQ DYNWI XHMG
			// GWC KRNV B GFFF NTW Q BZZBWOE.
			encodedData := machine.EncipherString(input, false)
			output := encodedData.Format(options)

			lines := strings.Split(output, "\n")
			Expect(len(lines)).To(Equal(4))

			// Examine last line
			blocks := strings.Split(lines[3], " ")
			Expect(len(blocks)).To(Equal(7))

			aWord := blocks[0]
			Expect(len(aWord)).To(Equal(3))

			aWord = blocks[2]
			Expect(len(aWord)).To(Equal(1))

			aWord = blocks[6]
			Expect(len(aWord)).To(Equal(8))
			Expect(aWord[7]).To(Equal(byte('.')))

		})

		It("Can format output by blocks & lines", func() {
			options := enigma.NewFormatOptions(5, 7, false)
			machine.ConfigureRotors("NIL")

			input := `This is the way the world ends
This is the way the world ends
This is the way the world ends
Not with a bang but a whimper.
`

			// OPGND XHVWJ KJNXS TZOMG YVFPK UFQMI ANRLF
			// FZZNS JVCUV AKFOM YLFHC RCDHF JKPGF GSZTY
			// MFCAA SMMJB AZWEC CDYXG AXMZH
			encodedData := machine.EncipherString(input, true)
			output := encodedData.Format(options)

			lines := strings.Split(output, "\n")
			Expect(len(lines)).To(Equal(3))

			words := strings.Split(strings.Trim(lines[0], " "), " ")
			Expect(len(words)).To(Equal(7))

			words = strings.Split(strings.Trim(lines[1], " "), " ")
			Expect(len(words)).To(Equal(7))

			words = strings.Split(strings.Trim(lines[2], " "), " ")
			Expect(len(words)).To(Equal(5))
		})
	})

	Context("Error Handling", func() {
		It("Cannot use unknown Reflector", func() {
			machine, err := enigma.NewEnigma(
				'X',
				[]string{"I", "II", "III"},
				[]string{},
			)

			Expect(machine).To(BeNil())
			Expect(err).To(MatchError("Unsupported Reflector: 'X'"))
		})

		It("Cannot contain duplicate rotors", func() {
			machine, err := enigma.NewEnigma(
				'B',
				[]string{"I", "I", "III"},
				[]string{},
			)

			Expect(machine).To(BeNil())
			Expect(err).To(MatchError("Duplicate Rotors Detected: [I @ 1]"))
		})

		It("Cannot contain duplicate plugboard settings", func() {
			machine, err := enigma.NewEnigma(
				'B',
				[]string{"I", "II", "III"},
				[]string{"XC", "QW", "ER", "CX"},
			)

			Expect(machine).To(BeNil())
			Expect(err).To(MatchError("Duplicate Plugboard Setting: [CX] [XC]"))
		})
	})
})
