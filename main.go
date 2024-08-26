package main

import (
	"fmt"

	. "github.com/paramsiddharth/esp32-morse/morse"
)

func main() {
	for {
		text := "SOS Param SOS"
		fmt.Println("<<<", text)
		code := *StrToMorseCode(text)
		fmt.Print("\r...")
		BlinkMorseCode(&code)
		renderedCode := *RenderMorseCode(&code)
		fmt.Println("\r>>>", renderedCode)
		fmt.Print("\r")
	}
}
