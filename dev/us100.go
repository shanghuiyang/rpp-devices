/*
US-100 is an ultrasonic distance meter used to measure the distance to objects.
US-100 supports both of interfaces: GPIO and UART.
min distance: 2cm
max distance: 450cm

Connect to Raspberry Pico:
GPIO Interface:
 - VCC: any 3.3v or 5v pin
 - GND: any gnd pin
 - Trig: any gnd pin
 - Echo: any gnd pin

 UART Interface:
 - VCC: any 3.3v or 5v pin
 - GND: any gnd pin
 - ...............................................
 - !!! NOTE: TX->TXD, RX-RXD, NOT TX->RXD, RX-TXD
 - ...............................................
 - TX: must connect to GP0 or GP4 or GP8 or GP16
 - RX: must connect to GP1 or GP5 or GP9 or GP17
*/
package dev

import (
	"time"

	"machine"
)

const (
	us100Timeout = 14000000 // Nanosecond, 476m
	us100MaxDist = 450      // cm
)

var (
	trigData byte = 0x55
)

// US100 ...
type US100 struct {
	iface InterfaceType

	// ttl mode
	trig machine.Pin
	echo machine.Pin

	// uart mode
	uart *machine.UART
}

// NewUS100GPIO creates US100 using GPOI interface
func NewUS100GPIO(trig, echo uint8) (*US100, error) {
	us := &US100{
		iface: GPIO,
		trig:  machine.Pin(trig),
		echo:  machine.Pin(echo),
	}
	us.trig.Configure(machine.PinConfig{Mode: machine.PinOutput})
	us.trig.Low()
	us.echo.Configure(machine.PinConfig{Mode: machine.PinInput})
	return us, nil
}

// NewUS100UART creates US100 using UART interface
func NewUS100UART(txPin, rxPin uint8, baud uint32) (*US100, error) {
	uart, err := getUart(txPin, rxPin)
	if err != nil {
		return nil, err
	}

	uart.Configure(machine.UARTConfig{
		BaudRate: baud,
		TX:       machine.Pin(txPin),
		RX:       machine.Pin(rxPin),
	})
	return &US100{
		iface: UART,
		uart:  uart,
	}, nil
}

// Value returns the distance in cm
func (us *US100) Dist() (float64, error) {
	if us.iface == UART {
		return us.distFromUART()
	}
	return us.distFromGPIO()
}

func (us *US100) distFromUART() (float64, error) {

	// send trigger data
	err := us.uart.WriteByte(trigData)
	if err != nil {
		return 0, err
	}

	// read data
	d0, err := us.uart.ReadByte()
	if err != nil {
		return 0, err
	}
	d1, err := us.uart.ReadByte()
	if err != nil {
		return 0, err
	}
	// calc distance in cm
	return float64((uint16(d0)<<8)|uint16(d1)) / 10.0, nil
}

func (us *US100) distFromGPIO() (float64, error) {
	us.trig.Low()
	delayUs(1)
	us.trig.High()
	delayUs(1)

	for i := 0; !us.echo.Get(); i++ {
		if i >= us100Timeout {
			return us100MaxDist, nil
		}
		delayNs(1)
	}

	t := time.Now()
	for i := 0; us.echo.Get(); i++ {
		if i >= us100Timeout {
			return us100MaxDist, nil
		}
		delayNs(1)
	}

	dist := time.Since(t).Seconds() * voiceSpeed / 2.0
	return dist, nil
}
