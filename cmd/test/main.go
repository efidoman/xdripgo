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
}
