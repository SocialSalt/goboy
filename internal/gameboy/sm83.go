package gameboy

import "log"

type SM83 struct {
	Clock     uint32
	Registers *Register
	Bus       byte
	Cart      *GBCart
	Memory    *Memory

	FetchData      uint16
	MemDest        uint16
	CurOpcode      byte
	CurInstruction *Instruction

	halted bool
	step   bool
}

func (cpu *SM83) Step() {
	if !cpu.halted {
		err := cpu.GetInstruction()
		if err != nil {
			log.Fatalf("cpu crashed while getting instruction: %v\ncpu state: %+v", err, cpu)
		}
		err = cpu.GetData()
		if err != nil {
			log.Fatalf("cpu crashed while getting data: %v\ncpu state: %+v", err, cpu)
		}

	}
}

func (cpu *SM83) GetInstruction() error {
	opcode, err := cpu.Memory.Read(cpu.Registers.PC)
	if err != nil {
		return err
	}
	cpu.CurOpcode = opcode
	cpu.CurInstruction = &instructions[opcode]
	cpu.Registers.PC++
	return nil
}

func (cpu *SM83) GetData() error {
	inst := cpu.CurInstruction

	switch inst.AddrMode {
	case A_IMP:
		return nil
	case A_R:
		cpu.FetchData = cpu.Registers.Read(inst.RegType1)
	case A_R_D8:
		fetched, err := cpu.Memory.Read(cpu.Registers.PC)
		if err != nil {
			return err
		}
		cpu.FetchData = uint16(fetched)
		cpu.Registers.PC++

	}

	return nil
}
