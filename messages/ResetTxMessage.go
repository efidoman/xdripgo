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

	/*
		crc := packet.MarshalUint16(packet.Crc16(d))
		copy(d[1:3], crc)
	*/
	m.Data = d

	return m
}

// packet format for ResetTxMessage
// +--------+----------------+
// | [0]    | [1-2]          |
// +--------+----------------+
// | opcode | CRC            |
// +--------+----------------+
// | 42     | cd 73          |
// +--------+----------------+
