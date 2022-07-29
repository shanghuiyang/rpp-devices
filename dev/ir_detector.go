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

// IRDetector implements Detector interface
type IRDetector struct {
	out machine.Pin
}

// NewIRDetector ...
func NewIRDetector(out uint8) *IRDetector {
	ir := &IRDetector{
		out: machine.Pin(out),
	}
	ir.out.Configure(machine.PinConfig{Mode: machine.PinInput})
	return ir
}

// Detected ...
func (ir *IRDetector) Detected() bool {
	return ir.out.Get()
}
