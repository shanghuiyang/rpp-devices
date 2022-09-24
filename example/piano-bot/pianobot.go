package main

import (
	"github.com/shanghuiyang/rpp-devices/dev"
)

type pianobot struct {
	s1 dev.ServoMotor
	s2 dev.ServoMotor
	s3 dev.ServoMotor
	s4 dev.ServoMotor
}

func newPianoBot(s1, s2, s3, s4 dev.ServoMotor) *pianobot {
	return &pianobot{
		s1: s1,
		s2: s2,
		s3: s3,
		s4: s4,
	}
}

func (pb *pianobot) play(i uint) {
	switch i {
	case 0:
		pb.play0()
	case 1:
		pb.play1()
	case 2:
		pb.play2()
	case 3:
		pb.play3()
	case 4:
		pb.play4()
	case 5:
		pb.play5()
	case 6:
		pb.play6()
	case 7:
		pb.play7()
	case 8:
		pb.play8()
	}
}

func (pb *pianobot) reset() {
	pb.s1.Roll(0)
	delayMs(200)
	pb.s2.Roll(0)
	delayMs(200)
	pb.s3.Roll(0)
	delayMs(200)
	pb.s4.Roll(0)
	delayMs(200)
}

func (pb *pianobot) play0() {
	delayMs(300)
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
