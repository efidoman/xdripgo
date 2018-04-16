package messages

import (
	"encoding/binary"
	"github.com/efidoman/xdripgo/mathutils"
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

	crc := mathutils.MarshalUint16(mathutils.Crc16(d))

	e := make([]byte, 3)

        e[0] = d[0]
	copy(e[1:3], crc)
	m.Data = e

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
