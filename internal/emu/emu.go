package emu

import (
	"log"
	"time"

	"github.com/SocialSalt/goboy/internal/cart"
	"github.com/SocialSalt/goboy/internal/cpu"
)

type EmuContext struct {
	Running bool
	Paused  bool
	Ticks   uint64
}

func NewEmuContext() EmuContext {
	return EmuContext{false, false, 0}
}

func Run(cartPath string) {

	cart, err := cart.LoadCart(cartPath)
	if err != nil {
		log.Fatal("failed to load cart %s", err)
	}

	ctx := NewEmuContext()
	ctx.Running = true
	ctx.Ticks = 0

	cpu, err := cpu.NewCPU()
	if err != nil {
		log.Fatal("Failed to initalize cpu %s", err)
	}

	for ctx.Running {
		if ctx.Paused {
			time.Sleep(10 * time.Millisecond)
			continue
		}
		cpu.Step()

		ctx.Ticks++
	}
}
