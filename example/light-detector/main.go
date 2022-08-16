package main

import (
	"time"

	"github.com/shanghuiyang/rpp-devices/dev"
)

const (
	ldPin  = 10
	ledPin = 25
)

func main() {
	ld := dev.NewLightDetector(ldPin)
	led := dev.NewLedImp(ledPin)
	for {
		time.Sleep(time.Millisecond * 100)
		if ld.Detected() {
			led.On()
			continue
		}
		led.Off()
	}
}
