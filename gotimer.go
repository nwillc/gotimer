package main

import (
	"flag"
	"fmt"
	"github.com/nwillc/gotimer/setup"
)

func main() {
	flag.Parse()

	fmt.Println("Counting down:", *setup.Flags.Time)
}
