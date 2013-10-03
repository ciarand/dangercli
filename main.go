package main

import (
	"github.com/ciarand/dangerlib"
	"os"
)

func main() {
	if string(os.Args[1]) == "server" {
		dangerlib.Serve()
		return
	}
	dangerlib.Connect(os.Args[1:])
	return
}
