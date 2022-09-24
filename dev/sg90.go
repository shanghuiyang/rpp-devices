/*
SG90 is servo motor which can roll angels from 0~180 degree.

Connect to Raspberry Pi:
 - the red line:	3.3v or 5v pin
 - the brown line: 	gnd pin
 - the yellow line:	any data pin
*/
package dev

import (
	"machine"
)

const pwmPeriod = 20e6 // 20ms

// SG90 implements Motor interface
type SG90 struct {
	pin machine.Pin
	pwm PWM
	ch  uint8
}

// NewSG90 ...
func NewSG90(pin uint8, pwm PWM) (*SG90, error) {
	sg := &SG90{
		pin: machine.Pin(pin),
		pwm: pwm,
	}
	sg.pin.Configure(machine.PinConfig{Mode: machine.PinPWM})
	sg.pwm.Configure(machine.PWMConfig{
		Period: pwmPeriod,
	})
	ch, err := sg.pwm.Channel(sg.pin)
	if err != nil {
		return nil, err
	}
	sg.ch = ch
	return sg, nil
}

// Roll ...
// angle: [-90, 90]
// angle < 0: roll anticlockwise
// angel = 0: ahead
// angle > 0: roll clockwise
// e.g.
//
//       -30  0   30
//         \  |  /
//          \ | /
//           \|/
//   -90 ---- * ---- 90
//         +-----+
//         |     |
//         |     | sg90
//         | (*) |
//         +-----+
//
func (sg *SG90) Roll(angle float64) {
	if angle < -90 || angle > 90 {
		return
	}

	us := 1500 - angle*100/9 // angle to microseconds
	v := uint32(sg.pwm.Top()) * uint32(us) / (pwmPeriod / 1000)
	sg.pwm.Set(sg.ch, v)
	return
}

// SetSpeed ...
func (sg *SG90) SetSpeed(speed int) {
	// no implement
}
