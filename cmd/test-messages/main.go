// test

package main

import (
	"fmt"
	"github.com/efidoman/xdripgo/messages"
	"time"
)

func main() {
	// Test Messsages

	data := make([]byte, 17)

	data[0] = 0x03
	a := messages.NewAuthChallengeRxMessage(data)
	fmt.Printf("%T Opcode (0x03) = %x\n", a, a.Opcode)
	fmt.Printf("   TokenHash = %x\n", a.TokenHash)
	fmt.Printf("   Challenge = %x\n", a.Challenge)

	b := messages.NewAuthChallengeTxMessage(data)
	fmt.Printf("%T Opcode (0x04) = %x\n", b, b.Opcode)
	fmt.Printf("   Data = %x\n", b.Data)

	c := messages.NewAuthRequestTxMessage()
	fmt.Printf("%T Opcode (0x01) = %x\n", c, c.Opcode)
	fmt.Printf("   Data = %x\n", c.Data)

	d := messages.NewBackfillTxMessage(0, 64)
	fmt.Printf("%T Opcode (0x50) = %x\n", d, d.Opcode)
	fmt.Printf("   Data = %x\n", d.Data)

	e := messages.NewBondRequestTxMessage()
	fmt.Printf("%T -- Opcode (0x07) = %x\n", e, e.Opcode)
	fmt.Printf("   Data = %x\n", e.Data)

	d2 := make([]byte, 2)
	d2[0] = 0x08
	f := messages.NewBondRequestRxMessage(d2)
	fmt.Printf("%T Opcode (0x08) = %x\n", f, f.Opcode)
	fmt.Printf("   Succeeded = %x\n", f.Succeeded)

	g := messages.NewCalibrateGlucoseTxMessage(100, 0) // 1970
	fmt.Printf("%T Opcode (0x34) = %x\n", g, g.Opcode)
	fmt.Printf("   Data = %x\n", g.Data)

	d1 := make([]byte, 5)
	d1[0] = 0x35
	h := messages.NewCalibrateGlucoseRxMessage(d1)
	fmt.Printf("%T Opcode (0x35) = %x\n", h, h.Opcode)
	fmt.Printf("   Data = %x\n", h.Data)

	i := messages.NewCalibrationDataTxMessage()
	fmt.Printf("%T Opcode (0x32) = %x\n", i, i.Opcode)
	fmt.Printf("   Data = %x\n", i.Data)

	data[0] = 0x42
	j := messages.NewResetTxMessage()
	fmt.Printf("%T Opcode (0x42) = %x\n", j, j.Opcode)
	fmt.Printf("   Data = %x\n", j.Data)

	data[0] = 0x33
	k := messages.NewCalibrationDataRxMessage(data)
	fmt.Printf("%T Opcode (0x33) = %x\n", k, k.Opcode)
	fmt.Printf("   Glucose = %x\n", k.Glucose)
	fmt.Printf("   Timestamp = %x\n", k.Timestamp)

	l := messages.NewDisconnectTxMessage()
	fmt.Printf("%T Opcode (0x09) = %x\n", l, l.Opcode)
	fmt.Printf("   Data = %x\n", l.Data)

	now := time.Now()
	secs := now.Unix()
	m := messages.NewKeepAliveTxMessage(uint32(secs))
	fmt.Printf("%T Opcode (0x06) = %x\n", m, m.Opcode)
	fmt.Printf("   Data = %x\n", m.Data)
}
