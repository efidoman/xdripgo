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

func SensorStateString(state byte) string {
	s := "--"
	switch state {
	case 0x00:
		s = "None"
	case 0x01:
		s = "Stopped"
	case 0x02:
		s = "Warmup"
	case 0x03:
		s = "Unused"
	case 0x04:
		s = "First calibration"
	case 0x05:
		s = "Second calibration"
	case 0x06:
		s = "OK"
	case 0x07:
		s = "Needs calibration"
	case 0x08:
		s = "Calibration Error 1"
	case 0x09:
		s = "Calibration Error 0"
	case 0x0a:
		s = "Calibration Linearity Fit Failure"
	case 0x0b:
		s = "Sensor Failed Due to Counts Aberration"
	case 0x0c:
		s = "Sensor Failed Due to Residual Aberration"
	case 0x0d:
		s = "Out of Calibration Due To Outlier"
	case 0x0e:
		s = "Outlier Calibration Request - Need a Calibration"
	case 0x0f:
		s = "Session Expired"
	case 0x10:
		s = "Session Failed Due To Unrecoverable Error"
	case 0x11:
		s = "Session Failed Due To Transmitter Error"
	case 0x12:
		s = "Temporary Session Failure"
	case 0x13:
		s = "Reserved"
	case 0x80:
		s = "Calibration State - Start"
	case 0x81:
		s = "Calibration State - Start Up"
	case 0x82:
		s = "Calibration State - First of Two Calibrations Needed"
	case 0x83:
		s = "Calibration State - High Wedge Display With First BG"
	case 0x84:
		s = "Unused Calibration State - Low Wedge Display With First BG"
	case 0x85:
		s = "Calibration State - Second of Two Calibrations Needed"
	case 0x86:
		s = "Calibration State - In Calibration Transmitter"
	case 0x87:
		s = "Calibration State - In Calibration Display"
	case 0x88:
		s = "Calibration State - High Wedge Transmitter"
	case 0x89:
		s = "Calibration State - Low Wedge Transmitter"
	case 0x8a:
		s = "Calibration State - Linearity Fit Transmitter"
	case 0x8b:
		s = "Calibration State - Out of Cal Due to Outlier Transmitter"
	case 0x8c:
		s = "Calibration State - High Wedge Display"
	case 0x8d:
		s = "Calibration State - Low Wedge Display"
	case 0x8e:
		s = "Calibration State - Linearity Fit Display"
	case 0x8f:
		s = "Calibration State - Session Not in Progress"
	default:
		s = "Unknown: 0x" + string(state)
	}
	return s
}
