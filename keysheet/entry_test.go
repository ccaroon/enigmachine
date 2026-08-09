package keysheet_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ccaroon/enigmachine/keysheet"
)

var _ = Describe("Entry", func() {

	Context("Generates", func() {
		It("Valid Rotor Selection", func() {
			entry := keysheet.GenerateEntry()

			Expect(entry).To(Not(BeNil()))
			Expect(len(entry.Rotors)).To(Equal(3))

			Expect(entry.Rotors[0]).To(Not(BeEmpty()))
			Expect(entry.Rotors[1]).To(Not(BeEmpty()))
			Expect(entry.Rotors[2]).To(Not(BeEmpty()))

			// Each Rotor must be different
			Expect(entry.Rotors[0]).To(Not(Equal(entry.Rotors[1])))
			Expect(entry.Rotors[0]).To(Not(Equal(entry.Rotors[2])))
			Expect(entry.Rotors[1]).To(Not(Equal(entry.Rotors[2])))
		})

		It("Valid Ring Settings", func() {
			entry := keysheet.GenerateEntry()

			Expect(entry).To(Not(BeNil()))
			Expect(len(entry.RingSettings)).To(Equal(3))

			Expect(entry.RingSettings[0]).To(Not(Equal(0)))
			Expect(entry.RingSettings[1]).To(Not(Equal(0)))
			Expect(entry.RingSettings[2]).To(Not(Equal(0)))

			for idx := range 3 {
				value := entry.RingSettings[idx]
				Expect(value >= 1 && value <= 26).To(BeTrue())

			}

			// Each ring setting must be different
			Expect(entry.RingSettings[0]).To(Not(Equal(entry.RingSettings[1])))
			Expect(entry.RingSettings[0]).To(Not(Equal(entry.RingSettings[2])))
			Expect(entry.RingSettings[1]).To(Not(Equal(entry.RingSettings[2])))
		})

		It("A Valid Reflector", func() {
			entry := keysheet.GenerateEntry()
			Expect(entry).To(Not(BeNil()))

			Expect(entry.Reflector).To(Equal('B'))
		})

		It("Valid Plugboard Swaps", func() {
			entry := keysheet.GenerateEntry()

			Expect(entry).To(Not(BeNil()))
			Expect(len(entry.PlugboardSwaps)).To(Equal(10))

			for idx := range 10 {
				Expect(entry.PlugboardSwaps[idx]).To(HaveLen(2))
			}
		})

		It("Valid Day Keys", func() {
			entry := keysheet.GenerateEntry()

			Expect(entry).To(Not(BeNil()))
			Expect(len(entry.DayKeys)).To(Equal(4))

			for idx := range 4 {
				Expect(entry.DayKeys[idx]).To(HaveLen(3))
			}
		})
	})

	Context("Output", func() {
		It("Can format an entry for printing", func() {})
		entry := keysheet.GenerateEntry()
		entry.Day = 27

		Expect(entry).To(Not(BeNil()))

		output := entry.Format()
		Expect(output).To(Not(BeEmpty()))

		Expect(output).To(ContainSubstring("27."))
	})
})
