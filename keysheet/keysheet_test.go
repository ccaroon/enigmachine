package keysheet_test

import (
	"fmt"
	"math/rand/v2"
	"os"
	"path"
	"strings"

	"github.com/bykof/gostradamus"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ccaroon/enigmachine/keysheet"
	"github.com/ccaroon/enigmachine/util"
)

var _ = Describe("Keysheet", func() {
	Context("Generates", func() {
		It("Valid Key Sheet", func() {
			ks := keysheet.GenerateKeySheet("caroon.org", 8, 2026)

			Expect(ks).To(Not(BeNil()))

			Expect(ks.Network).To(Equal("caroon.org"))
			Expect(ks.Month).To(Equal(8))
			Expect(ks.Year).To(Equal(2026))
			Expect(len(ks.Entries)).To(Equal(31))
		})
	})

	Context("Persistence", func() {
		It("Can build data path", func() {
			var network string = "navy"
			var month int = 2
			var year int = 1792

			dataDir := util.GetDataDir()

			path := keysheet.BuildPath(network, month, year)

			Expect(path).To(Not(BeEmpty()))
			Expect(path).To(ContainSubstring(dataDir))
			Expect(path).To(ContainSubstring(network))
			Expect(path).To(ContainSubstring(
				fmt.Sprintf("%4d-%02d.yml", year, month),
			))
		})

		It("Can save a Key Sheet", func() {
			var network = "airforce"
			var month int = rand.IntN(11) + 1
			var year int = rand.IntN(99) + 1900
			var savePath string = fmt.Sprintf("/tmp/%s-%4d-%02d.yml", network, year, month)

			ks := keysheet.GenerateKeySheet(network, month, year)

			// Check for existence of file - should NOT exist
			_, err := os.Open(savePath)
			Expect(err).To(Not(BeNil()))
			Expect(os.IsNotExist(err)).To(BeTrue())

			err = ks.Save(savePath)
			Expect(err).To(BeNil())

			// Check that file now exists
			file, err := os.Open(savePath)
			Expect(err).To(BeNil())

			// clean up
			err = file.Close()
			Expect(err).To(BeNil())

			err = os.Remove(savePath)
			Expect(err).To(BeNil())
		})

		It("Can load an existing Key Sheet", func() {
			var network = "airforce"
			var month int = rand.IntN(11) + 1
			var year int = rand.IntN(99) + 1900
			var savePath string = fmt.Sprintf("/tmp/%s-%4d-%02d.yml", network, year, month)

			ks1 := keysheet.GenerateKeySheet(network, month, year)
			err := ks1.Save(savePath)
			Expect(err).To(BeNil())

			ks2, err := keysheet.LoadKeySheet(savePath)
			Expect(err).To(BeNil())

			Expect(ks2.Network).To(Equal(ks1.Network))
			Expect(ks2.Month).To(Equal(ks1.Month))
			Expect(ks2.Year).To(Equal(ks1.Year))
			Expect(len(ks2.Entries)).To(Equal(len(ks1.Entries)))
		})

		It("Can load active/current Key Sheet", func() {
			os.Setenv("XDG_DATA_HOME", "/tmp")
			dt := gostradamus.Now()

			network := "marines"
			month := dt.Month()
			year := dt.Year()

			ks1 := keysheet.GenerateKeySheet(network, month, year)

			dataDir := util.GetDataDir()
			Expect(dataDir).To(Equal("/tmp"))

			savePath := keysheet.BuildPath(network, month, year)
			Expect(strings.HasPrefix(savePath, "/tmp")).To(BeTrue())

			err := os.MkdirAll(path.Dir(savePath), 0755)
			err = ks1.Save(savePath)
			Expect(err).To(BeNil())

			ks2, err := keysheet.LoadActiveKeySheet(ks1.Network)
			Expect(err).To(BeNil())

			Expect(ks2.Network).To(Equal(ks1.Network))
			Expect(ks2.Month).To(Equal(ks1.Month))
			Expect(ks2.Year).To(Equal(ks1.Year))
		})
	})

	Context("Humans", func() {
		It("Can format Key Sheet for Human Consumption", func() {
			var network string = "green peace"
			var month int = 3
			var year int = 1772

			ks := keysheet.GenerateKeySheet(network, month, year)

			viewStr := ks.String()
			Expect(viewStr).To(Not(BeEmpty()))

			Expect(viewStr).To(ContainSubstring("GREEN PEACE"))
			Expect(viewStr).To(ContainSubstring("March 1772"))
		})
	})
})
