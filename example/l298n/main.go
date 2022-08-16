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
	l298n := dev.NewL298N(in1, in2, in3, in4)
	l298n.MotorAForward()
	time.Sleep(3 * time.Second)
	l298n.MotorAStop()

	l298n.MotorBBackward()
	time.Sleep(3 * time.Second)
	l298n.MotorBStop()
}
