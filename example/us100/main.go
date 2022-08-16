package main

import (
	"time"

	"github.com/shanghuiyang/rpp-devices/dev"
)

const (
	trigPin = 12
	echoPin = 13
	ledPin  = 25
)

func main() {
	// new a us100 using GPIO interface
	us100, err := dev.NewUS100GPIO(trigPin, echoPin)
	if err != nil {
		return
	}
	// or using UART interface
	// us100, err := dev.NewUS100UART(12, 13, 9600)

	led := dev.NewLedImp(ledPin)
	for {
		dist, err := us100.Dist()
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
