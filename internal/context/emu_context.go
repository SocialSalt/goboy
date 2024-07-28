package context

type EmuContext struct {
	Running bool
	Paused  bool
	Ticks   uint64
}

func NewEmuContext() EmuContext {
	return EmuContext{false, false, 0}
}
