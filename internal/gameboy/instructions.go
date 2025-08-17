package gameboy

import "fmt"

type Instruction struct {
	// Name  string // Mnemonic of the instruction
	// Len   int    // length in bytes
	// Dur   int    // duration
	// Flags []rune // flags affected

	InstrType InstructionType
	AddrMode  AddressingMode
	RegType1  RegisterType
	RegType2  RegisterType
	CondType  ConditionType
	param     byte
}

func (i *Instruction) LogState() string {
	if i == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%+v", i)
}

type InstructionType int

const (
	I_NONE InstructionType = iota
	I_NOP
	I_LD
	I_INC
	I_DEC
	I_RLCA
	I_ADD
	I_RRCA
	I_STOP
	I_RLA
	I_JR
	I_RRA
	I_DAA
	I_CPL
	I_SCR
	I_CCF
	I_HALT
	I_ADC
	I_SUB
	I_SBC
	I_AND
	I_XOR
	I_OR
	I_CP
	I_POP
	I_JP
	I_PUSH
	I_RET
	I_CB
	I_CALL
	I_RETI
	I_LDH
	I_JPHL
	I_DI
	I_EI
	I_RST
	I_ERR
	// CB instructions
	I_RLC
	I_RRC
	I_RL
	I_RR
	I_SLA
	I_SRA
	I_SWAP
	I_SRL
	I_BIT
	I_RES
	I_SET
)

var instructionTypeNames = []string{
	"I_NONE",
	"I_NOP",
	"I_LD",
	"I_INC",
	"I_DEC",
	"I_RLCA",
	"I_ADD",
	"I_RRCA",
	"I_STOP",
	"I_RLA",
	"I_JR",
	"I_RRA",
	"I_DAA",
	"I_CPL",
	"I_SCR",
	"I_CCF",
	"I_HALT",
	"I_ADC",
	"I_SUB",
	"I_SBC",
	"I_AND",
	"I_XOR",
	"I_OR",
	"I_CP",
	"I_POP",
	"I_JP",
	"I_PUSH",
	"I_RET",
	"I_CB",
	"I_CALL",
	"I_RETI",
	"I_LDH",
	"I_JPHL",
	"I_DI",
	"I_EI",
	"I_RST",
	"I_ERR",
	// CB instructions
	"I_RLC",
	"I_RRC",
	"I_RL",
	"I_RR",
	"I_SLA",
	"I_SRA",
	"I_SWAP",
	"I_SRL",
	"I_BIT",
	"I_RES",
	"I_SET",
}

type AddressingMode int

const (
	A_IMP AddressingMode = iota
	A_R_D16
	A_R_R
	A_MR_R
	A_R
	A_R_D8
	A_R_MR
	A_R_HLI
	A_AM_R_HLD
	A_HLI_R
	A_HLD_R
	A_R_A8
	A_A8_R
	A_HL_SPR
	A_D16
	A_D8
	A_D16_R
	A_MR_D8
	A_MR
	A_A16_R
	A_R_A16
)

type RegisterType int

const (
	R_NONE RegisterType = iota
	R_A
	R_F
	R_B
	R_C
	R_D
	R_E
	R_H
	R_L
	R_AF
	R_BC
	R_DE
	R_HL
	R_SP
	R_PC
)

type ConditionType int

const (
	C_NONE ConditionType = iota
	C_NZ
	C_Z
	C_NC
	C_C
)

var instructions = [0x100]Instruction{
	0x00: {InstrType: I_NOP},
	0x05: {InstrType: I_DEC, AddrMode: A_R, RegType1: R_B},
	0x0E: {InstrType: I_LD, AddrMode: A_R_D8, RegType1: R_C},
	// 0x31: {InstrType: I_LD},
	0x76: {InstrType: I_HALT},
	0xAF: {InstrType: I_XOR, AddrMode: A_R, RegType1: R_A},
	0xC3: {InstrType: I_JP, AddrMode: A_D16},
	0xF3: {InstrType: I_DI},
	0xFB: {InstrType: I_EI},
}
