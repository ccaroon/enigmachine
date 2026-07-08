package enigma_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ccaroon/enigmachine/enigma"
)

var _ = Describe("Plugboard", func() {
	var pb *enigma.Plugboard

	BeforeEach(func() {
		cablePairs := []byte{
			byte('K'), byte('Z'),
			byte('X'), byte('G'),
			byte('O'), byte('I'),
			byte('A'), byte('B'),
		}
		pb = enigma.NewPlugboard(cablePairs)
	})

	Context("Creation", func() {
		It("Should create a valid plugboard given valid inputs", func() {
			Expect(pb.NumCables()).To(Equal(4))

			cable1 := pb.GetCable(0)
			Expect(cable1.Plug1).To(Equal(byte('K')))
			Expect(cable1.Plug2).To(Equal(byte('Z')))
		})

		It("Should create a valid plugboard given uneven cable pairings", func() {
			cablePairs := []byte{
				byte('A'), byte('Z'),
				byte('F'), byte('X'),
				// dangling/unconnected
				byte('M'),
			}
			pb2 := enigma.NewPlugboard(cablePairs)

			Expect(pb2.NumCables()).To(Equal(2))

			cable1 := pb2.GetCable(0)
			Expect(cable1.Plug1).To(Equal(byte('A')))
			Expect(cable1.Plug2).To(Equal(byte('Z')))

			cable2 := pb2.GetCable(1)
			Expect(cable2.Plug1).To(Equal(byte('F')))
			Expect(cable2.Plug2).To(Equal(byte('X')))

			cable3 := pb2.GetCable(2)
			Expect(cable3).To(BeNil())

			Expect(pb.FindCable('M')).To(BeNil())
		})
	})

	Context("Usage", func() {
		It("Can find a cable", func() {
			cable := pb.FindCable('X')
			Expect(cable).To(Not(BeNil()))
			Expect(cable.Follow('X')).To(Equal(byte('G')))
			Expect(cable.Follow('G')).To(Equal(byte('X')))

			// not connected
			cable2 := pb.FindCable('J')
			Expect(cable2).To(BeNil())
		})

		It("Can map one letter to another", func() {
			Expect(pb.Map('K')).To(Equal(byte('Z')))
			Expect(pb.Map('O')).To(Equal(byte('I')))
		})

		It("Should map an unconnected letter to itself", func() {
			// unconnected
			Expect(pb.Map('Q')).To(Equal(byte('Q')))
			Expect(pb.Map('J')).To(Equal(byte('J')))
		})

	})

})
