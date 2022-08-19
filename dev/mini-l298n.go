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

// MiniL298N implements MotorDriver interface
type MiniL298N struct {
	motorA Motor
	motorB Motor
}

// NewMiniL298N ...
func NewMiniL298N(in1, in2, in3, in4 uint8) *MiniL298N {
	l := &MiniL298N{
		motorA: NewDCMotor(in1, in2),
		motorB: NewDCMotor(in3, in4),
	}
	return l
}

// Forward ...
func (l *MiniL298N) Get(m MotorName) Motor {
	if m == MotorA {
		return l.motorA
	}
	if m == MotorB {
		return l.motorB
	}
	return nil
}
