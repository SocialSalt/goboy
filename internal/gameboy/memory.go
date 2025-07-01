package gameboy

import "errors"

type Memory struct {
	Cart GBCart
	RAM  [0x100]byte

	// there are 2 banks of VRAM each is 0x2000
	VRAM     [0x4000]byte
	VRAMBank byte

	// there are 9 banks of WRAM each is 0x1000
	WRAM     [0x9000]byte
	WRAMBank byte
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

	case addr >= 0xC000 && addr < 0xE000:
		// work RAM
		offset := uint16(m.WRAMBank)*0x1000 - 0xC000
		return m.WRAM[addr+offset], nil
	case addr >= 0xE000 && addr < 0xFDFF:
		// Echo RAM, prohibited
		// TODO: implement what is described in docs
		// https://gbdev.io/pandocs/Memory_Map.html#echo-ram
		return 0, errors.New("attempted read from echo memory")
	default:
		return 0, nil
	}
}

func (m *Memory) ReadRAM(addr uint16) byte {
	return m.RAM[addr]
}

func (m *Memory) WriteRAM(addr uint16, value byte) error {
	return nil
}

func (m *Memory) ReadVRAM(addr uint16) byte {
	return m.VRAM[addr]
}
