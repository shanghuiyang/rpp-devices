package main

import (
	"time"

	"github.com/shanghuiyang/rpp-devices/dev"
)

const ledPin = 28

func main() {
	led := dev.NewLedImp(ledPin)
	for {
		led.On()
		time.Sleep(time.Millisecond * 500)

		led.Off()
		time.Sleep(time.Millisecond * 500)
	}
}
