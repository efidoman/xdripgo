package messages

import (
	"github.com/efidoman/xdripgo/packet"
	"time"
)

type SessionStartTxMessage struct {
	Opcode byte
	Data   []byte
}

func NewSessionStartTxMessage(startTime uint32) SessionStartTxMessage {
	var m SessionStartTxMessage

	m.Opcode = 0x26

	d := make([]byte, 9)
	d[0] = m.Opcode

	tstart := packet.MarshalUint32(startTime)
	copy(d[1:5], tstart)

	now := time.Now()
	secs := now.Unix()
	tnow := packet.MarshalUint32(uint32(secs))
	copy(d[6:9], tnow)

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
// | 26     |                |
// +--------+----------------+
