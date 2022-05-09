package pkge

import (
	"github.com/nunnatsa/issue/cmd/cmdb"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestE(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "E Suite")
}

var _ = Describe("test package E", func() {
	It("should work", func() {
		Expect(E()).To(Succeed())
		Expect(cmdb.B()).To(Succeed())
	})
})
