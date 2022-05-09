package cmdb

import (
	"fmt"
	"time"
)

func B() error {
	fmt.Println("I'm CMD B")
	time.Sleep(time.Second)
	return nil
}
