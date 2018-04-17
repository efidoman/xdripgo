package messages

type DisconnectTxMessage struct {
	Opcode byte
	Data   []byte
}

func NewDisconnectTxMessage() DisconnectTxMessage {
	var m DisconnectTxMessage

	m.Opcode = 0x09

	d := make([]byte, 1)
	d[0] = m.Opcode

	m.Data = d

	return m
}

// packet format for DisconnectTxMessage
// +---------------+
// | [0]           |
// +---------------+
// | opcode        |
// +---------------+
// | 09            |
// +---------------+
