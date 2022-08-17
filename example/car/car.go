package main

import (
	"time"

	"github.com/shanghuiyang/rpp-devices/dev"
)

type car struct {
	l298n *dev.L298N
	sg90  *dev.SG90
}

func (c *car) Forward() {
	c.l298n.MotorAForward()
}

func (c *car) Backward() {
	c.l298n.MotorABackward()
}

func (c *car) Stop() {
	c.l298n.MotorAStop()
}

func (c *car) Turn(angle float64, turnTimeMs int) {
	c.sg90.Roll(angle * (-1))
	c.Backward()
	delayMs(time.Duration(turnTimeMs))
	c.sg90.Roll(0)
}
