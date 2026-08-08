package keysheet_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestKeysheet(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Keysheet Suite")
}
