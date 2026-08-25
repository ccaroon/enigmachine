package actions_test

import (
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ccaroon/enigmachine/actions"
)

var _ = Describe("Encode", func() {

	It("Can Encode using Default Settings", func() {
		args := actions.EncodeDecodeArgs{
			Input:       "CRAIG",
			Action:      "encode",
			BlockSize:   5,
			LineSize:    10,
			KeepOrigFmt: false,
		}

		output := strings.TrimSpace(actions.EncodeDecode(args))
		Expect(output).To(Not(BeEmpty()))

		Expect(len(output)).To(Equal(5))
		Expect(output).To(Equal("QCZQF"))
	})

	It("Can Encode using User Settings", func() {
		args := actions.EncodeDecodeArgs{
			Input:       "We are the hollow men",
			Action:      "encode",
			KeySpec:     "B:II,III,I:AX,FG,PQ",
			RotorCfg:    "MEN",
			BlockSize:   5,
			LineSize:    10,
			KeepOrigFmt: false,
		}

		output := strings.TrimSpace(actions.EncodeDecode(args))
		Expect(output).To(Not(BeEmpty()))

		Expect(len(output)).To(Equal(20))

		blocks := strings.Split(output, " ")
		Expect(len(blocks)).To(Equal(4))
		Expect(len(blocks[0])).To(Equal(5))

		Expect(output).To(Equal("VCZJG EFCMC TFHKU OF"))
	})

	It("Can Encode a File", func() {
		var blockSize int = 5
		var lineSize int = 10
		args := actions.EncodeDecodeArgs{
			Input:       "@testdata/encode-test.txt",
			Action:      "encode",
			BlockSize:   blockSize,
			LineSize:    lineSize,
			KeepOrigFmt: false,
		}

		output := strings.TrimSpace(actions.EncodeDecode(args))
		Expect(output).To(Not(BeEmpty()))

		lines := strings.Split(output, "\n")
		Expect(len(lines)).To(Equal(5))

		aLine := strings.TrimSpace(lines[0])
		blocks := strings.Split(aLine, " ")
		Expect(len(blocks)).To(Equal(lineSize))
		Expect(len(blocks[0])).To(Equal(blockSize))
	})

	// It("Can Encode a File keeping Original Format", func() {
	// 	var blockSize int = 5
	// 	var lineSize int = 10
	// 	args := actions.EncodeDecodeArgs{
	// 		Input:       "@testdata/encode-test.txt",
	// 		Action:      "encode",
	// 		BlockSize:   blockSize,
	// 		LineSize:    lineSize,
	// 		KeepOrigFmt: true,
	// 	}

	// 	output := strings.TrimSpace(actions.EncodeDecode(args))
	// 	Expect(output).To(Not(BeEmpty()))

	// 	lines := strings.Split(output, "\n")
	// 	Expect(len(lines)).To(Equal(5))

	// 	aLine := strings.TrimSpace(lines[0])
	// 	blocks := strings.Split(aLine, " ")
	// 	Expect(len(blocks)).To(Equal(lineSize))
	// 	Expect(len(blocks[0])).To(Equal(blockSize))
	// })

})
