package main

import (
	"fmt"
	"time"
)

func main() {

	for {
		select {
		case <-time.Tick(time.Second):
			fmt.Println("tick")
		}
	}

}
