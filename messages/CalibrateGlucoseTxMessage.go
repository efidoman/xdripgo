package messages

import (
	"encoding/binary"
	"github.com/efidoman/xdripgo/packet"
)

type CalibrateGlucoseTxMessage struct {
	Opcode byte
	Data   []byte
}

func NewCalibrateGlucoseTxMessage(glucose uint16, timestamp uint32) CalibrateGlucoseTxMessage {
	var m CalibrateGlucoseTxMessage

	m.Opcode = 0x34

	d := make([]byte, 7)
	d[0] = m.Opcode

	g := make([]byte, 2)
	t := make([]byte, 4)

	binary.LittleEndian.PutUint16(g, glucose)
	binary.LittleEndian.PutUint32(t, timestamp)

	copy(d[1:3], g)
	copy(d[3:7], t)

	d = packet.AppendCrc16(d)
	m.Data = d

	return m
}

// +--------+---------+----------------------+-------+
// | [0]    | [1-2]   | [3-6]                | [7-8] |
// +--------+---------+----------------------+-------+
// | opcode | glucose | dexcomTimeInSeconds  | CRC   |
// +--------+---------+----------------------+-------+
// | 34     | cb 00   | 35 20 00 00          | b3 f3 |
// +--------+---------+----------------------+-------+
