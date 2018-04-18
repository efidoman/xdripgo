package messages

import (
	"github.com/efidoman/xdripgo/packet"
	"log"
)

type TransmitterTimeRxMessage struct {
	Opcode           byte
	Status           uint8
	CurrentTime      uint32
	SessionStartTime uint32
}

func NewTransmitterTimeRxMessage(data []byte) TransmitterTimeRxMessage {
	var m TransmitterTimeRxMessage

	m.Opcode = 0x25

	if len(data) != 16 {
		log.Print("Cannot create new TransmitterTimeRxMessage - Length not 16 bytes. Length = ", len(data))
	} else {
		if data[0] != m.Opcode {
			log.Print("Cannot create TransmitterTimeRxMessage - Response is not correct. Opcode should be = ", m.Opcode, " but data response Opcode is ", data[0])
		}
		m.Status = uint8(data[1])
		m.CurrentTime = packet.UnmarshalUint32(data[2:6])
		m.SessionStartTime = packet.UnmarshalUint32(data[6:10])
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
