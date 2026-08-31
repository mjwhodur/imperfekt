package charmaps

type Chartable int

const (
	Italic Chartable = iota
	PC437
	PC852
)

func GetTableCode(tab Chartable) []byte {
	switch tab {
	case PC852:
		return []byte{10, 0}
	default:
		panic("unhandled default case")
	}
}
