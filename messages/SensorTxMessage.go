package messages

import (
	"github.com/efidoman/xdripgo/packet"
)

type SensorTxMessage struct {
	Opcode byte
	Data   []byte
}

func NewSensorTxMessage() SensorTxMessage {
	var m SensorTxMessage

	m.Opcode = 0x2e

	d := make([]byte, 1)
	d[0] = m.Opcode
	d = packet.AppendCrc16(d)
	m.Data = d
	return m
}
