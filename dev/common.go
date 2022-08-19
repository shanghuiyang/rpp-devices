package dev

import (
	"errors"
	"machine"
	"time"
)

type (
	LogicLevel    int
	InterfaceType int
)

type StepperMode int
type MotorName int

const (
	MotorA MotorName = iota
	MotorB
	MotorC
	MotorD
	MotorE
	MotorF
)

const (
	FullMode StepperMode = iota
	HalfMode
	QuarterMode
	EighthMode
	SixteenthMode
)

const (
	// voice speed in cm/s
	voiceSpeed = 34000.0
)

const (
	Low  LogicLevel = 0
	High LogicLevel = 1
)

const (
	GPIO InterfaceType = iota
	I2C
	SPI
	UART
	USB
)

func getUart(txPin, rxPin uint8) (*machine.UART, error) {
	tx := machine.Pin(txPin)
	rx := machine.Pin(rxPin)

	if tx == machine.GP0 && rx == machine.GP1 {
		return machine.UART0, nil
	}

	if tx == machine.GP12 && rx == machine.GP13 {
		return machine.UART0, nil
	}

	if tx == machine.GP16 && rx == machine.GP17 {
		return machine.UART0, nil
	}

	if tx == machine.GP4 && rx == machine.GP5 {
		return machine.UART1, nil
	}

	if tx == machine.GP8 && rx == machine.GP9 {
		return machine.UART1, nil
	}
	return nil, errors.New("invalid tx & rx pin for uart interface")
}

func delayNs(d time.Duration) {
	time.Sleep(d * time.Nanosecond)
}

func delayUs(d time.Duration) {
	time.Sleep(d * time.Microsecond)
}

func delayMs(d time.Duration) {
	time.Sleep(d * time.Millisecond)
}

func delaySec(d time.Duration) {
	time.Sleep(d * time.Second)
}

func delayMin(d time.Duration) {
	time.Sleep(d * time.Minute)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
