package main

import (
	"github.com/shanghuiyang/rpp-devices/dev"
)

type pianobot struct {
	play [9]func()
	s1   dev.ServoMotor
	s2   dev.ServoMotor
	s3   dev.ServoMotor
	s4   dev.ServoMotor
}

func newPianoBot(s1, s2, s3, s4 dev.ServoMotor) *pianobot {
	pb := &pianobot{
		s1: s1,
		s2: s2,
		s3: s3,
		s4: s4,
	}
	pb.play = [9]func(){
		pb.play0,
		pb.play1,
		pb.play2,
		pb.play3,
		pb.play4,
		pb.play5,
		pb.play6,
		pb.play7,
		pb.play8,
	}
	return pb
}

func (pb *pianobot) play0() {
	delayMs(200)
}

func (pb *pianobot) play1() {
	pb.s1.Roll(-50)
	delayMs(200)
	pb.s1.Roll(0)
	delayMs(200)
}

func (pb *pianobot) play2() {
	pb.s1.Roll(50)
	delayMs(200)
	pb.s1.Roll(0)
	delayMs(200)
}

func (pb *pianobot) play3() {
	pb.s2.Roll(-50)
	delayMs(200)
	pb.s2.Roll(0)
	delayMs(200)
}

func (pb *pianobot) play4() {
	pb.s2.Roll(50)
	delayMs(200)
	pb.s2.Roll(0)
	delayMs(200)
}

func (pb *pianobot) play5() {
	pb.s3.Roll(-50)
	delayMs(200)
	pb.s3.Roll(0)
	delayMs(200)
}

func (pb *pianobot) play6() {
	pb.s3.Roll(50)
	delayMs(200)
	pb.s3.Roll(0)
	delayMs(200)
}

func (pb *pianobot) play7() {
	pb.s4.Roll(-70)
	delayMs(200)
	pb.s4.Roll(0)
	delayMs(200)
}

func (pb *pianobot) play8() {
	pb.s4.Roll(70)
	delayMs(200)
	pb.s4.Roll(0)
	delayMs(200)
}
