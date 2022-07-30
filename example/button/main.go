package main

import (
	"fmt"
	"time"

	"github.com/shanghuiyang/rpp-devices/dev"
)

const btnPin = 26

func main() {
	btn := dev.NewButtonImp(btnPin)
	for {
		if !btn.Pressed() {
			time.Sleep(10 * time.Millisecond)
			continue
		}
		fmt.Printf("button is pressed")
		time.Sleep(500 * time.Millisecond)

	}
}
