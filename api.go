package xdripgo

func TransmitterStatusString(status byte) string {
	s := "--"
	switch status {
	case 0x00:
		s = "OK"
	case 0x81:
		s = "Low battery"
	case 0x83:
		s = "Expired"
	default:
		s = "Unknown: 0x" + string(status)
	}
	return s
}
