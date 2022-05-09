package cmda

import (
	"fmt"
	"time"

	"github.com/nunnatsa/issue/pkg/pkge"
)

func A() error {
	fmt.Println("I'm CMD A")
	time.Sleep(time.Second)
	return pkge.E()
}