package AuthChallengeRxMessage

import (
	"log"
)

type AuthChallengeRxMessage struct {
	Opcode    byte
	TokenHash []byte
	Challenge []byte
}

func New(data []byte) AuthChallengeRxMessage {
        var m AuthChallengeRxMessage
        m.Opcode = 0x03

	if data[0] != m.Opcode {
		log.Fatal("Cannot create AuthChallengeRxMessage - Response from AuthChallengeTxMessage not correct. Opcode should be = ", m.Opcode, " but data response Opcode is ", data[0])
	}
	if len(data) != 17  {
		log.Fatal("Cannot create new AuthChallengeRxMessage - Length not 17 bytes. Length = ", len(data))
	}
        m.TokenHash = data[0:0]
        m.Challenge = data[9:]
        return m
}

