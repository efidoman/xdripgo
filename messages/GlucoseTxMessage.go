package messages

import (
	"github.com/efidoman/xdripgo/packet"
)

type GlucoseTxMessage struct {
	Opcode byte
	Data   []byte
}

func NewGlucoseTxMessage() GlucoseTxMessage {
	var m GlucoseTxMessage

	m.Opcode = 0x30

	d := make([]byte, 1)
	d[0] = m.Opcode

	d = packet.AppendCrc16(d)

	/*
		crc := packet.MarshalUint16(packet.Crc16(d))
		copy(d[1:3], crc)
	*/
	m.Data = d

	return m
}
