package printers

type Printer struct {
	Address string
}

type GenericESCP2Printer struct {
	Printer
}

func (p *GenericESCP2Printer) Print(text string) {

}
