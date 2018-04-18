package messages

import (
	"github.com/efidoman/xdripgo/packet"
)

type CalibrationDataTxMessage struct {
	Opcode byte
	Data   []byte
}

func NewCalibrationDataTxMessage() CalibrationDataTxMessage {
	var m CalibrationDataTxMessage

	m.Opcode = 0x32

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
// | 32     | 11 16          |
// +--------+----------------+
