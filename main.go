package main

import (
	"fmt"
	"time"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/cmd/instance/locks"
)

func main() {
	beatInfo := beat.Info{
		Beat:    "testing",
		Version: "9",
	}
	lock := locks.New(beatInfo)
	fmt.Printf("Hello, World\n")
	fmt.Printf("lock is %v\n", lock)
	err := lock.Lock()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Got Lock\n")
	time.Sleep(3 * time.Second)
	err = lock.Unlock()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Released Lock\n")
}
