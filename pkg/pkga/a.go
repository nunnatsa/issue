package pkga

import (
	"fmt"
	"time"

	"github.com/nunnatsa/issue/pkg/pkgd"
)

func A() error {
	fmt.Println("I'm A")
	time.Sleep(time.Second)
	return pkgd.D()
}
