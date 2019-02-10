package cipher

func PaddingOracleBreak(
	bSize int, data, IV []byte,
	checkPadding func(data, IV []byte) bool) ([]byte, error) {

	ans := make([]byte, len(data))

	// append IV to data, treat it as the first block
	fData := append(IV, data...)

	// Iterate over all blocks, apart from IV
	numBlocks := len(fData) / bSize

	for i := numBlocks - 1; i > 0; i-- {

		curBlockStart := i * bSize
		prevBlockStart := (i - 1) * bSize

		for j := bSize - 1; j >= 0; j-- {
			tmp := make([]byte, curBlockStart+bSize)
			copy(tmp, fData)

			for k := j + 1; k < bSize; k++ {
				tmp[prevBlockStart+k] = fData[prevBlockStart+k] ^ ans[curBlockStart+k-bSize] ^ byte(bSize-j)
			}

			for k := 0; k <= 0xff; k++ {
				if i == numBlocks-1 && j == bSize-1 && k == 0 {
					continue
				}
				tmp2 := make([]byte, curBlockStart+bSize)
				copy(tmp2, tmp)

				tmp2[prevBlockStart+j] = fData[prevBlockStart+j] ^ byte(k)

				if checkPadding(tmp2[bSize:], tmp2[:bSize]) {
					ans[prevBlockStart+j] = byte(k) ^ byte(bSize-j)
					break
				}

			}

		}
	}
	return ans, nil
}
