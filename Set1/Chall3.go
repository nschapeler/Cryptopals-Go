package Set1

type decryptionResultPretty struct {
	key     string
	payload string
}

type decryptionResult struct {
	key     []byte
	payload []byte
	score   uint32
}

// Score is the count of normal letters and spaces
func score(phrase []byte) uint32 {
	var score uint32 = 0

	for i := 0; i < len(phrase); i++ {
		if isAlphaOrSpace(phrase[i]) {
			score++
		}

	}

	return score
}

// Highest Score is best
func breakSingleByteXOR(cipherText []byte) decryptionResult {
	var payload []byte
	// Initialize as highest possible value
	var highestScore uint32 = 0
	var key byte

	// Try all possible byte values as key
	for i := 0; i < 256; i++ {
		xored := xorByteWise(cipherText, byte(i))

		if score(xored) >= highestScore {
			payload = xored
			key = byte(i)
			highestScore = score(xored)
		}
	}

	return decryptionResult{key: []byte{key}, payload: payload, score: highestScore}

}

func decrypt(cipherText string) decryptionResultPretty {
	res := breakSingleByteXOR(hexToByteArray(cipherText))
	return decryptionResultPretty{key: string(res.key), payload: string(res.payload)}
}

func isAlphaOrSpace(c byte) bool {
	// Small letter OR Capital letter OR Space
	return (c > 96 && c < 123) || (c < 91 && c > 64) || c == 32
}
