package main

import (
	"time"

	"github.com/shanghuiyang/rpp-devices/dev"
)

type car struct {
	motor dev.Motor
	servo dev.ServoMotor
}

func (c *car) Forward() {
	c.motor.Forward()
}

func (c *car) Backward() {
	c.motor.Backward()
}

func (c *car) Stop() {
	c.motor.Stop()
}

func (c *car) Turn(angle float64, turnTimeMs int) {
	c.servo.Roll(angle * (-1))
	c.Backward()
	delayMs(time.Duration(turnTimeMs))
	c.servo.Roll(0)
}
