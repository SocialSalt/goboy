package cart

import (
	"errors"
	"os"
)

type Cart interface{}

type GBCart struct {
	Data   []byte
	Header GBRomHeader
}

type GBACart struct {
	Data []byte
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

func (h *GBRomHeader) ComputeHeaderChecksum() {

}

func (h *GBRomHeader) ComputeGlobalChecksum() {

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
