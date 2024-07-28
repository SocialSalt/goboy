package cpu

type SharpProcessor interface {
}

type SM83 struct {
	Clock     int32
	Registers [7]byte
	Bus       byte
}
