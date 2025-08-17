package gameboy

type GameboyHardware struct {
}

func (h *GameboyHardware) Read(addr uint16) (byte, error) {
	return 0, nil
}

func (h *GameboyHardware) Write(addr uint16, value byte) error {
	return nil
}
