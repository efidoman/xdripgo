package messages

type KeepAliveTxMessage struct {
	Opcode byte
	Data   []byte
}

func NewKeepAliveTxMessage(timestamp uint8) KeepAliveTxMessage {
	var m KeepAliveTxMessage

	m.Opcode = 0x06

	d := make([]byte, 2)
	d[0] = m.Opcode
	d[1] = timestamp

	m.Data = d

	return m
}
