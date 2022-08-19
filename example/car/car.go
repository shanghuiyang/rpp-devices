package main

import (
	"time"

	"github.com/shanghuiyang/rpp-devices/dev"
)

type car struct {
	driver dev.MotorDriver
	servo  dev.ServoMotor
}

func (c *car) Forward() {
	c.driver.Get(dev.MotorA).Forward()
}

func (c *car) Backward() {
	c.driver.Get(dev.MotorA).Backward()
}

func (c *car) Stop() {
	c.driver.Get(dev.MotorA).Stop()
}

func (c *car) Turn(angle float64, turnTimeMs int) {
	c.servo.Roll(angle * (-1))
	c.Backward()
	delayMs(time.Duration(turnTimeMs))
	c.servo.Roll(0)
}
