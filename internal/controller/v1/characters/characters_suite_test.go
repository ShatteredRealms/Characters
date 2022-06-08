package characters_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCharacters(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Characters Suite")
}
