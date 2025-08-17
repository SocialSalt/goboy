package gameboy

import (
	"context"
	"fmt"
	"log/slog"
)

type SM83 struct {
	Clock     uint32
	Registers *Register
	Bus       byte
	Memory    *Memory

	FetchedData    uint16
	MemDest        uint16
	CurOpcode      byte
	CurInstruction *Instruction

	halted bool
	step   bool
}

func NewCPU(cart *GBCart) *SM83 {
	return &SM83{
		Registers: NewRegister(),
		Memory:    NewMemory(cart),
	}
}

func (cpu *SM83) LogState() string {
	if cpu == nil {
		return "<nil>"
	}
	return fmt.Sprintf(
		"&{Clock: %v, Bus: %v, FetchedData: %v, CurOpcode: %X, CurInstruction: %v, Registers: %v}",
		cpu.Clock,
		cpu.Bus,
		cpu.FetchedData,
		cpu.CurOpcode,
		cpu.CurInstruction.LogState(),
		cpu.Registers.LogState(),
	)
}

func (cpu *SM83) IncrCycles(n int) {}

func first[T, U any](val T, _ U) T {
	return val
}

func (cpu *SM83) Step() error {
	if !cpu.halted {
		err := cpu.GetInstruction()
		if err != nil {
			slog.Error("cpu crashed while getting instruction", "err", err, "cpu state", cpu.LogState())
			return fmt.Errorf("cpu crashed while getting instruction: %w\ncpu state: %+v", err, cpu)
		}
		err = cpu.GetData()
		if err != nil {
			slog.Error("cpu crashed while getting data", "err", err, "cpu state", cpu.LogState())
			return fmt.Errorf("cpu crashed while getting data: %w\ncpu state: %+v", err, cpu)
		}

		slog.Debug(fmt.Sprintf(
			"prepare to execute:\n\tPC: 0x%04X\n\tinstr type & name: %v %v\n\topcode: 0x%X\n\tfetched data: 0x%X\n\tdata: 0x%X 0x%X",
			cpu.Registers.PC,
			cpu.CurInstruction.InstrType,
			instructionTypeNames[cpu.CurInstruction.InstrType],
			cpu.CurOpcode,
			cpu.FetchedData,
			first(cpu.Memory.Read(cpu.Registers.PC+1)),
			first(cpu.Memory.Read(cpu.Registers.PC+2)),
		))
		err = cpu.Execute()
		if err != nil {
			return err
		}
	}
	return nil
}

func (cpu *SM83) GetInstruction() error {
	opcode, err := cpu.Memory.Read(cpu.Registers.PC)
	slog.Log(context.TODO(), -8, "got opcode", "opcode", fmt.Sprintf("0x%X", opcode))
	if err != nil {
		return err
	}
	cpu.CurOpcode = opcode
	cpu.CurInstruction = &instructions[opcode]
	slog.Log(context.TODO(), -8, "got instr", "instr", cpu.CurInstruction)
	cpu.Registers.PC++
	return nil
}

func (cpu *SM83) GetData() error {
	inst := cpu.CurInstruction

	slog.Log(context.TODO(), -8, "getting data from memory", "addressing mode", inst.AddrMode)

	switch inst.AddrMode {
	case A_IMP:
		return nil

	case A_R:
		cpu.FetchedData = cpu.Registers.Read(inst.RegType1)
		return nil

	case A_R_D8:
		fetched, err := cpu.Memory.Read(cpu.Registers.PC)
		if err != nil {
			return err
		}
		cpu.IncrCycles(1)
		cpu.FetchedData = uint16(fetched)
		cpu.Registers.PC++

	case A_D16:
		fetched_lo, err := cpu.Memory.Read(cpu.Registers.PC)
		if err != nil {
			return err
		}
		cpu.IncrCycles(1)
		cpu.Registers.PC++

		fetched_hi, err := cpu.Memory.Read(cpu.Registers.PC)
		if err != nil {
			return err
		}
		cpu.IncrCycles(1)
		cpu.Registers.PC++

		cpu.FetchedData = uint16(fetched_lo) | (uint16(fetched_hi) << 8)

	default:
		return fmt.Errorf("Unknown addressing mode: %v", inst.AddrMode)
	}

	return nil
}

func (cpu *SM83) Execute() error {
	proc := instructionProcesses[cpu.CurInstruction.InstrType]
	if proc == nil {
		slog.Debug("failed to load instruction process", "CurInstruction", cpu.CurInstruction)
		return fmt.Errorf("failed to load instruction process %v", cpu.CurInstruction)
	}
	return proc(cpu)
}

func (cpu *SM83) CheckCondition() bool {
	z := cpu.Registers.ZFlag()
	c := cpu.Registers.CFlag()
	switch cpu.CurInstruction.CondType {
	case C_NONE:
		return true
	case C_C:
		return c
	case C_NC:
		return !c
	case C_Z:
		return z
	case C_NZ:
		return !z
	default:
		return false
	}
}
