package emu

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/SocialSalt/goboy/internal/gameboy"
)

type EmuContext struct {
	Running bool
	Paused  bool
	Ticks   uint64
}

func NewEmuContext() *EmuContext {
	return &EmuContext{false, false, 0}
}

func Run(cartPath string) error {

	cart, err := gameboy.LoadCart(cartPath)
	if err != nil {
		slog.Info(fmt.Sprintf("failed to load cart with error: %v", err))
		return fmt.Errorf("failed to load cart %w", err)
	}

	ctx := NewEmuContext()
	ctx.Running = true
	ctx.Ticks = 0

	cpu := gameboy.NewCPU(&cart)

	for ctx.Running {
		if ctx.Paused {
			time.Sleep(10 * time.Millisecond)
			continue
		}
		err := cpu.Step()
		if err != nil {
			return fmt.Errorf("Error while taking cpu step: %w", err)
		}

		ctx.Ticks++
	}
	return nil
}
