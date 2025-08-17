package gameboy

import "fmt"

type Register struct {
	AF uint16
	BC uint16
	DE uint16
	HL uint16
	SP uint16
	PC uint16
}

func NewRegister() *Register {
	return &Register{PC: 0x100, AF: 0x100}
}

func (r *Register) LogState() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%+v", r)
}

func (r *Register) Read(rt RegisterType) uint16 {
	// val >> 8 gets the high byte by shifting the
	// value 8 bits to the right
	// val & 0xFF gets the low byte by anding with
	// something that only has low values
	switch rt {
	case R_A:
		return r.AF >> 8
	case R_F:
		return r.AF & 0xFF
	case R_B:
		return r.BC >> 8
	case R_C:
		return r.BC & 0xFF
	case R_D:
		return r.DE >> 8
	case R_E:
		return r.DE & 0xFF
	case R_H:
		return r.HL >> 8
	case R_L:
		return r.HL & 0xFF
	case R_AF:
		return r.AF
	case R_BC:
		return r.BC
	case R_DE:
		return r.DE
	case R_HL:
		return r.HL
	case R_SP:
		return r.SP
	case R_PC:
		return r.PC
	}
	return 0
}
