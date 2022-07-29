/*
BuzzerImp is a buzzer module used to generate "beep, beep, ..." sound.

Connect to Raspberry Pico for a 3-pin buzzer mobule:
 - vcc: any 3.3v pin
 - gnd: any gnd pin
 - i/o: any data pin

Connect to Raspberry Pico for a 2-pin buzzer mobule:
  - vcc(the longer pin):  any data pin
  - gnd(the shorter pin): any gnd pin
*/

package dev

import (
	"machine"
	"time"
)

// BuzzerImp implements Buzzer interface
type BuzzerImp struct {
	pin machine.Pin
}

// NewBuzzerImp ...
func NewBuzzerImp(pin uint8) *BuzzerImp {
	b := &BuzzerImp{
		pin: machine.Pin(pin),
	}
	b.pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	b.pin.Low()
	return b
}

// On ...
func (b *BuzzerImp) On() {
	b.pin.High()
}

// Off ...
func (b *BuzzerImp) Off() {
	b.pin.Low()
}

// Beep beeps [n] times with an interval in [interval] millisecond
func (b *BuzzerImp) Beep(n int, intervalMs int) {
	d := time.Duration(intervalMs)
	for i := 0; i < n; i++ {
		b.On()
		delayMs(d)
		b.Off()
		delayMs(d)
	}
}
