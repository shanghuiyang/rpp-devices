package main

import (
	"time"

	"github.com/shanghuiyang/rpp-devices/dev"
)

const (
	keepLighting = time.Second * 60
)

var (
	led    = dev.NewLedImp(22)
	trigAt = time.Now()
)

func main() {
	ld := dev.NewLightDetector(18)
	us, err := dev.NewUS100GPIO(17, 16)
	if err != nil {
		panic(err)
	}

	go timeToTurnOffLight()
	for {
		if ld.Detected() {
			led.Off()
			time.Sleep(time.Millisecond * 10)
			continue
		}
		dist, err := us.Dist()
		if err != nil {
			time.Sleep(time.Millisecond * 10)
			continue
		}
		if dist > 50 {
			time.Sleep(time.Millisecond * 10)
			continue
		}
		led.On()
		trigAt = time.Now()
		time.Sleep(time.Second)
	}

}

func timeToTurnOffLight() {
	for {
		if time.Since(trigAt) > keepLighting {
			led.Off()
		}
		time.Sleep(time.Second)
	}
}
