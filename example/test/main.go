package main

import (
	"time"

	"github.com/shanghuiyang/rpp-devices/dev"
)

func main() {
	b := dev.NewButtonImp(26)
	led := dev.NewLedImp(28)
	buz := dev.NewBuzzerImp(27)
	for {
		if !b.Pressed() {
			time.Sleep(time.Millisecond * 10)
			continue
		}

		go led.On()
		go buz.On()
		time.Sleep(time.Millisecond * 100)

		go led.Off()
		go buz.Off()
		time.Sleep(time.Millisecond * 500)
	}
}
