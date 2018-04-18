package messages

import (
	"github.com/efidoman/xdripgo/packet"
)

type ResetTxMessage struct {
	Opcode byte
	Data   []byte
}

func NewResetTxMessage() ResetTxMessage {
	var m ResetTxMessage

	m.Opcode = 0x42

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
// | 42     | cd 73          |
// +--------+----------------+
