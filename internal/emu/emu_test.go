package emu_test

import (
	"log/slog"
	"testing"

	"github.com/SocialSalt/goboy/internal/emu"
	"github.com/stretchr/testify/require"
)

func TestEmu(t *testing.T) {
	err := emu.Run("/Users/ethanwalker/Library/Application Support/OpenEmu/Game Library/roms/Game Boy/Pokemon - Crystal Version (USA, Europe) (Rev 1).gbc")
	require.NoError(t, err)
}

func TestEmu2(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	err := emu.Run("../../data/cpu_instrs/cpu_instrs.gb")
	require.NoError(t, err)
}
