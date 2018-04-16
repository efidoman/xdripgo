package xdripgo

// messages.go

import (
	"log"
)

// AuthChallengeRxMessage
type AuthChallengeRxMessage struct {
	opcode    byte
	tokenHash []byte
	challenge []byte
}

func (this *AuthChallengeRxMessage) NewAuthChallengeRxMessage(data []byte) {
	this.opcode = 0x03
	if len(data) != 17 || data[0] != this.opcode {
		log.Fatal("cannot create new AuthChallengeRxMessage")
	}
	this.tokenHash = data[0:9]
	this.challenge = data[9:]
}
