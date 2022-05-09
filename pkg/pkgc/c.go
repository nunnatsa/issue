package pkgc

import (
	"fmt"
	"time"

	"github.com/nunnatsa/issue/pkg/pkgd"
)

func C() error {
	fmt.Println("I'm C")
	time.Sleep(time.Second)
	return pkgd.D()
}
