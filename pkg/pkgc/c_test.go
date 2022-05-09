package pkgc

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/nunnatsa/issue/cmd/cmdb"
	"github.com/nunnatsa/issue/pkg/pkgd"
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
		Expect(pkgd.D()).To(Succeed())
		Expect(cmdb.B())
	})
})
