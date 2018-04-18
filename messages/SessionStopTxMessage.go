package messages

import (
	"github.com/efidoman/xdripgo/packet"
)

type SessionStopTxMessage struct {
	Opcode byte
	Data   []byte
}

func NewSessionStopTxMessage(stopTime uint32) SessionStopTxMessage {
	var m SessionStopTxMessage

	m.Opcode = 0x28

	d := make([]byte, 5)
	d[0] = m.Opcode

	tstop := packet.MarshalUint32(stopTime)
	copy(d[1:5], tstop)

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
