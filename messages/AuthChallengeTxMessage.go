package messages

type AuthChallengeTxMessage struct {
	Opcode		byte
	Data		[]byte
}

func NewAuthChallengeTxMessage(ChallengeHash []byte) AuthChallengeTxMessage {
        var m AuthChallengeTxMessage
        m.Opcode = 0x04
        len := len(ChallengeHash)+1
        d := make([]byte, len)
        d[0] = m.Opcode
        copy(d[1:],ChallengeHash)
        m.Data = d 

        return m
}

