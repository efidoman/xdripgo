package messages

import (
	"github.com/efidoman/xdripgo/packet"
	"log"
)

type SessionStopRxMessage struct {
	Opcode            byte
	Status            uint8
	Received          uint8
	RequestedStopTime uint32
	SessionStopTime   uint32
	TransmitterTime   uint32
}

func NewSessionStopRxMessage(data []byte) SessionStopRxMessage {
	var m SessionStopRxMessage

	m.Opcode = 0x29

	if len(data) != 17 {
		log.Print("Cannot create new SessionStopRxMessage - Length not 16 bytes. Length = ", len(data))
	} else {
		if data[0] != m.Opcode {
			log.Print("Cannot create SessionStopRxMessage - Response is not correct. Opcode should be = ", m.Opcode, " but data response Opcode is ", data[0])
		}
		m.Status = uint8(data[1])
		m.Received = uint8(data[2])
		m.RequestedStopTime = packet.UnmarshalUint32(data[3:7])
		m.SessionStopTime = packet.UnmarshalUint32(data[7:11])
		m.TransmitterTime = packet.UnmarshalUint32(data[11:15])
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
