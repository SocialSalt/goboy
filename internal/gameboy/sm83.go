package gameboy

type SM83 struct {
	Clock     int32
	Registers Register
	Bus       byte
	Cart      GBCart
	Memory    Memory
}

var instructions map[byte]func() = map[byte]func(){}

func (cpu *SM83) nop() {}
