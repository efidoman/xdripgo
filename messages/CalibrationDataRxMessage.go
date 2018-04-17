package messages

import (
	//	"encoding/binary"
	"github.com/efidoman/xdripgo/mathutils"
	"log"
)

type CalibrationDataRxMessage struct {
	Opcode    byte
	Data      []byte
	Glucose   uint16
	Timestamp uint32
}

func NewCalibrationDataRxMessage(data []byte) CalibrationDataRxMessage {
	var m CalibrationDataRxMessage
	m.Opcode = 0x33

	if len(data) != 19 {
		log.Print("Cannot create new CalibrationDataRxMessage - Length not 2 bytes. Length = ", len(data))
	} else {
		if data[0] != m.Opcode {
			log.Print("Cannot create CalibrationDataRxMessage - Response is not correct. Opcode should be = ", m.Opcode, " but data response Opcode is ", data[0])
		}

		m.Glucose = mathutils.UnmarshalUint16(data[11:13])
		m.Timestamp = mathutils.UnmarshalUint32(data[13:17])
	}
	return m
}
