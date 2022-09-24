package main

import (
	"time"
)

func delayMs(d time.Duration) {
	time.Sleep(d * time.Millisecond)
}

func delaySec(d time.Duration) {
	time.Sleep(d * time.Second)
}
