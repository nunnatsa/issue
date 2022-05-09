package pkgd

import (
	"github.com/nunnatsa/issue/cmd/cmdb"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestD(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "D Suite")
}

var _ = Describe("test package D", func() {
	It("should work", func() {
		Expect(D()).To(Succeed())
		Expect(cmdb.B()).To(Succeed())
	})
})
