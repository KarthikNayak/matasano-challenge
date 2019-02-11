package helpers

func IntToLittleEndianBytes(val uint64) []byte {
	b := make([]byte, 8)
	b[0] = byte((val & 0x00000000000000ff) >> 0)
	b[1] = byte((val & 0x000000000000ff00) >> 8)
	b[2] = byte((val & 0x0000000000ff0000) >> 16)
	b[3] = byte((val & 0x00000000ff000000) >> 24)
	b[4] = byte((val & 0x000000ff00000000) >> 32)
	b[5] = byte((val & 0x0000ff0000000000) >> 40)
	b[6] = byte((val & 0x00ff000000000000) >> 48)
	b[7] = byte((val & 0xff00000000000000) >> 56)

	return b
}

func LittleEndianToInt(b []byte) uint64 {
	if len(b) != 8 {
		return 0
	}

	val := uint64(b[0])
	val = val | (uint64(b[1]) << 8)
	val = val | (uint64(b[2]) << 16)
	val = val | (uint64(b[3]) << 24)
	val = val | (uint64(b[4]) << 32)
	val = val | (uint64(b[5]) << 40)
	val = val | (uint64(b[6]) << 48)
	val = val | (uint64(b[7]) << 56)

	return val
}
