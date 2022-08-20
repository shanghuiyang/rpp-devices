/*
L298N is a motor driver used to control the direction and speed of DC motors.

Spec:
           _________________________________________
          |                                         |
          |                                         |
    OUT1 -|                 L298N                   |- OUT3
    OUT2 -|                                         |- OUT4
          |                                         |
          |_________________________________________|
              |   |   |     |   |   |   |   |   |
             12v GND  5V   ENA IN1 IN2 IN3 IN4 ENB

Pins:
 - OUT1: dc motor A+
 - OUT2: dc motor A-
 - OUT3: dc motor B+
 - OUT4: dc motor B-

 - 12v: +battery
 - GND: -battery (and any gnd pin of raspberry pi if motors and raspberry pi use different battery sources)
 - IN1: any data pin
 - IN2: any data pin
 - IN3: any data pin
 - IN4: any data pin
 - EN1: any data pin
 - EN2: any data pin

*/
package dev

import (
	"machine"
)

// L298N ...
type L298N struct {
	MotorA MotorDriver
	MotorB MotorDriver
}

// l298nMotorDriver implements MotorDriver interface
type l298nMotorDriver struct {
	in1 machine.Pin
	in2 machine.Pin
	en  machine.Pin
	pwm PWM
	ch  uint8
}

// NewL298N ...
func NewL298N(in1, in2, in3, in4, ena, enb uint8) *L298N {
	l := &L298N{
		MotorA: newL298NMotorDriver(in1, in2, ena),
		MotorB: newL298NMotorDriver(in3, in4, enb),
	}
	return l
}

func newL298NMotorDriver(in1, in2, en uint8) *l298nMotorDriver {
	m := &l298nMotorDriver{
		in1: machine.Pin(in1),
		in2: machine.Pin(in2),
		en:  machine.Pin(en),
		pwm: machine.PWM0,
	}
	m.in1.Configure(machine.PinConfig{Mode: machine.PinOutput})
	m.in2.Configure(machine.PinConfig{Mode: machine.PinOutput})
	m.in1.Low()
	m.in2.Low()

	m.en.Configure(machine.PinConfig{Mode: machine.PinPWM})
	m.pwm.Configure(machine.PWMConfig{
		Period: pwmPeriod,
	})
	ch, err := m.pwm.Channel(m.en)
	if err != nil {
		return nil
	}
	m.ch = ch
	return m
}

// Forward ...
func (m *l298nMotorDriver) Forward() {
	m.in1.High()
	m.in2.Low()
}

// Backward ...
func (m *l298nMotorDriver) Backward() {
	m.in1.Low()
	m.in2.High()
}

// Stop ...
func (m *l298nMotorDriver) Stop() {
	m.in1.Low()
	m.in2.Low()
}

// Speed ...
func (m *l298nMotorDriver) SetSpeed(percent uint32) {
	if percent == 0 || percent > 100 {
		return
	}
	v := m.pwm.Top() / (100 / percent)
	m.pwm.Set(m.ch, v)
}
