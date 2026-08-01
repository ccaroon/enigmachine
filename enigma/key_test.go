package enigma_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ccaroon/enigmachine/enigma"
)

var _ = Describe("Key", func() {
	It("Can Parse a Key Specification String", func() {
		key, err := enigma.ParseKeySpec("B:I,II,III:AZ,XY,EM")

		Expect(err).To(BeNil())

		Expect(key.ReflId).To(Equal("B"))
		Expect(key.RotorIds).To(Equal([]string{"I", "II", "III"}))
		Expect(key.RotorCfg).To(Equal("AAA"))
		Expect(key.PbSpec).To(Equal([]string{"AZ", "XY", "EM"}))
	})

	It("Fails on invalid Key Spec: Empty", func() {
		key, err := enigma.ParseKeySpec("")

		Expect(key).To(BeNil())
		Expect(err).To(Not(BeNil()))
		Expect(err).To(MatchError("Invalid Key Specification: []"))
	})

	It("Fails on invalid Key Spec: No Rotors", func() {
		key, err := enigma.ParseKeySpec("B")

		Expect(key).To(BeNil())
		Expect(err).To(Not(BeNil()))
		Expect(err).To(MatchError("Invalid Key Specification. Missing Rotor Specs: [B] "))
	})

	It("Fails on invalid Key Spec: Wrong Number of Rotors", func() {
		key, err := enigma.ParseKeySpec("B:I,II")

		Expect(key).To(BeNil())
		Expect(err).To(Not(BeNil()))
		Expect(err).To(MatchError("Invalid Rotor Specs: [I,II]"))
	})

	It("Fails on invalid Key Spec: Bad Rotor Format", func() {
		key, err := enigma.ParseKeySpec("B:I@B,II@A,@D")

		Expect(key).To(BeNil())
		Expect(err).To(Not(BeNil()))
		Expect(err).To(MatchError("Invalid Rotor: [@D]"))
	})
})
