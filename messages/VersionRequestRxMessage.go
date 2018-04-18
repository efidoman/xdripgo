package messages

import (
	"fmt"
	"log"
)

type VersionRequestRxMessage struct {
	Opcode  byte
	Status  uint8
	Version string
}

func NewVersionRequestRxMessage(data []byte) VersionRequestRxMessage {
	var m VersionRequestRxMessage

	m.Opcode = 0x4b

	if len(data) != 19 {
		log.Print("Cannot create new VersionRequestRxMessage - Length not 19 bytes. Length = ", len(data))
	} else {
		if data[0] != m.Opcode {
			log.Print("Cannot create VersionRequestRxMessage - Response is not correct. Opcode should be = ", m.Opcode, " but data response Opcode is ", data[0])
		}
		m.Status = uint8(data[1])
		v1 := uint8(data[2])
		v2 := uint8(data[3])
		v3 := uint8(data[4])
		v4 := uint8(data[5])
		m.Version = fmt.Sprintf("%d.%d.%d.%d", v1, v2, v3, v4)
	}
	return m
}

// TODO: packet format
// +--------+----------------+
// | [0]    | [1-2]          |
// +--------+----------------+
// | opcode |                |
// +--------+----------------+
// | 29     |                |
// +--------+----------------+
// 4b000100040adf2900002800037000f0006e35
//
// opcode:            0x4b
// status:            0x00
// version:           1.0.4.10 (0x0100040a)
// BT version:        223.41.0.0 (0xdf290000)
// H/W rev:           40 (0x28)
// other F/W version: 0.3.112 (0x000370)
// asic:              0x00f0

// 4b000100040adf2900004800037000f0007486
//
// opcode:            4b
// status:            00
// version:           0100040a
// BT version:        df290000
// H/W rev:           48
// other F/W version: 000370
// asic:              00f0007486
