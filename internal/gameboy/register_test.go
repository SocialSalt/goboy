package gameboy_test

import (
	"testing"

	"github.com/SocialSalt/goboy/internal/gameboy"
	"github.com/stretchr/testify/require"
)

func TestGetBit(t *testing.T) {
	x := byte(0b0001)
	result := gameboy.GetBit(x, 0)
	require.True(t, result)
	result = gameboy.GetBit(x, 1)
	require.False(t, result)

	x = byte(0b1001)
	result = gameboy.GetBit(x, 0)
	require.True(t, result)
	result = gameboy.GetBit(x, 3)
	require.True(t, result)

	x = byte(0b0100_1000)
	result = gameboy.GetBit(x, 0)
	require.False(t, result)
	result = gameboy.GetBit(x, 1)
	require.False(t, result)
	result = gameboy.GetBit(x, 2)
	require.False(t, result)
	result = gameboy.GetBit(x, 3)
	require.True(t, result)
	result = gameboy.GetBit(x, 4)
	require.False(t, result)
	result = gameboy.GetBit(x, 5)
	require.False(t, result)
	result = gameboy.GetBit(x, 6)
	require.True(t, result)
	result = gameboy.GetBit(x, 7)
	require.False(t, result)
	// test past the end of the data
	result = gameboy.GetBit(x, 8)
	require.False(t, result)
	result = gameboy.GetBit(x, 9)
	require.False(t, result)

	y := uint16(0b0000_0010_0100_1000)
	result = gameboy.GetBit(y, 0)
	require.False(t, result)
	result = gameboy.GetBit(y, 1)
	require.False(t, result)
	result = gameboy.GetBit(y, 2)
	require.False(t, result)
	result = gameboy.GetBit(y, 3)
	require.True(t, result)
	result = gameboy.GetBit(y, 4)
	require.False(t, result)
	result = gameboy.GetBit(y, 5)
	require.False(t, result)
	result = gameboy.GetBit(y, 6)
	require.True(t, result)
	result = gameboy.GetBit(y, 7)
	require.False(t, result)
	result = gameboy.GetBit(y, 8)
	require.False(t, result)
	result = gameboy.GetBit(y, 9)
	require.True(t, result)
}

func TestSetBit(t *testing.T) {
	// byte
	x := byte(0b0001_0001)
	x = gameboy.SetBit(x, 2, true)
	require.Equal(t, byte(0b0001_0101), x)
	x = gameboy.SetBit(x, 3, true)
	require.Equal(t, byte(0b0001_1101), x)
	x = gameboy.SetBit(x, 7, true)
	require.Equal(t, byte(0b1001_1101), x)

	x = gameboy.SetBit(x, 2, false)
	require.Equal(t, byte(0b1001_1001), x)
	x = gameboy.SetBit(x, 3, false)
	require.Equal(t, byte(0b1001_0001), x)
	x = gameboy.SetBit(x, 7, false)
	require.Equal(t, byte(0b0001_0001), x)

	// setting outside range is fine
	x = byte(0b0001_0101)
	x = gameboy.SetBit(x, 8, true)
	require.Equal(t, byte(0b10101), x)

	x = gameboy.SetBit(x, 8, false)
	require.Equal(t, byte(0b10101), x)

	// uint16
	y := uint16(0b0001_0001_0001_0001)
	y = gameboy.SetBit(y, 9, true)
	require.Equal(t, uint16(0b0001_0011_0001_0001), y)
	y = gameboy.SetBit(y, 14, true)
	require.Equal(t, uint16(0b0101_0011_0001_0001), y)

	y = gameboy.SetBit(y, 5, false)
	require.Equal(t, uint16(0b0101_0011_0001_0001), y)
	y = gameboy.SetBit(y, 0, false)
	require.Equal(t, uint16(0b0101_0011_0001_0000), y)
}

func TestSetHi(t *testing.T) {
	v := uint16(0xABCD)
	x := byte(0x12)

	v = gameboy.SetHi(v, x)
	require.Equal(t, uint16(0x12CD), v)
}

func TestSetLo(t *testing.T) {
	v := uint16(0xABCD)
	x := byte(0x12)

	v = gameboy.SetLo(v, x)
	require.Equal(t, uint16(0xAB12), v)
}
