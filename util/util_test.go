package util_test

import (
	"os"

	"github.com/ccaroon/enigmachine/util"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Util", func() {
	It("Can get the data home/directory", func() {
		// XDG_DATA_HOME NOT set
		err := os.Unsetenv("XDG_DATA_HOME")
		Expect(err).To(BeNil())
		Expect(os.Getenv("XDG_DATA_HOME")).To(BeEmpty())

		dataDir := util.GetDataDir()
		Expect(dataDir).To(Not(BeEmpty()))
		Expect(dataDir).To(Equal(os.ExpandEnv("$HOME/.local/share")))

		// XDG_DATA_HOME IS set
		err = os.Setenv("XDG_DATA_HOME", "/home/zaphod/data")
		Expect(err).To(BeNil())

		dataDir = util.GetDataDir()
		Expect(dataDir).To(Equal("/home/zaphod/data"))
	})
})
