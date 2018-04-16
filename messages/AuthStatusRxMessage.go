package messages

import (
	"log"
)

type AuthStatusRxMessage struct {
	Opcode        byte
	Authenticated byte
	Bonded        byte
}

func NewAuthStatusRxMessage(data []byte) AuthStatusRxMessage {
	var m AuthStatusRxMessage
	m.Opcode = 0x05

	if data[0] != m.Opcode {
		log.Print("Cannot create AuthStatusRxMessage - Response from AuthStatusRxMessage not correct. Opcode should be = ", m.Opcode, " but data response Opcode is ", data[0])
	}
	if len(data) != 3 {
		log.Print("Cannot create new AuthStatusRxMessage - Length not 3 bytes. Length = ", len(data))
	}
	m.Authenticated = data[1]
	m.Bonded = data[2]
	return m
}
