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
		MotorA: newMiniL298NMotor(in1, in2),
		MotorB: newMiniL298NMotor(in3, in4),
	}
	return l
}

// minil298nMotor implements MotorDriver interface
type minil298nMotor struct {
	in1 machine.Pin
	in2 machine.Pin
}

func newMiniL298NMotor(in1, in2 uint8) *minil298nMotor {
	m := &minil298nMotor{
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
func (m *minil298nMotor) Forward() {
	m.in1.High()
	m.in2.Low()
}

// Backward ...
func (m *minil298nMotor) Backward() {
	m.in1.Low()
	m.in2.High()
}

// Stop ...
func (m *minil298nMotor) Stop() {
	m.in1.Low()
	m.in2.Low()
}

// SetSpeed ...
func (m *minil298nMotor) SetSpeed(percent uint) {

}
