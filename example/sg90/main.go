package main

import (
	"time"

	"github.com/shanghuiyang/rpp-devices/dev"
)

const pin = 16

func main() {
	sg90, err := dev.NewSG90(pin)
	if err != nil {
		return
	}

	for i := -90; i <= 90; i += 10 {
		sg90.Roll(float64(i))
		time.Sleep(time.Millisecond * 1000)
	}
}
