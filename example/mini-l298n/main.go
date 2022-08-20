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
	motorA := dev.NewDCMotor(l298n.MotorA)
	motorB := dev.NewDCMotor(l298n.MotorB)

	motorA.Forward()
	time.Sleep(3 * time.Second)
	motorA.Stop()

	motorB.Forward()
	time.Sleep(3 * time.Second)
	motorB.Stop()
}
