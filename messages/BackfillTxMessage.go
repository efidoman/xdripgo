package messages

import (
	"encoding/binary"
	"github.com/efidoman/xdripgo/packet"
)

type BackfillTxMessage struct {
	Opcode byte
	Data   []byte
}

func NewBackfillTxMessage(timestampStart uint32, timestampEnd uint32) BackfillTxMessage {
	var m BackfillTxMessage

	m.Opcode = 0x50

	d := make([]byte, 9)
	d[0] = m.Opcode

	bs := make([]byte, 4)
	be := make([]byte, 4)

	binary.LittleEndian.PutUint32(bs, timestampStart)
	binary.LittleEndian.PutUint32(be, timestampEnd)

	copy(d[1:5], bs)
	copy(d[5:9], be)

	d = packet.AppendCrc16(d)

	/*
		crc := packet.MarshalUint16(packet.Crc16(d))

		e := make([]byte, 11)

		copy(e[0:9], d)
		copy(e[9:11], crc)
		m.Data = e
	*/
	m.Data = d

	return m
}

// NOTE: this is a total guess
// +--------+--------------------------+------------------------+--------+
// | [0]    | [1-4]                    | [5-8]                  | [9-10] |
// +--------+--------------------------+------------------------+--------+
// | opcode | dexcomStartTimeInSeconds | dexcomEndTimeInSeconds | CRC    |
// +--------+--------------------------+------------------------+--------+
// | 50     | 9e 32 66 00              | ce 5c 66 00            | 87 77  |
// +--------+--------------------------+------------------------+--------+
