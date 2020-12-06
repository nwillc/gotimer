package main

import (
	"flag"
	"fmt"
	"github.com/nwillc/gotimer/setup"
	"time"
)

func main() {
	flag.Parse()

	duration, err := time.ParseDuration(*setup.Flags.Time)
	if err != nil {
		panic(err)
	}
	fmt.Println("Duration:", duration)
	finished := make(chan bool)
	completed := time.Now().Add(duration)
	go func(finished chan bool) {
		for {
			time.Sleep(time.Second)
			fmt.Println("-")
			if time.Now().After(completed) {
				finished <- true
			}
		}
	}(finished)
	<- finished
}
