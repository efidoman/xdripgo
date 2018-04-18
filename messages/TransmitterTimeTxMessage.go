package messages

import (
	"github.com/efidoman/xdripgo/packet"
)

type TransmitterTimeTxMessage struct {
	Opcode byte
	Data   []byte
}

func NewTransmitterTimeTxMessage() TransmitterTimeTxMessage {
	var m TransmitterTimeTxMessage

	m.Opcode = 0x24
	d := make([]byte, 1)
	d[0] = m.Opcode
	d = packet.AppendCrc16(d)
	m.Data = d
	return m
}

// TODO: packet format
// +--------+----------------+
// | [0]    | [1-2]          |
// +--------+----------------+
// | opcode |                |
// +--------+----------------+
// | 28     |                |
// +--------+----------------+
