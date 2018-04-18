package messages

type BondRequestTxMessage struct {
	Opcode byte
	Data   []byte
}

func NewBondRequestTxMessage() BondRequestTxMessage {
	var m BondRequestTxMessage
	m.Opcode = 0x07
	d := make([]byte, 1)
	d[0] = m.Opcode
	m.Data = d

	return m
}

// packet format
// +--------+
// | [0]    |
// +--------+
// | opcode |
// +--------+
// | 07     |
// +--------+
