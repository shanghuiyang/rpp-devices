/*
MiniL298N is a motor driver used to control the direction and speed of DC motors.

Pins:
 - MotorA+: dc motor A+
 - MotorA-: dc motor A-
 - MotorB+: dc motor B+
 - MotorB-: dc motor B-

 -  + : 3.3v/5v pin or battery(+)
 -  - : any gnd pin or battery(-)
 - IN1: any data pin
 - IN2: any data pin
 - IN3: any data pin
 - IN4: any data pin

*/
package dev

import (
	"machine"
)

// MiniL298N ...
type MiniL298N struct {
	MotorA MotorDriver
	MotorB MotorDriver
}

// NewMiniL298N ...
func NewMiniL298N(in1, in2, in3, in4 uint8) *MiniL298N {
	l := &MiniL298N{
		MotorA: newMiniL298NMotorDriver(in1, in2),
		MotorB: newMiniL298NMotorDriver(in3, in4),
	}
	return l
}

// minil298nMotor implements MotorDriver interface
type miniL298NMotorDriver struct {
	in1 machine.Pin
	in2 machine.Pin
}

func newMiniL298NMotorDriver(in1, in2 uint8) *miniL298NMotorDriver {
	m := &miniL298NMotorDriver{
		in1: machine.Pin(in1),
		in2: machine.Pin(in2),
	}
	m.in1.Configure(machine.PinConfig{Mode: machine.PinOutput})
	m.in2.Configure(machine.PinConfig{Mode: machine.PinOutput})
	m.in1.Low()
	m.in2.Low()
	return m
}

// Forward ...
func (m *miniL298NMotorDriver) Forward() {
	m.in1.High()
	m.in2.Low()
}

// Backward ...
func (m *miniL298NMotorDriver) Backward() {
	m.in1.Low()
	m.in2.High()
}

// Stop ...
func (m *miniL298NMotorDriver) Stop() {
	m.in1.Low()
	m.in2.Low()
}

// SetSpeed ...
func (m *miniL298NMotorDriver) SetSpeed(percent uint32) {

}
