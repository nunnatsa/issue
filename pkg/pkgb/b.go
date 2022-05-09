package pkgb

import (
	"fmt"
	"time"

	"github.com/nunnatsa/issue/pkg/pkgd"
)

func B() error {
	fmt.Println("I'm B")
	time.Sleep(time.Second)
	return pkgd.D()
}
