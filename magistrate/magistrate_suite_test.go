package magistrate_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMagistrate(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "Magistrate Suite")
}
