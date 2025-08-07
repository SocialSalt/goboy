package gameboy

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
)

type GBCart struct {
	Data   []byte
	Header GBRomHeader
}

type GBRomHeader struct {
	EntryPoint           []byte
	NintendoLogo         []byte
	Title                []byte
	ManufacturerCode     []byte
	CGBFlag              byte
	NewLicenseeCode      []byte
	SGBFlag              byte
	CartrigeType         byte
	ROMSize              byte
	RAMSize              byte
	DestinationCode      byte
	OldLicenseeCode      byte
	MaskRomVersionNumber byte
	HeaderChecksum       byte
	GlobalChecksum       []byte
}

func NewGBRomHeader(data []byte) GBRomHeader {
	header := GBRomHeader{
		EntryPoint:           data[0x100:0x104],
		NintendoLogo:         data[0x104:0x134],
		Title:                data[0x134:0x144],
		ManufacturerCode:     data[0x13F:0x143],
		CGBFlag:              data[0x143],
		NewLicenseeCode:      data[0x144:0x146],
		SGBFlag:              data[0x146],
		CartrigeType:         data[0x147],
		ROMSize:              data[0x148],
		RAMSize:              data[0x149],
		DestinationCode:      data[0x14A],
		OldLicenseeCode:      data[0x14B],
		MaskRomVersionNumber: data[0x14C],
		HeaderChecksum:       data[0x14D],
		GlobalChecksum:       data[0x14E:0x150],
	}

	return header
}

func (c *GBCart) Read(addr uint16) (byte, error) {

	return c.Data[addr], nil

}

func (c *GBCart) Write(addr uint16, value byte) error {
	return fmt.Errorf("no writing to rom")
}

func ComputeHeaderChecksum(cart GBCart) byte {
	var checksum byte = 0
	for _, d := range cart.Data[0x134:0x14D] {
		checksum = checksum + (^d)
	}
	return checksum
}

func ComputeGlobalChecksum(cart GBCart) []byte {
	var checksum uint16 = 0
	for i, d := range cart.Data {
		if i == 0x14E || i == 0x14F {
			continue
		}
		checksum = checksum + uint16(d)
	}
	ret := make([]byte, 2)
	binary.BigEndian.PutUint16(ret, checksum)
	return ret
}

func LoadCart(cartPath string) (GBCart, error) {
	fileInfo, err := os.Stat(cartPath)
	if errors.Is(err, os.ErrNotExist) {
		return GBCart{}, err
	}

	fp, err := os.Open(cartPath)
	if err != nil {
		return GBCart{}, err
	}

	data := make([]byte, fileInfo.Size())
	_, err = fp.Read(data)
	if err != nil {
		return GBCart{}, err
	}

	return GBCart{Data: data, Header: NewGBRomHeader(data)}, nil
}
