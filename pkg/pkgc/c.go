package pkgc

import (
	"fmt"
	"time"

	"github.com/nunnatsa/issue/pkg/pkge"
)

func C() error {
	fmt.Println("I'm C")
	time.Sleep(time.Second)
	return pkge.E()
}
