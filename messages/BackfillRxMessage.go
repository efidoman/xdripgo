package messages

import (
//	"encoding/binary"
//      "github.com/efidoman/xdripgo/mathutils"
)

type BackfillRxMessage struct {
	Opcode         byte
	unknown1       []byte
	unknown2       []byte
	timestampStart uint32
	timestampEnd   uint32
}

func NewBackfillRxMessage(data []byte) BackfillRxMessage {
	var m BackfillRxMessage

	m.Opcode = 0x51

	return m
}

//function BackfillRxMessage(data) {
//  if ((data.length !== 20) || (data[0] !== opcode) || !crc.crcValid(data)) {
//    throw new Error('cannot create new BackfillRxMessage');
//  }
//  this.status = data.readUInt8(1);
//  this.unknown1 = data.readUInt8(2); // seen 1 or 2 (could mean data was returned?)
//  this.unknown2 = data.readUInt8(3); // seen 0, 1 or 2 (don't know)
//  this.timestampStart = data.readUInt32LE(4);
//  this.timestampEnd = data.readUInt32LE(8);
//}
