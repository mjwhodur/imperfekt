package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mjwhodur/imperfekt/commands/escp2"
	"github.com/mjwhodur/imperfekt/commands/escp2/charmaps"
	"golang.org/x/text/encoding/charmap"
)

func main() {
	surname := "Michał Hodur\n"
	addr := "Very secret Address\n"
	addline2 := "Przemyśl, Woj. Podkarpackie\n"

	dev, err := os.OpenFile("/dev/usb/lp2", os.O_RDWR, 0)

	if err != nil {
		log.Fatal(err)
	}
	encoder := charmap.CodePage852.NewEncoder()
	_, err = dev.WriteString(string(escp2.AssignCharacterTable(0, charmaps.PC852)))
	_, err = dev.WriteString(string(escp2.SelectCharacterTable(0)))
	if err != nil {
		panic(err)
	}
	for _, text := range []string{surname, addr, addline2} {

		tr, _ := encoder.String(text)
		write, err := dev.WriteString(tr)
		if err != nil {
			return
		}
		fmt.Println("Written bytes: ", write)
	}

}
