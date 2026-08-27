package actions_test

import (
	"os"
	"path"
	"slices"
	"strings"

	"github.com/bykof/gostradamus"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ccaroon/enigmachine/actions"
	"github.com/ccaroon/enigmachine/keysheet"
)

func genCurrentKeySheet(network string) *keysheet.KeySheet {
	now := gostradamus.Now()

	err := os.Setenv("XDG_DATA_HOME", "testdata/tmp")
	Expect(err).To(BeNil())

	ks := keysheet.GenerateKeySheet(network, now.Month(), now.Year())
	ksFile := keysheet.BuildPath(network, now.Month(), now.Year())

	err = os.MkdirAll(path.Dir(ksFile), 0755)
	Expect(err).To(BeNil())

	err = ks.Save(ksFile)
	Expect(err).To(BeNil())

	return ks
}

var _ = Describe("EncodeDecode", func() {

	activeKeySheet := genCurrentKeySheet("army")
	Expect(activeKeySheet).ToNot(BeNil())

	Context("Encode", func() {
		It("Can Encode using Default Settings", func() {
			args := actions.EncodeDecodeArgs{
				Input:       "CRAIG",
				Action:      "encode",
				BlockSize:   5,
				LineSize:    10,
				KeepOrigFmt: false,
			}

			output, err := actions.EncodeDecode(args)
			Expect(err).To(BeNil())
			output = strings.TrimSpace(output)
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

			output, err := actions.EncodeDecode(args)
			Expect(err).To(BeNil())
			output = strings.TrimSpace(output)
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
				Input:       "@testdata/clear-text.txt",
				Action:      "encode",
				BlockSize:   blockSize,
				LineSize:    lineSize,
				KeepOrigFmt: false,
			}

			output, err := actions.EncodeDecode(args)
			Expect(err).To(BeNil())
			output = strings.TrimSpace(output)
			Expect(output).To(Not(BeEmpty()))

			lines := strings.Split(output, "\n")
			Expect(len(lines)).To(Equal(5))

			aLine := strings.TrimSpace(lines[0])
			blocks := strings.Split(aLine, " ")
			Expect(len(blocks)).To(Equal(lineSize))
			Expect(len(blocks[0])).To(Equal(blockSize))
		})

		It("Can Encode a File keeping Original Format", func() {
			args := actions.EncodeDecodeArgs{
				Input:       "@testdata/clear-text.txt",
				Action:      "encode",
				KeepOrigFmt: true,
			}

			output, err := actions.EncodeDecode(args)
			Expect(err).To(BeNil())
			output = strings.TrimSpace(output)
			Expect(output).To(Not(BeEmpty()))

			lines := strings.Split(output, "\n")
			Expect(len(lines)).To(Equal(5))

			// Lines should not be the same length
			Expect(len(lines[0])).ToNot(Equal(len(lines[1])))

			// Lines should end with '.'
			aLine := strings.TrimSpace(lines[0])
			Expect(strings.HasSuffix(aLine, ".")).To(BeTrue())

			// blocks (words) on a line should NOT be the same length
			blocks := strings.Split(aLine, " ")
			blockLens := make([]int, len(blocks))
			for idx, block := range blocks {
				blockLens[idx] = len(block)
			}

			min := slices.Min(blockLens)
			max := slices.Max(blockLens)
			Expect(min).ToNot(Equal(max))
		})

		It("Can Encode using a Key Sheet", func() {
			ksEntry := activeKeySheet.ActiveEntry()
			args := actions.EncodeDecodeArgs{
				Input:     "@testdata/clear-text.txt",
				Action:    "encode",
				Network:   activeKeySheet.Network,
				BlockSize: 5,
				LineSize:  15,
			}

			output, err := actions.EncodeDecode(args)
			Expect(err).To(BeNil())
			output = strings.TrimSpace(output)
			Expect(output).To(Not(BeEmpty()))

			lines := strings.Split(output, "\n")
			blocks := strings.Split(lines[0], " ")

			// Check that first block is the day Key
			block1 := blocks[0]
			Expect(len(block1)).To(Equal(5))
			dayKey := block1[2:]

			Expect(slices.Contains(ksEntry.DayKeys, dayKey)).To(BeTrue())

		})
	})

	Context("Decode", func() {

	})
})
