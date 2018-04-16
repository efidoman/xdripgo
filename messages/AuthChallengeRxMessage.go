package messages

import (
	"log"
)

type AuthChallengeRxMessage struct {
	Opcode    byte
	TokenHash []byte
	Challenge []byte
}

func NewAuthChallengeRxMessage(data []byte) AuthChallengeRxMessage {
	var m AuthChallengeRxMessage
	m.Opcode = 0x03

	if data[0] != m.Opcode {
		log.Print("Cannot create AuthChallengeRxMessage - Response from AuthChallengeTxMessage not correct. Opcode should be = ", m.Opcode, " but data response Opcode is ", data[0])
	}
	if len(data) != 17 {
		log.Print("Cannot create new AuthChallengeRxMessage - Length not 17 bytes. Length = ", len(data))
	}
	m.TokenHash = data[0:9]
	m.Challenge = data[9:]
	return m
}
