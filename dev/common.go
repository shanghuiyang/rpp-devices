package dev

import (
	"time"
)

type (
	LogicLevel    int
	InterfaceType int
)

type StepperMode int

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
