package pkgd

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/nunnatsa/issue/pkg/pkge"
)

func TestD(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "D Suite")
}

var _ = Describe("test package D", func() {
	It("should work", func() {
		Expect(D()).To(Succeed())
		Expect(pkge.E()).To(Succeed())
	})
})
