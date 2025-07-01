package gameboy_test

import (
	"testing"

	"github.com/SocialSalt/goboy/internal/cart"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

var nintendoLogo []byte = []byte{
	0xCE, 0xED, 0x66, 0x66, 0xCC, 0x0D, 0x00, 0x0B, 0x03, 0x73, 0x00, 0x83, 0x00, 0x0C, 0x00, 0x0D, 0x00, 0x08, 0x11, 0x1F, 0x88, 0x89, 0x00, 0x0E, 0xDC, 0xCC, 0x6E, 0xE6, 0xDD, 0xDD, 0xD9, 0x99, 0xBB, 0xBB, 0x67, 0x63, 0x6E, 0x0E, 0xEC, 0xCC, 0xDD, 0xDC, 0x99, 0x9F, 0xBB, 0xB9, 0x33, 0x3E,
}

func TestGBCartReadPokemonCrystal(t *testing.T) {
	cartPath := "/Users/ethanwalker/Library/Application Support/OpenEmu/Game Library/roms/Game Boy/Pokemon - Crystal Version (USA, Europe) (Rev 1).gbc"

	loadedCart, err := cart.LoadCart(cartPath)

	if err != nil {
		t.Fatalf("Failed to load cart data %s", err)
	}
	if diff := cmp.Diff(nintendoLogo, loadedCart.Header.NintendoLogo); diff != "" {
		t.Errorf("Failed to load nintendo logo. (-want +got) %s", diff)
	}

	if diff := cmp.Diff("PM_CRYSTAL\x00BYTE\xc0", string(loadedCart.Header.Title)); diff != "" {
		t.Errorf("Failed to load the title correctly. (-want +got) %s", diff)
	}

	if diff := cmp.Diff([]byte{0x42, 0x59, 0x54, 0x45}, loadedCart.Header.ManufacturerCode); diff != "" {
		t.Errorf("Failed to load the manufacturer code correctly. (-want +got) %s", diff)
	}

	if diff := cmp.Diff(byte(0xC0), loadedCart.Header.CGBFlag); diff != "" {
		t.Errorf("Failed to load the cgb flag correctly. (-want +got) %s", diff)
	}

	if diff := cmp.Diff("01", string(loadedCart.Header.NewLicenseeCode)); diff != "" {
		t.Errorf("Failed to load the new lisencsee code code correctly. (-want +got) %s", diff)
	}

	if diff := cmp.Diff(byte(0x00), loadedCart.Header.SGBFlag); diff != "" {
		t.Errorf("Failed to load the sgb flag correctly. (-want +got) %s", diff)
	}

	if diff := cmp.Diff(byte(0x10), loadedCart.Header.CartrigeType); diff != "" {
		t.Errorf("Failed to load the cartrige type correctly. (-want +got) %s", diff)
	}

	if diff := cmp.Diff(byte(0x06), loadedCart.Header.ROMSize); diff != "" {
		t.Errorf("Failed to load the rom size correctly. (-want +got) %s", diff)
	}

	if diff := cmp.Diff(byte(0x03), loadedCart.Header.RAMSize); diff != "" {
		t.Errorf("Failed to load the ram size correctly. (-want +got) %s", diff)
	}

	if diff := cmp.Diff(byte(0x33), loadedCart.Header.OldLicenseeCode); diff != "" {
		t.Errorf("Failed to load the old licensee code correctly. (-want +got) %s", diff)
	}

	if diff := cmp.Diff(byte(0x01), loadedCart.Header.MaskRomVersionNumber); diff != "" {
		t.Errorf("Failed to load the mask rom version number correctly. (-want +got) %s", diff)
	}

	if diff := cmp.Diff(byte(0x26), loadedCart.Header.HeaderChecksum); diff != "" {
		t.Errorf("Failed to load the header checksum correctly. (-want +got) %s", diff)
	}

	if diff := cmp.Diff([]byte{0x18, 0xD2}, loadedCart.Header.GlobalChecksum); diff != "" {
		t.Errorf("Failed to load the global checksum correctly. (-want +got) %s", diff)
	}
}

func TestGBCartReadLinksAwakening(t *testing.T) {
	cartPath := "/Users/ethanwalker/Library/Application Support/OpenEmu/Game Library/roms/Game Boy/Legend of Zelda, The - Link's Awakening (USA, Europe) (Rev 2).gb"

	loadedCart, err := cart.LoadCart(cartPath)

	if err != nil {
		t.Fatalf("Failed to load cart data %s", err)
	}
	if diff := cmp.Diff(nintendoLogo, loadedCart.Header.NintendoLogo); diff != "" {
		t.Errorf("Failed to load nintendo logo. (-want +got) %s", diff)
	}

	if diff := cmp.Diff("ZELDA\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00", string(loadedCart.Header.Title)); diff != "" {
		t.Errorf("Failed to load the title correctly. (-want +got) %s", diff)
	}

	if diff := cmp.Diff([]byte{0x00, 0x00, 0x00, 0x00}, loadedCart.Header.ManufacturerCode); diff != "" {
		t.Errorf("Failed to load the manufacturer code correctly. (-want +got) %s", diff)
	}

	if diff := cmp.Diff(byte(0x00), loadedCart.Header.CGBFlag); diff != "" {
		t.Errorf("Failed to load the cgb flag correctly. (-want +got) %s", diff)
	}

	if diff := cmp.Diff([]byte{0x0, 0x0}, loadedCart.Header.NewLicenseeCode); diff != "" {
		t.Errorf("Failed to load the new lisencsee code code correctly. (-want +got) %s", diff)
	}

	if diff := cmp.Diff(byte(0x00), loadedCart.Header.SGBFlag); diff != "" {
		t.Errorf("Failed to load the sgb flag correctly. (-want +got) %s", diff)
	}

	if diff := cmp.Diff(byte(0x03), loadedCart.Header.CartrigeType); diff != "" {
		t.Errorf("Failed to load the cartrige type correctly. (-want +got) %s", diff)
	}

	if diff := cmp.Diff(byte(0x04), loadedCart.Header.ROMSize); diff != "" {
		t.Errorf("Failed to load the rom size correctly. (-want +got) %s", diff)
	}

	if diff := cmp.Diff(byte(0x02), loadedCart.Header.RAMSize); diff != "" {
		t.Errorf("Failed to load the ram size correctly. (-want +got) %s", diff)
	}

	if diff := cmp.Diff(byte(0x01), loadedCart.Header.OldLicenseeCode); diff != "" {
		t.Errorf("Failed to load the old licensee code correctly. (-want +got) %s", diff)
	}

	if diff := cmp.Diff(byte(0x02), loadedCart.Header.MaskRomVersionNumber); diff != "" {
		t.Errorf("Failed to load the mask rom version number correctly. (-want +got) %s", diff)
	}

	if diff := cmp.Diff(byte(0x6A), loadedCart.Header.HeaderChecksum); diff != "" {
		t.Errorf("Failed to load the header checksum correctly. (-want +got) %s", diff)
	}

	if diff := cmp.Diff([]byte{0x3A, 0xEE}, loadedCart.Header.GlobalChecksum); diff != "" {
		t.Errorf("Failed to load the global checksum correctly. (-want +got) %s", diff)
	}
}

func TestHeaderChecksum(t *testing.T) {
	cartPath := "/Users/ethanwalker/Library/Application Support/OpenEmu/Game Library/roms/Game Boy/Legend of Zelda, The - Link's Awakening (USA, Europe) (Rev 2).gb"

	loadedCart, err := cart.LoadCart(cartPath)
	require.NoError(t, err, "cart failed to load")
	checksum := cart.ComputeHeaderChecksum(loadedCart)
	require.Equal(t, loadedCart.Header.HeaderChecksum, checksum)

	cartPath = "/Users/ethanwalker/Library/Application Support/OpenEmu/Game Library/roms/Game Boy/Pokemon - Crystal Version (USA, Europe) (Rev 1).gbc"

	loadedCart, err = cart.LoadCart(cartPath)
	require.NoError(t, err, "cart failed to load")
	checksum = cart.ComputeHeaderChecksum(loadedCart)
	require.Equal(t, loadedCart.Header.HeaderChecksum, checksum)
}

func TestGlobalChecksum(t *testing.T) {
	cartPath := "/Users/ethanwalker/Library/Application Support/OpenEmu/Game Library/roms/Game Boy/Legend of Zelda, The - Link's Awakening (USA, Europe) (Rev 2).gb"

	loadedCart, err := cart.LoadCart(cartPath)
	require.NoError(t, err, "cart failed to load")
	checksum := cart.ComputeGlobalChecksum(loadedCart)
	// require.Equal(t, loadedCart.Header.GlobalChecksum, checksum)

	cartPath = "/Users/ethanwalker/Library/Application Support/OpenEmu/Game Library/roms/Game Boy/Pokemon - Crystal Version (USA, Europe) (Rev 1).gbc"

	loadedCart, err = cart.LoadCart(cartPath)
	require.NoError(t, err, "cart failed to load")
	checksum = cart.ComputeGlobalChecksum(loadedCart)
	require.Equal(t, loadedCart.Header.GlobalChecksum, checksum)
}
