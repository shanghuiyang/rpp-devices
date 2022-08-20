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

// NewL298N ...
func NewL298N(in1, in2, in3, in4, ena, enb uint8) *L298N {
	l := &L298N{
		MotorA: newL298NMotor(in1, in2, ena),
		MotorB: newL298NMotor(in3, in4, enb),
	}
	return l
}

// l298nMotor implements MotorDriver interface
type l298nMotor struct {
	in1 machine.Pin
	in2 machine.Pin
	en  machine.Pin
	pwm PWM
	ch  uint8
}

func newL298NMotor(in1, in2, en uint8) *l298nMotor {
	m := &l298nMotor{
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
func (l *l298nMotor) Forward() {
	l.in1.High()
	l.in2.Low()
}

// Backward ...
func (l *l298nMotor) Backward() {
	l.in1.Low()
	l.in2.High()
}

// Stop ...
func (l *l298nMotor) Stop() {
	l.in1.Low()
	l.in2.Low()
}

// Speed ...
func (l *l298nMotor) SetSpeed(percent uint) {
	if percent == 0 || percent > 100 {
		return
	}
	v := uint32(100 / percent)
	l.pwm.Set(l.ch, v)
}
