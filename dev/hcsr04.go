/*
HC-SR04 is an ultrasonic distance meter used to measure the distance to objects.
min distance: 2cm
max distance: 600cm

Connect to Raspberry Pi:
  - vcc:	any 3.3v or 5v pin
  - gnd:	any gnd pin
  - trig:	any data pin
  - echo:	any data pin

*/
package dev

import (
	"time"

	"machine"
)

const (
	hcsr04Timeout = 18000000 // Nanosecond, 612m
	hcsr04MaxDist = 600      // cm
)

// HCSR04 implements DistanceMeter interface
type HCSR04 struct {
	trig machine.Pin
	echo machine.Pin
}

// NewHCSR04 ...
func NewHCSR04(trig int8, echo int8) *HCSR04 {
	hc := &HCSR04{
		trig: machine.Pin(trig),
		echo: machine.Pin(echo),
	}
	hc.trig.Configure(machine.PinConfig{Mode: machine.PinOutput})
	hc.trig.Low()
	hc.echo.Configure(machine.PinConfig{Mode: machine.PinInput})
	return hc
}

// Value returns distance in cm to objects
func (hc *HCSR04) Dist() (float64, error) {
	hc.trig.Low()
	delayUs(1)
	hc.trig.High()
	delayUs(1)

	for i := 0; !hc.echo.Get(); i++ {
		if i >= hcsr04Timeout {
			return hcsr04MaxDist, nil
		}
		delayNs(1)
	}

	t := time.Now()
	for i := 0; hc.echo.Get(); i++ {
		if i >= hcsr04Timeout {
			return hcsr04MaxDist, nil
		}
		delayNs(1)
	}

	dist := time.Since(t).Seconds() * voiceSpeed / 2.0
	return dist, nil
}
