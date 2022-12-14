package main

import (
	"time"

	"github.com/shanghuiyang/rpp-devices/dev"
)

var (
	btn = dev.NewButtonImp(26)
	led = dev.NewLedImp(28)
	buz = dev.NewBuzzerImp(27)
)

func main() {
	// us100, err := dev.NewUS100UART(12, 13, 9600)
	us100, err := dev.NewUS100GPIO(12, 13)
	if err != nil {
		panic(err)
	}

	for {
		if btn.Pressed() {
			alert()
			continue
		}

		dist, err := us100.Dist()
		if err == nil && dist < 100 {
			alert()
			continue
		}
		time.Sleep(time.Millisecond * 10)
	}

}

func alert() {
	go led.On()
	go buz.On()
	time.Sleep(time.Millisecond * 100)

	go led.Off()
	go buz.Off()
	time.Sleep(time.Millisecond * 500)
}
