package pkgd

import (
	"fmt"
	"time"
)

func D() error {
	fmt.Println("I'm D")
	time.Sleep(time.Second)
	return nil
}
