package Set1

type decryptionResultPretty struct {
	key     string
	payload string
}

type decryptionResult struct {
	key     byte
	payload []byte
}

// Each letters score is calculated by taking its frequency in the english language multiplied by 100
func getFrequency(letter byte) uint32 {
	letterToScore := map[byte]uint32{
		'a': 850,
		'b': 207,
		'c': 454,
		'd': 338,
		'e': 1116,
		'f': 181,
		'g': 247,
		'h': 300,
		'i': 754,
		'j': 20,
		'k': 110,
		'l': 549,
		'm': 301,
		'n': 665,
		'o': 716,
		'p': 317,
		'q': 20,
		'r': 758,
		's': 573,
		't': 695,
		'u': 363,
		'v': 101,
		'x': 29,
		'y': 178,
		'z': 27,
	}

	return letterToScore[letter]

}

// Make all our letters small
func normalizeLetter(letter byte) byte {
	// Letter is capital
	if letter >= 65 && letter <= 90 {
		return letter + 32
	}
	return letter
}

// Score is the sum of the differences in frequency of the letters in english and in our phrase
func score(phrase []byte) uint32 {

	frequencies := make(map[byte]uint32)
	var score uint32 = 0

	// Count the frequencies of our letters
	for i := 0; i < len(phrase); i++ {
		// First 32 ASCII symbols are extremely unlikely to appear in any phrase meant to be read by a human
		if phrase[i] < 32 {
			return ^uint32(0)
		}

		// This works because go initializes map values to their zero value, so 0 for uint32
		frequencies[normalizeLetter(phrase[i])] += 1
	}

	// Iterate over all small letters and compare frequencies
	for i := 97; i < 123; i++ {
		score += diff((frequencies[byte(i)]*100)/uint32(len(phrase)), getFrequency(byte(i)))
	}

	return score

}

// Lowest Score is best
func getLowestScoreDecryption(cipherText []byte) decryptionResult {
	var payload []byte
	// Initialize as highest possible value
	var lowestScore uint32 = ^uint32(0)
	var key byte

	// Try all possible byte values as key
	for i := 0; i < 256; i++ {
		xored := xorByteWise(cipherText, byte(i))

		if score(xored) < lowestScore {
			payload = xored
			key = byte(i)
			lowestScore = score(xored)
		}
	}

	return decryptionResult{key: key, payload: payload}

}

func decrypt(cipherText string) decryptionResultPretty {
	res := getLowestScoreDecryption(hexToByteArray(cipherText))
	return decryptionResultPretty{key: string(res.key), payload: string(res.payload)}
}

func diff(a, b uint32) uint32 {
	if a < b {
		return b - a
	}
	return a - b
}
