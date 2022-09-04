package main

import (
	"time"

	"github.com/shanghuiyang/rpp-devices/dev"
)

const (
	btnPin = 26
	ledPin = 25
)

func main() {
	btn := dev.NewButtonImp(btnPin)
	led := dev.NewLedImp(ledPin)
	for {
		if !btn.Pressed() {
			time.Sleep(10 * time.Millisecond)
			continue
		}
		led.Blink(1, 200)
		time.Sleep(500 * time.Millisecond)
	}
}
