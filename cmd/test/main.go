// test

package main

import (
	"fmt"
	"github.com/efidoman/xdripgo/messages"
	"log"
)

func main() {
// Test Messsages

        data := make([]byte, 17)
        data[0] = 0x03

	a := messages.NewAuthChallengeRxMessage(data)
	log.Print("AuthChallengeRxMessage Opcode = ", a.Opcode)
//	log.Print("AuthChallengeRxMessage Data = ", a.Data)

	b := messages.NewAuthChallengeTxMessage(data)
	log.Print("AuthChallengeTxMessage Opcode = ", b.Opcode)
	log.Print("AuthChallengeTxMessage Data = ", b.Data)

	c := messages.NewAuthRequestTxMessage()
	log.Print("AuthRequestTxMessage Opcode = ", c.Opcode)
	log.Print("AuthRequestTxMessage Data = ", c.Data)

	d := messages.NewBackfillTxMessage(0, 64)
	log.Print("BackfillTxMessage Opcode = ", d.Opcode)
	log.Print("BackfillTxMessage Data = ", d.Data)

	e := messages.NewBondRequestTxMessage()
	log.Print("BondRequestTxMessage Opcode = ", e.Opcode)
	log.Print("BondRequestTxMessage Data = ", e.Data)

        data[0] = 0x08
	f := messages.NewBondRequestRxMessage(data)
	log.Print("BondRequestRxMessage Opcode = ", f.Opcode)
	log.Print("BondRequestRxMessage Succeeded = ", f.Succeeded)

	g := messages.NewCalibrateGlucoseTxMessage(100, 0) // 1970
	log.Print("CalibrateGlucoseTxMessage Opcode = ", g.Opcode)
	log.Print("CalibrateGlucoseTxMessage Data = ", g.Data)

        d1 := make([]byte, 5)
        d1[0] = 0x35
	h := messages.NewCalibrateGlucoseRxMessage(d1)
	log.Print("CalibrateGlucoseRxMessage Opcode (should be 35) = ", h.Opcode)
	log.Print("CalibrateGlucoseRxMessage Data = ", h.Data)

	i := messages.NewCalibrationDataTxMessage()
	log.Print("CalibrationDataTxMessage Opcode = ", i.Opcode)
	log.Print("CalibrationDataTxMessage Data (should be 32) = ", i.Data)

        data[0] = 0x42
	j := messages.NewResetTxMessage()
	fmt.Printf("ResetTxMessage Opcode (should be 42) = %x\n", j.Opcode)
	fmt.Printf("ResetTxMessage Data = %x\n", j.Data)

}
