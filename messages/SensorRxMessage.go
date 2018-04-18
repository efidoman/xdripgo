package messages

import (
	"github.com/efidoman/xdripgo/packet"
	"log"
)

type SensorRxMessage struct {
	Opcode     byte
	Status     uint8
	Timestamp  uint32
	Unfiltered uint32
	Filtered   uint32
}

func NewSensorRxMessage(data []byte) SensorRxMessage {
	var m SensorRxMessage
	m.Opcode = 0x2f

	if len(data) != 16 {
		log.Print("Cannot create new SensorRxMessage - Length not 16 bytes. Length = ", len(data))
	} else {
		if data[0] != m.Opcode {
			log.Print("Cannot create SensorRxMessage - Response is not correct. Opcode should be = ", m.Opcode, " but data response Opcode is ", data[0])
		}

		m.Status = uint8(data[1])
		m.Timestamp = packet.UnmarshalUint32(data[2:6])
		m.Unfiltered = packet.UnmarshalUint32(data[6:10])
		m.Filtered = packet.UnmarshalUint32(data[10:14])
	}
	return m
}

// 2f 00 bf 2f 07 00 60 67 02 00 c0 6d 02 00 f9 7b
// 2f 00 cc 1e 08 00 20 31 02 00 e0 36 02 00 7a af
// Timestamp is crap - do time.Now()

// packet format - example TODO - finish documenting
// TODO: use transmitterStartDate to calculate timestamp - has to happen
//   readDate = transmitterStartDate + glucoseMessage.timestamp * 1000
// +--------+--------+-------------+-------------+-------------+
// | [0]    | [1]    | [2-5]       | [6-9]       | [10-13]     |
// +--------+--------+-------------+-------------+-------------+
// | opcode | Status | Timstmp crap| Unfiltered  | Filtered    |
// +--------+--------+-------------+-------------+-------------+
// | 2f     | 00 cc  | 1e 08 00 20 | 31 02 00 e0 | 36 02 7a af |
// +--------+--------+-------------+-------------+-------------+
