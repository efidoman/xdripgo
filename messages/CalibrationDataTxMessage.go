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

	/*
		crc := packet.MarshalUint16(packet.Crc16(d))

		e := make([]byte, 3)

		e[0] = d[0]
		copy(e[1:3], crc)
		m.Data = e
	*/
	m.Data = d

	return m
}

// NOTE: this is a total guess
// +--------+--------------------------+------------------------+--------+
// | [0]    | [1-4]                    | [5-8]                  | [9-10] |
// +--------+--------------------------+------------------------+--------+
// | opcode | dexcomStartTimeInSeconds | dexcomEndTimeInSeconds | CRC    |
// +--------+--------------------------+------------------------+--------+
// | 50     | 9e 32 66 00              | ce 5c 66 00            | 87 77  |
// +--------+--------------------------+------------------------+--------+