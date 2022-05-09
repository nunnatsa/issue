package pkge

import (
	"fmt"
	"time"
)

func E() error {
	fmt.Println("I'm E")
	time.Sleep(time.Second)
	return nil
}
