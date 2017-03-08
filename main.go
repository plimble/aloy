package main

import "github.com/plimble/aloy/api"

func main() {
	srv := api.New()

	srv.Listen(":4400")
}
