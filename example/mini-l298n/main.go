package main

import (
	"time"

	"github.com/shanghuiyang/rpp-devices/dev"
)

const (
	in1 = 14
	in2 = 15
	in3 = 12
	in4 = 13
)

func main() {
	l298n := dev.NewMiniL298N(in1, in2, in3, in4)
	l298n.Get(dev.MotorA).Forward()
	time.Sleep(3 * time.Second)
	l298n.Get(dev.MotorA).Stop()

	l298n.Get(dev.MotorB).Backward()
	time.Sleep(3 * time.Second)
	l298n.Get(dev.MotorB).Stop()
}
