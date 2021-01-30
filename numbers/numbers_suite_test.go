package numbers

import (
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCodewarsGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Numbers Package Suite")
}