package main

import (
	"time"

	"github.com/shanghuiyang/rpp-devices/dev"
)

const (
	trigPin = 16
	echoPin = 17
	ledPin  = 25
)

func main() {
	hcsr := dev.NewHCSR04(trigPin, echoPin)
	led := dev.NewLedImp(ledPin)
	for {
		dist, err := hcsr.Dist()
		if err != nil {
			time.Sleep(time.Millisecond * 100)
			continue
		}
		if dist > 10 {
			time.Sleep(time.Millisecond * 100)
			continue
		}
		led.Blink(1, 100)
	}
}
