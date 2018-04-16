package messages

import (
	"log"
)

type CalibrateGlucoseRxMessage struct {
	Opcode byte
	Data   []byte
}

func NewCalibrateGlucoseRxMessage(data []byte) CalibrateGlucoseRxMessage {
	var m CalibrateGlucoseRxMessage

	m.Opcode = 0x35

	if len(data) != 5 {
		log.Print("Cannot create new CalibrateGlucoseRxMessage - Length not 5 bytes. Length = ", len(data))
	}
	if data[0] != m.Opcode {
		log.Print("Cannot create CalibrateGlucoseRxMessage - Response from CalibrateGlucoseRxMessage not correct. Opcode should be = ", m.Opcode, " but data response Opcode is ", data[0])
	}
// TODO: work out what the payload is
// presumably calibration succeeded / rejected
	m.Data = data

	return m
}

// example: 350000552e
// this one above is the usual one (cal successful)
// got this one when I think the cal was rejected
//          3500085daf

// perhaps its something like
// opcode:    35
// something: 00
// status:    00 OK
//            06 second calibration
//            08 rejected - enter another one
//            0b no can do - sensor stopped

// crc:       5daf

// if we try to calibrate when the sensor is stopped:
// 35000b3e9f
