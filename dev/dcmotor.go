package dev

import (
	"machine"
)

type DCMotor struct {
	in1 machine.Pin
	in2 machine.Pin
}

// NewDCMotor ...
func NewDCMotor(in1, in2 uint8) *DCMotor {
	m := &DCMotor{
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
func (m *DCMotor) Forward() {
	m.in1.High()
	m.in2.Low()
}

// Backward ...
func (m *DCMotor) Backward() {
	m.in1.Low()
	m.in2.High()
}

// Stop ...
func (m *DCMotor) Stop() {
	m.in1.Low()
	m.in2.Low()
}

// Speed ...
func (m *DCMotor) SetSpeed(percent uint) {

}
