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
	m := AuthChallengeRxMessage.New(data)
	log.Print("AuthChallengeRxMessage Opcode = ", m.Opcode)

}
