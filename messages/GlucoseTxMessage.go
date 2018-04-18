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
	m.Data = d

	return m
}

// packet format
// +--------+----------------+
// | [0]    | [1-2]          |
// +--------+----------------+
// | opcode | CRC            |
// +--------+----------------+
// | 30     | 53 36          |
// +--------+----------------+
