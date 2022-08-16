package main

import (
	"github.com/shanghuiyang/rpp-devices/dev"
)

const (
	l298nIn1 = 7
	l298nIn2 = 22
	l298nIn3 = 10
	l298nIn4 = 11
	sg90Pin  = 16
	trigPin  = 20
	echoPin  = 21
	ledPin   = 25
)

const (
	turnAngle      = 30
	turnTimeSec    = 2
	backwardDistCM = 20
	turnDistCM     = 40
)

const (
	forward  operator = "forward"
	backward operator = "backward"
	stop     operator = "stop"
	turn     operator = "turn"
)

var (
	thecar    *car
	distMeter dev.DistanceMeter
)

type operator string

func main() {
	l298n := dev.NewL298N(l298nIn1, l298nIn2, l298nIn3, l298nIn4)
	sg90, err := dev.NewSG90(sg90Pin)
	if err != nil {
		return
	}
	us100, err := dev.NewUS100GPIO(trigPin, echoPin)
	if err != nil {
		return
	}
	thecar = &car{
		l298n: l298n,
		sg90:  sg90,
	}
	distMeter = us100

	var (
		fwd  = false
		op   = forward
		chOp = make(chan operator, 4)
	)

	sg90.Roll(0)
	delayMs(100)
	for {
		select {
		case p := <-chOp:
			op = p
		default:
			// 	do nothing
		}

		switch op {
		case backward:
			thecar.Backward()
			delaySec(2)
			chOp <- stop
			continue
		case stop:
			fwd = false
			thecar.Stop()
			delayMs(20)
			chOp <- turn
			continue
		case turn:
			fwd = false
			thecar.Turn(turnAngle, turnTimeSec)
			chOp <- forward
			continue
		case forward:
			if !fwd {
				fwd = true
				thecar.Forward()
				go lookingForObs(chOp)
			}
			delayMs(5)
			continue
		}
	}
}

func lookingForObs(chOp chan operator) {
	for {
		d, err := distMeter.Dist()
		for i := 0; err != nil && i < 3; i++ {
			delayMs(100)
			d, err = distMeter.Dist()
		}
		if err != nil {
			continue
		}

		if d < backwardDistCM {
			chOp <- stop
			chOp <- backward
			return
		}
		if d < turnDistCM {
			chOp <- stop
			return
		}
	}
}
