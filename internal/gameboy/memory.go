package gameboy

import (
	"errors"
	"fmt"
)

type Memory struct {
	Cart GBCart
	RAM  [0x100]byte

	// there are 2 banks of VRAM each is 0x2000
	VRAM     [0x4000]byte
	VRAMBank byte

	// there are 9 banks of WRAM each is 0x1000
	WRAM     [0x9000]byte
	WRAMBank byte

	// interrupt enable register
	IE byte

	// object attribute memory
	OAM [0x100]byte

	IO GameboyHardware
}

func (m *Memory) Read(addr uint16) (byte, error) {
	switch {
	case addr < 0x8000:
		// read from the rom
		// 0x0000 - 0x3FFF is the main 16KiB ROM bank
		// 0x4000 - 0x7FFF is the 16KiB bank on the cartidge
		return m.Cart.Read(addr)

	case addr >= 0x8000 && addr < 0xA000:
		// read from VRAM
		offset := uint16(m.VRAMBank)*0x2000 - 0x8000
		return m.VRAM[addr+offset], nil

	case addr >= 0xA000 && addr < 0xC000:
		// cart RAM
		return m.Cart.Read(addr)

	case addr >= 0xC000 && addr < 0xD000:
		// work RAM bank 0
		return m.WRAM[addr-0xC000], nil

	case addr >= 0xD000 && addr < 0xE000:
		// work RAM bank 1 - 7
		offset := uint16(m.WRAMBank)*0x1000 - 0xC000
		return m.WRAM[addr+offset], nil

	case addr >= 0xE000 && addr < 0xFDFF:
		// Echo RAM, prohibited
		// TODO: implement what is described in docs
		// https://gbdev.io/pandocs/Memory_Map.html#echo-ram
		return 0, errors.New("attempted read from echo memory")

	case addr >= 0xFE00 && addr < 0xFEA0:
		// Object Attribute Memory
		return m.OAM[addr-0xFE00], nil

	case addr >= 0xFEA0 && addr < 0xFF00:
		// Not usable
		return 0, fmt.Errorf("memory address not usable: %X", addr)

	case addr >= 0xFF00 && addr < 0xFF80:
		// hardware I/O registers
		return m.IO.Read(addr)

	case addr >= 0xFF80 && addr < 0xFFFF:
		// High Ram
		return m.ReadHighRam(addr)

	case addr == 0xFFFF:
		// Interrupt Enable Register
		return m.IE, nil

	default:
		return 0, fmt.Errorf("undefined memory address read: %X", addr)
	}
}

func (m *Memory) ReadHighRam(addr uint16) (byte, error) {
	return 0, nil
}

func (m *Memory) Write(addr uint16, value byte) error {
	switch {
	case addr < 0x8000:
		// read from the rom
		// 0x0000 - 0x3FFF is the main 16KiB ROM bank
		// 0x4000 - 0x7FFF is the 16KiB bank on the cartidge
		return m.Cart.Write(addr, value)

	case addr >= 0x8000 && addr < 0xA000:
		// read from VRAM
		offset := uint16(m.VRAMBank)*0x2000 - 0x8000
		m.VRAM[addr+offset] = value

	case addr >= 0xA000 && addr < 0xC000:
		// cart RAM
		return m.Cart.Write(addr, value)

	case addr >= 0xC000 && addr < 0xD000:
		// work RAM bank 0
		m.WRAM[addr-0xC000] = value

	case addr >= 0xD000 && addr < 0xE000:
		// work RAM bank 1 - 7
		offset := uint16(m.WRAMBank)*0x1000 - 0xC000
		m.WRAM[addr+offset] = value

	case addr >= 0xE000 && addr < 0xFDFF:
		// Echo RAM, prohibited
		// TODO: implement what is described in docs
		// https://gbdev.io/pandocs/Memory_Map.html#echo-ram
		return errors.New("attempted read from echo memory")

	case addr >= 0xFE00 && addr < 0xFEA0:
		// Object Attribute Memory
		m.OAM[addr-0xFE00] = value

	case addr >= 0xFEA0 && addr < 0xFF00:
		// Not usable
		return fmt.Errorf("memory address not usable: %X", addr)

	case addr >= 0xFF00 && addr < 0xFF80:
		// hardware I/O registers
		return m.IO.Write(addr, value)

	case addr >= 0xFF80 && addr < 0xFFFF:
		// High Ram
		return m.WriteHighRam(addr, value)

	case addr == 0xFFFF:
		// Interrupt Enable Register
		m.IE = value

	default:
		return fmt.Errorf("undefined memory address read: %X", addr)
	}
	return nil
}

func (m *Memory) WriteHighRam(uint16, byte) error {
	return nil
}
