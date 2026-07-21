package enigma_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/ccaroon/enigmachine/enigma"
)

var _ = Describe("Plugboard", func() {
	var pb *enigma.Plugboard

	BeforeEach(func() {
		cablePairs := []rune{
			'K', 'Z',
			'X', 'G',
			'O', 'I',
			'A', 'B',
		}
		pb = enigma.NewPlugboard(cablePairs)
	})

	Context("Creation", func() {
		It("Should be able to create an 'empty' plugboard", func() {
			pb2 := enigma.NewPlugboard([]rune{})

			Expect(pb2.NumCables()).To(Equal(0))
			Expect(pb2.GetCable(0)).To(BeNil())
			Expect(pb2.Map('A')).To(Equal('A'))
			Expect(pb2.Map('B')).To(Equal('B'))
			Expect(pb2.Map('C')).To(Equal('C'))
			Expect(pb2.Map('X')).To(Equal('X'))
			Expect(pb2.Map('Y')).To(Equal('Y'))
			Expect(pb2.Map('Z')).To(Equal('Z'))

		})

		It("Should create a valid plugboard given valid inputs", func() {
			Expect(pb.NumCables()).To(Equal(4))

			cable1 := pb.GetCable(0)
			Expect(cable1.Plug1).To(Equal('K'))
			Expect(cable1.Plug2).To(Equal('Z'))
		})

		It("Should create a valid plugboard given uneven cable pairings", func() {
			cablePairs := []rune{
				'A', 'Z',
				'F', 'X',
				// dangling/unconnected
				'M',
			}
			pb2 := enigma.NewPlugboard(cablePairs)

			Expect(pb2.NumCables()).To(Equal(2))

			cable1 := pb2.GetCable(0)
			Expect(cable1.Plug1).To(Equal('A'))
			Expect(cable1.Plug2).To(Equal('Z'))

			cable2 := pb2.GetCable(1)
			Expect(cable2.Plug1).To(Equal('F'))
			Expect(cable2.Plug2).To(Equal('X'))

			cable3 := pb2.GetCable(2)
			Expect(cable3).To(BeNil())

			Expect(pb.FindCable('M')).To(BeNil())
		})
	})

	Context("Usage", func() {
		It("Can find a cable", func() {
			cable := pb.FindCable('X')
			Expect(cable).To(Not(BeNil()))
			Expect(cable.Follow('X')).To(Equal('G'))
			Expect(cable.Follow('G')).To(Equal('X'))

			// not connected
			cable2 := pb.FindCable('J')
			Expect(cable2).To(BeNil())
		})

		It("Can map one letter to another", func() {
			Expect(pb.Map('K')).To(Equal('Z'))
			Expect(pb.Map('O')).To(Equal('I'))
		})

		It("Should map an unconnected letter to itself", func() {
			// unconnected
			Expect(pb.Map('Q')).To(Equal('Q'))
			Expect(pb.Map('J')).To(Equal('J'))
		})
	})

})
