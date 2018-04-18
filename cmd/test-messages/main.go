// test

package main

import (
	"encoding/hex"
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

	n := messages.NewSensorTxMessage()
	fmt.Printf("%T Opcode (0x2e) = %x\n", n, n.Opcode)
	fmt.Printf("   Data = %x\n", n.Data)

	o := messages.NewGlucoseTxMessage()
	fmt.Printf("%T Opcode (0x30) = %x\n", o, o.Opcode)
	fmt.Printf("   Data = %x\n", o.Data)

	d3, _ := hex.DecodeString("31009d040000c41708000500057f6f12")
	p := messages.NewGlucoseRxMessage(d3)
	fmt.Printf("%T Opcode (0x31) = %x\n", p, p.Opcode)
	fmt.Printf("   Glucose = %d\n", p.Glucose)
	fmt.Printf("   GlucoseBytes = %d\n", p.GlucoseBytes)
	fmt.Printf("   Timestamp = %d\n", p.Timestamp)
	fmt.Printf("   State = %d\n", p.State)
	fmt.Printf("   Status = %d\n", p.Status)
	fmt.Printf("   Sequence = %d\n", p.Sequence)
	fmt.Printf("   Trend = %d\n", p.Trend)
	fmt.Printf("   GlucoseIsDisplayOnly = %d\n", p.GlucoseIsDisplayOnly)

	d4, _ := hex.DecodeString("2f00cc1e080020310200e03602007aaf")
	q := messages.NewSensorRxMessage(d4)
	fmt.Printf("%T Opcode (0x2f) = %x\n", q, q.Opcode)
	fmt.Printf("   Timestamp = %d\n", q.Timestamp)
	fmt.Printf("   Unfiltered = %d\n", q.Unfiltered)
	fmt.Printf("   Filtered = %d\n", q.Filtered)

	r := messages.NewSessionStartTxMessage(uint32(secs))
	fmt.Printf("%T Opcode (0x26) = %x\n", r, r.Opcode)
	fmt.Printf("   Data = %x\n", r.Data)

	d5, _ := hex.DecodeString("2700cc1e080020310200e03602007aaf00")
	s := messages.NewSessionStartRxMessage(d5)
	fmt.Printf("%T Opcode (0x27) = %x\n", s, s.Opcode)
	fmt.Printf("   Status = %d\n", s.Status)
	fmt.Printf("   Received = %d\n", s.Received)
	fmt.Printf("   RequestedStartTime = %d\n", s.RequestedStartTime)
	fmt.Printf("   SessionStartTime = %d\n", s.SessionStartTime)
	fmt.Printf("   TransmitterTime = %d\n", s.TransmitterTime)

	t := messages.NewSessionStopTxMessage(uint32(secs))
	fmt.Printf("%T Opcode (0x28) = %x\n", t, t.Opcode)
	fmt.Printf("   Data = %x\n", t.Data)

	d6, _ := hex.DecodeString("2900cc1e080020310200e03602007aaf00")
	u := messages.NewSessionStopRxMessage(d6)
	fmt.Printf("%T Opcode (0x29) = %x\n", u, u.Opcode)
	fmt.Printf("   Status = %d\n", u.Status)
	fmt.Printf("   Received = %d\n", u.Received)
	fmt.Printf("   RequestedStopTime = %d\n", u.RequestedStopTime)
	fmt.Printf("   SessionStopTime = %d\n", u.SessionStopTime)
	fmt.Printf("   TransmitterTime = %d\n", u.TransmitterTime)

	v := messages.NewTransmitterTimeTxMessage()
	fmt.Printf("%T Opcode (0x24) = %x\n", v, v.Opcode)
	fmt.Printf("   Data = %x\n", v.Data)

	d7, _ := hex.DecodeString("2500cc1e080020310200e03602007aaf")
	x := messages.NewTransmitterTimeRxMessage(d7)
	fmt.Printf("%T Opcode (0x25) = %x\n", x, x.Opcode)
	fmt.Printf("   Status = %d\n", x.Status)
	fmt.Printf("   currentTime = %d\n", x.CurrentTime)
	fmt.Printf("   SessionStartTime = %d\n", x.SessionStartTime)

	y := messages.NewVersionRequestTxMessage()
	fmt.Printf("%T Opcode (0x4a) = %x\n", y, y.Opcode)
	fmt.Printf("   Data = %x\n", y.Data)

	d8, _ := hex.DecodeString("4b000100040adf2900002800037000f0006e35")
	z := messages.NewVersionRequestRxMessage(d8)
	fmt.Printf("%T Opcode (0x25) = %x\n", z, z.Opcode)
	fmt.Printf("   Status = %d\n", z.Status)
	fmt.Printf("   Version = %s\n", z.Version)
}
