package Set1

import (
	"errors"
)

func kasiski(cipherText []byte, minKeySize uint32, maxKeySize uint32) (decryptionResult, error) {

	keySizes, err := getKeySizeDistances(cipherText, minKeySize, maxKeySize)

	if err != nil {
		return decryptionResult{}, err
	}

	highestScore := uint32(0)
	keyGlobal := []byte{}
	payloadGlobal := []byte{}

	// How many keysizes we want to consider
	count := 3
	bottom := getBottomXValues(keySizes, uint32(count))

	for i := 0; i < count; i++ {

		// For each element key is the key size, value is the normalized Hemming Distance here
		keySize := bottom[i].key

		// Generate equal sized blocks
		blocks := getBlocks(cipherText, keySize)

		// Multi-byte key for the current keysize
		key := make([]byte, keySize)

		for i := range blocks {
			// One byte key
			keyInner := breakSingleByteXOR(blocks[i]).key
			// We know key of singly byte xor is one byte, so we can safely assume its an array of length 1
			key[i] = keyInner[0]
		}

		// Applying the key again is same as decrypting
		payload := repeatingXOR(cipherText, key)

		if score(payload) > highestScore {
			keyGlobal = key
			highestScore = score(payload)
			payloadGlobal = payload
		}

	}

	return decryptionResult{
		key:     keyGlobal,
		payload: payloadGlobal,
	}, nil
}

// Put all bytes into an array fitting the index with which byte of the key they would be xored e.g.
// e.g. if keysize = 3, byte 0 of the key would be xored with byte 0, 3, 6, 9 etc. of the ciphertext.
// This means we have keySize different "blocks", each of which is at most (len(ciphertext)/keySize) +1  large.
func getBlocks(cipherText []byte, keySize uint32) [][]byte {
	// We pick blockSize as (len(ciphertext)/keySize) since we use golang's append function which at worst will have to resize keysize-1
	// slices once but avoids us having to handle the special overflow case
	blockSize := (uint32(len(cipherText)) / keySize)

	// Initialize 2d slice
	blocks := make([][]byte, keySize)
	for i := range blocks {
		blocks[i] = make([]byte, blockSize)
	}

	blockIndex := 0
	// Organize the bytes into the blocks like described above
	for i := range cipherText {
		if uint32(blockIndex) >= blockSize {
			blocks[uint32(i)%keySize] = append(blocks[uint32(i)%keySize], cipherText[i])
		}
		blocks[uint32(i)%keySize][blockIndex] = cipherText[i]

		// Increment the index within block if we're back at block 0
		if uint32(i)%keySize == 0 && i > 0 {
			blockIndex++
		}

	}

	return blocks

}

func getKeySizeDistances(cipherText []byte, minKeySize uint32, maxKeySize uint32) (map[uint32]float32, error) {
	distances := make(map[uint32]float32)

	for k := minKeySize; k <= maxKeySize; k++ {
		dist, err := getNormalizedDistance(cipherText, k)

		if err != nil {
			return nil, err
		}

		distances[k] = dist

	}

	return distances, nil

}

func getNormalizedDistance(cipherText []byte, keySize uint32) (float32, error) {
	totalDistance := uint32(0)

	sliceSize := keySize * 2

	blockCount := uint32(len(cipherText))/sliceSize - 1
	distanceCount := 0

	// Iterate over whole cypertext instead of first 4 chunks just to be extra sure
	for i := 0; i < int(blockCount); i++ {
		for j := i; j < int(blockCount); j++ {
			// Compare difference between first k bytes and second k bytes
			b1 := cipherText[(i * int(sliceSize)) : (i+1)*int(sliceSize)]
			b2 := cipherText[j*int(sliceSize) : (j+1)*int(sliceSize)]
			distance, err := getHammingDistance(b1, b2)

			if err != nil {
				return 0, err
			}

			totalDistance += distance
			distanceCount++
		}
	}

	// blockCount distances computed
	avgDistance := float32(totalDistance) / float32(distanceCount)

	// Normalize our avg distance by dividing by keysize
	return avgDistance / float32(sliceSize), nil

}

func getHammingDistance(a []byte, b []byte) (uint32, error) {
	if len(a) != len(b) {
		return 0, errors.New("Byte arrays have to have same length")
	}

	distance := uint32(0)

	for i := 0; i < len(a); i++ {
		distance += getHammingDistanceByteWise(a[i], b[i])
	}

	return distance, nil

}

func getHammingDistanceByteWise(a byte, b byte) uint32 {
	// Get differing bits
	c := a ^ b

	count := uint32(0)

	// Brian Kernighan Algorithm to count 1 bits
	for c != 0 {
		c = c & (c - 1)
		count++
	}

	return count

}
