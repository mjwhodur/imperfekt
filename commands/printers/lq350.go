package printers

// "Driver" For EPSON LQ-350

/*
	Commands supported by this printer:
	ESC @ 		escp2.InitializePrinter
	ESC U 		escp2.TurnUnidirectionalModeOnOff
	FF			FormFeed
	LF			LineFeed
	ESC 0		Select1_8_inchLineSpacing()
	ESC 2		Select1_6_inchLineSpacint()
	ESC 3		NOT IMPLEMENTED
	ESC +		SetN360thsinchLineSpacing
	CR			CarriageReturn()
	ESC ( C		SetPageLengthInDefinedUnit()
	ESC C		SetPageFormat
	ESC C 0		<<set page length in inches? Not sure>> <<maybe shall be ESC C NUL>>
	ESC Q		SetRightMargin()
	ESC I		SetLeftMargin
	ESC ( c		SetPageFormat()
	ESC N		SetBottomMargin()
	ESC O		CancelBottomMargin
	ESC $		SetAbsoluteHorizontalPrintPosition
	ESC \		SetRelativeHorizontalPrintPosition
	ESC (V		SetAbsoluteVerticalPrintPosition
	ESC (v		SetRelativeVerticalPrintPosition
	ESC D		SetHorizontalTabs
	HT			TabHorizontally()
	ESC B		Set Vertical Tabs
	VT			Tab Vertically
	ESC J		NOT IMPLEMENTED AdvancePrintPositionVertically
	ESC k		SelectTypeface()
	ESC x		SelectLQOrDraft()
	ESC y		<<UNKNOWN COMMAND>> High speed draft?!
	ESC X		SelectFontByPitchAndPoint
	ESC P		Select10CPI
	ESC M		Select12CPI
	ESC g		Select15CPI
	ESC p		Turn Proportional Mode On Off
	ESC 4		Select italic font
	ESC 5		Cancel italic font
	ESC E		Select Bold Font
	ESC F		Cancel Bold Font
	ESC !		Master Select
	ESC W		Turn Double Width printing on / off
	DC4			Cancel double width-printing one-line
	SO			Select double-width printing one-line
	DC2			Cancel Condensed Printing
	SI			Select condensed printing SelectCondensedPrinting()
	ESC w		Turn Double-height printing on/off
	ESC G		Select double-strike printing
	ESC H		Cancel double-strike printing
	ESC T		Cancel Superscript/subscript
	ESC S		Select superscript/subscript
	ESC -		Select underline on/off
	ESC ( -		Select line / score
	ESC q		<<SelectCharacterStyle>>
	ESC SPACE	<<SelectIntercharacterSpace>> C-108
	ESC c		<<Not implemented>> Select horizontal motion index HMI <- C-99
	ESC	( U		Set unit / set default unit
	ESC t		Select Character Table
	ESC ( t		Assign CharacterTable
	ESC R		<<NOT IMPLEMENTED>> Select international character set
	ESC %		<<NOT IMPLEMENTED>> Select iser-defined set
	ESC &		<<NOT IMPLEMENTED>> Define user-defined characters
	ESC :		<<NOT IMPLEMENTED>> Copy ROM to RAM
	ESC 6		<<NOT IMPLEMENTED>> Enable printing of upper-control-codes
	ESC 7		<<NOT IMPLEMENTED>> Enable upper control codes
	ESC (^		<<NOT IMPLEMENTED>> PrintDataAsCharacters
	ESC *		Select Bit Image
	ESC ( B		Bar Code Setup and print:
					EAN-13
					EAN-8
					2of5
					UPCA
					UPCE
					Code39
					Code128
					POSTNET
*/

type LQ350Printer struct {
	Printer
}

func NewLQ350Printer() *LQ350Printer {
	return &LQ350Printer{}
}

func (printer *LQ350Printer) Print(data byte) {}
