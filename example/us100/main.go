package main

import (
	"fmt"
	"time"

	"github.com/shanghuiyang/rpp-devices/dev"
)

func main() {
	// new a us100 using GPIO interface
	us100, err := dev.NewUS100GPIO(12, 13)
	// or using UART interface
	// us100, err := dev.NewUS100UART(12, 13, 9600)
	if err != nil {
		panic(err)
	}

	for {
		dist, err := us100.Dist()
		if err != nil {
			fmt.Printf("error: %v\n", err)
			time.Sleep(time.Millisecond * 10)
			continue
		}
		fmt.Printf("dist: %.2fcm\n", dist)
		time.Sleep(time.Millisecond * 10)
	}
}
