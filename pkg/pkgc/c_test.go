package pkgc

import (
	"github.com/nunnatsa/issue/cmd/cmdb"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/nunnatsa/issue/pkg/pkge"
)

func TestC(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "C Suite")
}

var _ = Describe("test package C", func() {
	It("should work", func() {
		Expect(C()).To(Succeed())
	})

	It("E should work", func() {
		Expect(pkge.E()).To(Succeed())
		Expect(cmdb.B())
	})
})
