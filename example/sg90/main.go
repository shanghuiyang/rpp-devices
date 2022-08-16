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

	// roll clockwise 30 degree
	sg90.Roll(30)
	time.Sleep(1 * time.Second)

	// roll to 0 degree
	sg90.Roll(0)
	time.Sleep(1 * time.Second)

	// roll anticlockwise 30 degree
	sg90.Roll(-30)
	time.Sleep(1 * time.Second)
}
