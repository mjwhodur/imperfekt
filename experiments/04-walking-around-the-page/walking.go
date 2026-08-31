package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mjwhodur/imperfekt/commands/escp2"
	"github.com/mjwhodur/imperfekt/commands/escp2/charmaps"
	"github.com/mjwhodur/imperfekt/commands/escp2/fonts"
)

func main() {

	dev, err := os.OpenFile("/dev/usb/lp3", os.O_WRONLY, 0)

	if err != nil {
		log.Fatal(err)
	}
	//encoder := charmap.CodePage852.NewEncoder()
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

	//for i, text := range []string{surname, addr, addline2, string([]byte{commandcodes.CR, commandcodes.LF})} {
	//	if i == 0 {
	//		dev.WriteString(string(escp2.SelectBoldFont()))
	//	} else {
	//		dev.WriteString(string(escp2.CancelBoldFont()))
	//	}
	//	fmt.Println("About to print:")
	//	fmt.Println(text)
	//	tr, encerr := encoder.String(text)
	//	if encerr != nil {
	//		fmt.Println("Errorr encoding text at index")
	//		fmt.Println(i)
	//	}
	//	write, e := dev.WriteString(tr)
	//	if e != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println("Written bytes: ", write)
	//}

	//_, err = dev.WriteString(string(escp2.Select10andhalfPoint10CPI()))
	fmt.Println("About to set position")
	//time.Sleep(5 * time.Second)

	_, err = dev.WriteString(string(escp2.Select10andhalfPoint10CPI()))
	if err != nil {
		fmt.Println("Errorr Selecting 10 and half point 15 CPI")
		fmt.Println(err)
	}
	_, err = dev.WriteString(string(escp2.SetDefaultUnit(10)))
	if err != nil {
		fmt.Println("Errorr Setting default unit")
		fmt.Println(err)
	}
	_, err = dev.WriteString(string(escp2.SetAbsoluteVerticalPrintPositionInMillimetersWithKnownMargin(70, 0)))
	if err != nil {
		fmt.Println("Errorr Setting Absolute Vertical Print Position")
		fmt.Println(err)
	}
	//if err != nil {
	//	fmt.Println("Errorr Setting Absolute Vertical Print Position")
	//	fmt.Println(err)
	//}
	//
	//_, _ = dev.WriteString("I should be somewhere aroung 5 centimeters from the top\n")

	//_, e := dev.WriteString(string(escp2.SetAbsoluteHorizontalPosition(208, 2)))
	_, e := dev.WriteString(string(escp2.SetAbsoluteHorizontalPrintPositionInMillimetersWithKnownMargin(70, 0)))
	if e != nil {
		fmt.Println("Errorr Setting Absolute Horizontal Print Position")
		fmt.Println(e)
	}

	_, _ = dev.Write(escp2.SelectBoldFont())
	_, _ = dev.WriteString("\n")
	_, e = dev.WriteString(string(escp2.SetAbsoluteHorizontalPrintPositionInMillimetersWithKnownMargin(70, 0)))
	_, _ = dev.Write(escp2.CancelBoldFont())
	_, _ = dev.WriteString("\n")
	_, e = dev.WriteString(string(escp2.SetAbsoluteHorizontalPrintPositionInMillimetersWithKnownMargin(70, 0)))
	_, _ = dev.WriteString("32-100 Krakow\n")

	//	dev.WriteString(string(escp2.SetAbsoluteHorizontalPrintPositionInMillimetersWithKnownMargin(50, 10))) // NIE DZIAŁA
	//	fmt.Println("About to print:")
	//	fmt.Println(text)
	//	tr, encerr := encoder.String(text)
	//	if encerr != nil {
	//		fmt.Println("Errorr encoding text at index")
	//		fmt.Println(i)
	//	}
	//	write, e := dev.WriteString(tr)
	//	if e != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println("Written bytes: ", write)
	//}
	//
	dev.WriteString(string(escp2.FormFeed()))

}
