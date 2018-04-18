package messages

import (
	"github.com/efidoman/xdripgo/packet"
)

type VersionRequestTxMessage struct {
	Opcode byte
	Data   []byte
}

func NewVersionRequestTxMessage() VersionRequestTxMessage {
	var m VersionRequestTxMessage

	m.Opcode = 0x4a
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
// | 4a     |                |
// +--------+----------------+
