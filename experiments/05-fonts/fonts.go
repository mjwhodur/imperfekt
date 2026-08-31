package main

import (
	"log"
	"os"

	"github.com/mjwhodur/imperfekt/commands/escp2"
)

func main() {
	dev, err := os.OpenFile("/dev/usb/lp2", os.O_WRONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	_, err = dev.WriteString(string(escp2.Select10andhalfPoint10CPI()))
	_, err = dev.WriteString("LQ 350 spacing\n")
	_, err = dev.WriteString(string([]byte{27, 48}))
	_, err = dev.WriteString("1/8 line spacing\n")
	_, err = dev.WriteString("1/8 line spacing\n")
	_, err = dev.WriteString(string([]byte{27, 50}))
	_, err = dev.WriteString("1/6 line spacing\n")
	_, err = dev.WriteString("1/8 line spacing\n")
	_, err = dev.WriteString(string([]byte{27, 43, 0}))
	_, err = dev.WriteString("no line spacing\n")
	_, err = dev.WriteString("no line spacing\n")
	_, err = dev.WriteString(string(escp2.Select10andhalfPoint12CPI()))
	_, err = dev.WriteString(string([]byte{27, 50}))
	_, err = dev.WriteString("LQ 350 spacing\n")
	_, err = dev.WriteString(string([]byte{27, 48}))
	_, err = dev.WriteString("1/8 line spacing\n")
	_, err = dev.WriteString("1/8 line spacing\n")
	_, err = dev.WriteString(string([]byte{27, 50}))
	_, err = dev.WriteString("1/6 line spacing\n")
	_, err = dev.WriteString("1/8 line spacing\n")
	_, err = dev.WriteString(string([]byte{27, 43, 0}))
	_, err = dev.WriteString("no line spacing\n")
	_, err = dev.WriteString("no line spacing\n")
	_, err = dev.WriteString(string(escp2.Select10andhalfPoint15CPI()))
	_, err = dev.WriteString(string([]byte{27, 50}))
	_, err = dev.WriteString("LQ 350 spacing\n")
	_, err = dev.WriteString(string([]byte{27, 48}))
	_, err = dev.WriteString("1/8 line spacing\n")
	_, err = dev.WriteString("1/8 line spacing\n")
	_, err = dev.WriteString(string([]byte{27, 50}))
	_, err = dev.WriteString("1/6 line spacing\n")
	_, err = dev.WriteString("1/8 line spacing\n")
	_, err = dev.WriteString(string([]byte{27, 43, 0}))
	_, err = dev.WriteString("no line spacing\n")
	_, err = dev.WriteString("no line spacing\n")

}
