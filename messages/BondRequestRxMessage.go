package messages

import (
	"log"
)

type BondRequestRxMessage struct {
	Opcode    byte
	Succeeded byte
}

func NewBondRequestRxMessage(data []byte) BondRequestRxMessage {
	var m BondRequestRxMessage
	m.Opcode = 0x08

	if len(data) != 2 {
		log.Print("Cannot create new BondRequestRxMessage - Length not 2 bytes. Length = ", len(data))
	} else {
		if data[0] != m.Opcode {
			log.Print("Cannot create BondRequestRxMessage - Response from BondRequestRxMessage not correct. Opcode should be = ", m.Opcode, " but data response Opcode is ", data[0])
		}
		m.Succeeded = data[1]
	}
	return m
}
