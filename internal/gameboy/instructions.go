package gameboy

type Instruction struct {
	Name  string // Mnemonic of the instruction
	Len   int    // length in bytes
	Dur   int    // duration
	Flags []rune // flags affected
}

var instrs = [0x100]func(*SM83){
	0x00: func(cpu *SM83) {
		cpu.nop()
	},
	0x01: func(cpu *SM83) {
	},
	0x02: func(cpu *SM83) {
	},
	0x03: func(cpu *SM83) {
	},
	0x04: func(cpu *SM83) {
	},
	0x05: func(cpu *SM83) {
	},
	0x06: func(cpu *SM83) {
	},
	0x07: func(cpu *SM83) {
	},
}
