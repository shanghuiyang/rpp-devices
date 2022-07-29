package main

import (
	// "machine"
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

		dis, err := us100.Dist()
		if err == nil && dis < 100 {
			alert()
			continue
		}
		time.Sleep(time.Millisecond * 10)
	}

}

func alert() {
	led.On()
	buz.On()
	time.Sleep(time.Millisecond * 100)

	led.Off()
	buz.Off()
	time.Sleep(time.Millisecond * 500)
}
