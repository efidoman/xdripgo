package messages

import (
	"fmt"
	"github.com/satori/go.uuid"
)

type AuthRequestTxMessage struct {
	Opcode         byte
	EndByte        byte
	Data           []byte
	SingleUseToken []byte
}

func NewAuthRequestTxMessage() AuthRequestTxMessage {
	var m AuthRequestTxMessage

	u1 := uuid.Must(uuid.NewV4())
	fmt.Printf("UUIDv4: %x\n", u1)

	m.Opcode = 0x01
	m.EndByte = 0x2

	len := 10 
	d := make([]byte, len)
	d[0] = m.Opcode
	d[len-1] = m.EndByte
	copy(d[1:len-1], u1.Bytes())
	m.Data = d

	return m
}
