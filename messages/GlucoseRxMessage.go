package messages

import (
	"github.com/efidoman/xdripgo/packet"
	"log"
)

type GlucoseRxMessage struct {
	Opcode               byte
	Glucose              uint16
	GlucoseBytes         uint16
	Timestamp            uint32
	State                uint8
	Status               uint8
	Sequence             uint32
	Trend                uint8
	GlucoseIsDisplayOnly uint8
}

func NewGlucoseRxMessage(data []byte) GlucoseRxMessage {
	var m GlucoseRxMessage
	m.Opcode = 0x31

	if len(data) != 16 {
		log.Print("Cannot create new GlucoseRxMessage - Length not 16 bytes. Length = ", len(data))
	} else {
		if data[0] != m.Opcode {
			log.Print("Cannot create CalibrationDataRxMessage - Response is not correct. Opcode should be = ", m.Opcode, " but data response Opcode is ", data[0])
		}

		m.Status = uint8(data[1])
		m.Sequence = packet.UnmarshalUint32(data[2:6])
		m.Timestamp = packet.UnmarshalUint32(data[6:10])
		m.GlucoseBytes = packet.UnmarshalUint16(data[10:12])
		m.Glucose = uint16(m.GlucoseBytes)
		m.State = uint8(data[12])
		m.Trend = uint8(data[13])
	}
	return m
}

// packet format - example TODO - finish documenting
// +--------+--------+-------------+-------------+---------+
// | [0]    | [1]    | [2-5]       | [6-9]       | [10-11] |
// +--------+--------+-------------+-------------+---------+
// | opcode | Status | Sequence    | Timestamp   | Glucose |
// +--------+--------+-------------+-------------+---------+
// | 31     | 00 9d  | 04 00 00 c4 | 17 08 00 05 | 00 05   |
// +--------+--------+-------------+-------------+---------+
