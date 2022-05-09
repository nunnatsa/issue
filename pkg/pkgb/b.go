package pkgb

import (
	"fmt"
	"time"
)

func B() error {
	fmt.Println("I'm B")
	time.Sleep(time.Second)
	return nil
}
