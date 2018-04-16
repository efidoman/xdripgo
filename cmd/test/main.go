// test

package main

import (
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
}
