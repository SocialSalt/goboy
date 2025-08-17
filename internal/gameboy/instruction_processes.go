package gameboy

import (
	"context"
	"errors"
	"log/slog"
)

type InstructionProcess func(cpu *SM83) error

func procNone(_ *SM83) error {
	return errors.New("encountered NONE instruction")
}

func procNOP(_ *SM83) error {
	slog.Log(context.TODO(), -8, "nop")
	return nil
}

func procLD(_ *SM83) error {
	return nil
}

func procJP(cpu *SM83) error {
	slog.Log(context.TODO(), -8, "jmp to", "FetchedData", cpu.FetchedData)
	if cpu.CheckCondition() {
		cpu.Registers.PC = cpu.FetchedData
	}
	return nil
}

func procDI(cpu *SM83) error {
	cpu.Memory.IME = false
	return nil
}

func procEI(cpu *SM83) error {
	cpu.Memory.IME = true
	return nil
}

func procXOR(cpu *SM83) error {
	val := cpu.Registers.Read(R_A) ^ (cpu.FetchedData & 0xFF)
	cpu.Registers.Write(R_A, val)
	cpu.Registers.SetFlags(cpu.Registers.Read(R_A) == 0, false, false, false)
	return nil
}

var instructionProcesses = [0x100]InstructionProcess{
	I_NONE: procNone,
	I_NOP:  procNOP,
	I_LD:   procLD,
	I_JP:   procJP,
	I_DI:   procDI,
	I_EI:   procEI,
	I_XOR:  procXOR,
}
