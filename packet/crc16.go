package packet

// go:generate crcgen -size 16 -poly 0x1021

// Compute CRC-16 using CCITT polynomial.
func Crc16(msg []byte) uint16 {
	res := uint16(0)
	for _, b := range msg {
		res = res<<8 ^ crc16Table[byte(res>>8)^b]
	}
	return res
}

func AppendCrc16(data []byte) []byte {
	n := len(data)
	crc := MarshalUint16(Crc16(data))
	d := make([]byte, n+2)
	copy(d[0:n], data)
	copy(d[n:n+2], crc)
	return d
}
