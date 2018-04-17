package messages

import (
	"github.com/efidoman/xdripgo/mathutils"
)

type KeepAliveTxMessage struct {
	Opcode byte
	Data   []byte
}

func NewKeepAliveTxMessage(timestamp uint32) KeepAliveTxMessage {
	var m KeepAliveTxMessage

	m.Opcode = 0x06

	d := make([]byte, 5)
	d[0] = m.Opcode

	tbuffer := mathutils.MarshalUint32(timestamp)

	copy(d[1:5], tbuffer)
	m.Data = d

	return m
}
