package blink

import (
	"fmt"
	"machine"
	"time"
)

func Blink(dur uint32, frac float32, offset float32) {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	if frac+offset > 1 {
		return
	}

	blinkDur := uint32(frac * float32(dur))
	startDur := uint32(offset * float32(dur))
	remainingDur := dur - blinkDur - startDur

	time.Sleep(time.Duration(startDur) * time.Millisecond)
	led.High()
	fmt.Printf("\rOOO")
	time.Sleep(time.Duration(blinkDur) * time.Millisecond)
	led.Low()
	fmt.Printf("\r---")
	time.Sleep(time.Duration(remainingDur) * time.Millisecond)
}
