/*
IRDetector is a sensor used to detected infrared ray.

Connect to Pico:
 - vcc: any 3.3v pin
 - gnd: any gnd pin
 - out: any data pin

*/
package dev

import (
	"machine"
)

// LightDetector implements Detector interface
type LightDetector struct {
	do machine.Pin
}

// NewLightDetector ...
func NewLightDetector(out uint8) *LightDetector {
	ld := &LightDetector{
		do: machine.Pin(out),
	}
	ld.do.Configure(machine.PinConfig{Mode: machine.PinInput})
	return ld
}

// Detected ...
func (ld *LightDetector) Detected() bool {
	return !ld.do.Get()
}
