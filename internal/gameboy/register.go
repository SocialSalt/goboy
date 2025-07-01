package gameboy

type Register struct {
	AF []byte
	B  byte
	C  byte
	D  byte
	E  byte
	H  byte
	L  byte
	SP uint16
	PC uint16
}
