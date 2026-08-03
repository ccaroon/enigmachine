package enigma_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ccaroon/enigmachine/enigma"
)

var _ = Describe("Cable", func() {
	It("Knows what letters it's connected to", func() {
		cable := enigma.Cable{
			Plug1: 'D',
			Plug2: 'X',
		}

		Expect(cable.ConnectedTo('D')).To(BeTrue())
		Expect(cable.ConnectedTo('X')).To(BeTrue())

		Expect(cable.ConnectedTo('A')).To(BeFalse())
	})

	It("Can follow plug1 to plug2", func() {
		cable := enigma.NewCable('R', 'X')

		Expect(cable.Follow('R')).To(Equal('X'))
		Expect(cable.Follow('X')).To(Equal('R'))
	})
})
