package main

import (
	"machine"
	"time"

	"github.com/shanghuiyang/rpp-devices/dev"
)

const (
	sg90Pin1 = 0
	sg90Pin2 = 1
	sg90Pin3 = 4
	sg90Pin4 = 5
	btnPin   = 26
	ledPin   = 25
)

var bot *pianobot

func main() {
	btn := dev.NewButtonImp(btnPin)
	led := dev.NewLedImp(ledPin)

	s1, err := dev.NewSG90(sg90Pin1, machine.PWM0)
	if err != nil {
		led.Blink(3, 500)
		return
	}

	s2, err := dev.NewSG90(sg90Pin2, machine.PWM0)
	if err != nil {
		led.Blink(5, 500)
		return
	}

	s3, err := dev.NewSG90(sg90Pin3, machine.PWM2)
	if err != nil {
		led.Blink(7, 500)
		return
	}

	s4, err := dev.NewSG90(sg90Pin4, machine.PWM2)
	if err != nil {
		led.Blink(9, 500)
		return
	}

	led.Blink(3, 100)
	bot = newPianoBot(s1, s2, s3, s4)
	bot.reset()
	for {
		if !btn.Pressed() {
			delayMs(10)
			continue
		}
		led.Blink(1, 200)
		play()
	}
}

func play() {
	for _, song := range songs {
		for _, m := range song {
			bot.play(m[0])
			delayMs(time.Duration(m[1]))
		}
		delaySec(1)
	}
}
