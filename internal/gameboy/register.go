package gameboy

import (
	"fmt"
)

type Register struct {
	AF uint16
	BC uint16
	DE uint16
	HL uint16
	SP uint16
	PC uint16
}

func NewRegister() *Register {
	return &Register{PC: 0x100, AF: 0x0100}
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

func (r *Register) Write(rt RegisterType, value uint16) {
	switch rt {
	case R_A:
		r.AF = SetHi(r.AF, byte(value))
	case R_F:
		r.AF = SetLo(r.AF, byte(value))
	case R_B:
		r.BC = SetHi(r.BC, byte(value))
	case R_C:
		r.BC = SetLo(r.BC, byte(value))
	case R_D:
		r.DE = SetHi(r.DE, byte(value))
	case R_E:
		r.DE = SetLo(r.DE, byte(value))
	case R_H:
		r.HL = SetHi(r.HL, byte(value))
	case R_L:
		r.HL = SetLo(r.HL, byte(value))
	case R_AF:
		r.AF = value
	case R_BC:
		r.BC = value
	case R_DE:
		r.DE = value
	case R_HL:
		r.HL = value
	case R_SP:
		r.SP = value
	case R_PC:
		r.PC = value
	}
}

type number interface {
	uint16 | byte
}

func GetBit[T number](value T, bit int) bool {
	return value&(1<<bit) != 0
}

func SetBit[T number](n T, bit int, value bool) T {
	if value {
		return n | (1 << bit)
	}
	return n & ^(1 << bit)
}

func SetHi(n uint16, v byte) uint16 {
	return (uint16(v) << 8) | (n & 0xFF)
}

func SetLo(n uint16, v byte) uint16 {
	return uint16(v) | (n & 0xFF00)
}

// zero flag is 7th bit of AF
func (r *Register) ZFlag() bool {
	return GetBit(r.AF, 7)
}

// subtraction flag is 6th bit of AF
func (r *Register) NFlag() bool {
	return GetBit(r.AF, 6)
}

// half carry flag is 5th bit of AF
func (r *Register) HFlag() bool {
	return GetBit(r.AF, 5)
}

// carry flag is 4th bit of AF
func (r *Register) CFlag() bool {
	return GetBit(r.AF, 4)
}

func (r *Register) SetZFlag(value bool) {
	r.AF = SetBit(r.AF, 7, value)
}

func (r *Register) SetCFlag(value bool) {
	r.AF = SetBit(r.AF, 4, value)
}

func (r *Register) SetFlags(z bool, n bool, h bool, c bool) {
	r.AF = SetBit(r.AF, 7, z)
	r.AF = SetBit(r.AF, 6, n)
	r.AF = SetBit(r.AF, 5, h)
	r.AF = SetBit(r.AF, 4, c)
}
