/*
LedImp is a led module.

Connect to Raspberry Pico:
  - gcc(the longer pin):  any data pin
  - gnd(the shorter pin): any gnd pin
*/
package dev

import (
	"machine"
	"time"
)

// LedImp implements Led interface
type LedImp struct {
	pin machine.Pin
}

// NewLedImp ...
func NewLedImp(pin uint8) *LedImp {
	led := &LedImp{
		pin: machine.Pin(pin),
	}
	led.pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led.pin.Low()
	return led
}

// On ...
func (led *LedImp) On() {
	led.pin.High()
}

// Off ...
func (led *LedImp) Off() {
	led.pin.Low()
}

// Blink is let led blink n time, interval Millisecond each time
func (led *LedImp) Blink(n int, intervalMs int) {
	d := time.Duration(intervalMs)
	for i := 0; i < n; i++ {
		led.On()
		delayMs(d)
		led.Off()
		delayMs(d)
	}
}
