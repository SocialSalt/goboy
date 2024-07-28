package cpu

// Rn: register number; Op2: 32bit operand; Rd: destination register
type ArmCpu interface {
	Tick()
	ParseInstruction()
	ADC(Rd int, Rn int, Op2 [4]byte, Carry int) int // add with carry
	ADD(Rd int, Rn int, Op2 int) int                // TODO: not sure that the types are correct
	AND()
	B()   // branch
	BIC() // bit clear
	BL()  // branch with link
	BX()  // branch and exchange
	CDP() // coprocessor data processing
	CMN() // compare negative
	CMP() // compare
	EOR() // exclusive or
	LDC() // load coprocessor from memory
	LDM() // load multiple registers
	LDR() // load register from memory
	MCR() // move cpu register to coprocessor register
}

type ARM7TDMI struct {
	Reg    [16]Register
	Mem    int // FIXME: define memory
	Cycles uint64
}
