package packet

import (
	"math"
)

// Marshaling and Unmarshaling of ints and floats in little-endian order.

func MarshalUint16(n uint16) []byte {
	return []byte{byte(n & 0xFF), byte(n >> 8)}
}

// nolint
func MarshalInt16(n int16) []byte {
	return MarshalUint16(uint16(n))
}

func MarshalUint32(n uint32) []byte {
	return append(MarshalUint16(uint16(n&0xFFFF)), MarshalUint16(uint16(n>>16))...)
}

func MarshalInt32(n int32) []byte {
	return MarshalUint32(uint32(n))
}

func UnmarshalUint16(v []byte) uint16 {
	return uint16(v[0]) | uint16(v[1])<<8
}

// nolint
func UnmarshalInt16(v []byte) int16 {
	return int16(UnmarshalUint16(v))
}

func UnmarshalUint32(v []byte) uint32 {
	return uint32(UnmarshalUint16(v[0:2])) | uint32(UnmarshalUint16(v[2:4]))<<16
}

func UnmarshalInt32(v []byte) int32 {
	return int32(UnmarshalUint32(v))
}

func UnmarshalUint64(v []byte) uint64 {
	return uint64(UnmarshalUint32(v[0:4])) | uint64(UnmarshalUint32(v[4:8]))<<32
}

func UnmarshalFloat64(v []byte) float64 {
	return math.Float64frombits(UnmarshalUint64(v))
}
