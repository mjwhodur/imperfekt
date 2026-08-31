package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mjwhodur/imperfekt/commands/escp2"
	"github.com/mjwhodur/imperfekt/commands/escp2/charmaps"
	"github.com/mjwhodur/imperfekt/commands/escp2/commandcodes"
	"github.com/mjwhodur/imperfekt/commands/escp2/fonts"
	"golang.org/x/text/encoding/charmap"
)

func main() {
	surname := "Michał Hodur\n"
	addr := "\n"
	addline2 := "Woj. Podkarpackie\n"

	dev, err := os.OpenFile("/dev/usb/lp2", os.O_WRONLY, 0)

	if err != nil {
		log.Fatal(err)
	}
	encoder := charmap.CodePage852.NewEncoder()
	_, err = dev.WriteString(string(escp2.InitializePrinter()))
	if err != nil {
		fmt.Println("Errorr initialazing printer")
		fmt.Println(err)
	}
	_, err = dev.WriteString(string(escp2.SetDefaultUnit(10))) // nie psuje
	_, err = dev.WriteString(string(escp2.SetPageFormatC6()))  // Już nie psuje
	if err != nil {
		fmt.Println("Errorr setting page format")
		fmt.Println(err)
	}
	//_, err = dev.WriteString(string(escp2.SetPageMarginsInMilimeters(30, 30, 50, 180)))
	_, err = dev.WriteString(string(escp2.AssignCharacterTable(0, charmaps.PC852)))
	if err != nil {
		fmt.Println("Errorr Assigning Character Table")
		fmt.Println(err)
	}
	_, err = dev.WriteString(string(escp2.SelectCharacterTable(0)))
	if err != nil {
		fmt.Println("Errorr Selecting Character Table")
		fmt.Println(err)
	}
	//if err != nil {
	//	panic(err)
	//}
	_, err = dev.WriteString(string(escp2.SelectFont(fonts.SansSerif)))
	//printData := []byte{commandcodes.ESC, '@', commandcodes.ESC, 'g', commandcodes.ESC, 'l', 10, commandcodes.ESC, 'Q', 75, 'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', 'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', 'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', 'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', 'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', 'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', 'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', 'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', 'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', 'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', 'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', 'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', 'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', 'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', 'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', 'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', 'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', 'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', 'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', commandcodes.CR, commandcodes.LF, commandcodes.FF}

	_, err = dev.WriteString(string(escp2.SelectCharacterTable(0)))

	_, err = dev.WriteString(string(escp2.Select10andhalfPoint15CPI()))
	for i, text := range []string{surname, addr, addline2, string([]byte{commandcodes.CR, commandcodes.LF})} {
		if i == 0 {
			dev.WriteString(string(escp2.SelectBoldFont()))
		} else {
			dev.WriteString(string(escp2.CancelBoldFont()))
		}
		fmt.Println("About to print:")
		fmt.Println(text)
		tr, encerr := encoder.String(text)
		if encerr != nil {
			fmt.Println("Errorr encoding text at index")
			fmt.Println(i)
		}
		write, e := dev.WriteString(tr)
		if e != nil {
			fmt.Println(err)
		}
		fmt.Println("Written bytes: ", write)
	}

	_, err = dev.WriteString(string(escp2.Select10andhalfPoint10CPI()))
	dev.WriteString(string(escp2.SetAbsoluteVerticalPrintPositionInMillimetersWithKnownMargin(50, 10))) // NIE DZIAŁA
	for i, text := range []string{"Piotr\n", "ul.  / __\n", "00-000 Kraków\n"} {
		if i == 0 {
			dev.WriteString(string(escp2.SelectBoldFont()))
		} else {
			dev.WriteString(string(escp2.CancelBoldFont()))
		}

		dev.WriteString(string(escp2.SetAbsoluteHorizontalPrintPositionInMillimetersWithKnownMargin(50, 10))) // NIE DZIAŁA
		fmt.Println("About to print:")
		fmt.Println(text)
		tr, encerr := encoder.String(text)
		if encerr != nil {
			fmt.Println("Errorr encoding text at index")
			fmt.Println(i)
		}
		write, e := dev.WriteString(tr)
		if e != nil {
			fmt.Println(err)
		}
		fmt.Println("Written bytes: ", write)
	}

	dev.WriteString(string(escp2.FormFeed()))

	//writ, err := dev.WriteString(string(printData))
	//fmt.Println("Wpisano: ", writ)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err = dev.Sync()
	//if err != nil {
	//	fmt.Println(err)
	//}

}
