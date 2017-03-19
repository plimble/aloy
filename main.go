package main

import (
	"fmt"
	"time"

	"github.com/plimble/aloy/api"
)

func main() {
	fmt.Println("Start app")
	time.Sleep(4 * time.Second)
	fmt.Println("Gogo")
	api.Run()
}
