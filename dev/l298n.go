/*
L298N is a motor driver used to control the direction and speed of DC motors.

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
 - EN1: must be one of GPIO 12, 13, 18 or 19 (pwn pins)
 - EN2: must be one of GPIO 12, 13, 18 or 19 (pwn pins)

*/
package dev

import (
	"machine"
)

// L298N implements MotorDriver interface
type L298N struct {
	in1 machine.Pin
	in2 machine.Pin
	in3 machine.Pin
	in4 machine.Pin
}

// NewL298N ...
func NewL298N(in1, in2, in3, in4 uint8) *L298N {
	l := &L298N{
		in1: machine.Pin(in1),
		in2: machine.Pin(in2),
		in3: machine.Pin(in3),
		in4: machine.Pin(in4),
	}
	l.in1.Configure(machine.PinConfig{Mode: machine.PinOutput})
	l.in2.Configure(machine.PinConfig{Mode: machine.PinOutput})
	l.in3.Configure(machine.PinConfig{Mode: machine.PinOutput})
	l.in4.Configure(machine.PinConfig{Mode: machine.PinOutput})
	l.in1.Low()
	l.in2.Low()
	l.in3.Low()
	l.in4.Low()
	return l
}

// Forward ...
func (l *L298N) MotorAForward() {
	l.in1.High()
	l.in2.Low()
}

// Backward ...
func (l *L298N) MotorABackward() {
	l.in1.Low()
	l.in2.High()
}

// Stop ...
func (l *L298N) MotorAStop() {
	l.in1.Low()
	l.in2.Low()
}

// Forward ...
func (l *L298N) MotorBForward() {
	l.in3.High()
	l.in4.Low()
}

// Backward ...
func (l *L298N) MotorBBackward() {
	l.in3.Low()
	l.in4.High()
}

// Stop ...
func (l *L298N) MotorBStop() {
	l.in3.Low()
	l.in4.Low()
}

// Speed ...
func (l *L298N) SetMotorASpeed(s uint32) {

}

// Speed ...
func (l *L298N) SetMotorBSpeed(s uint32) {

}
