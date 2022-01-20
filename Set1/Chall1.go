package Set1

func hexToBase64(hexString string) string {
	return byteArrayToBase64(hexToByteArray(hexString))
}

func hexToByteArray(hexString string) []byte {
	bytes := make([]byte, len(hexString)/2)

	hexCharToByte := map[byte]byte{
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'a': 10,
		'b': 11,
		'c': 12,
		'd': 13,
		'e': 14,
		'f': 15,
	}

	// Since all chars in a hex string take just one byte, we iterate over our String by doing
	for i := 0; i < len(hexString); i += 2 {
		bytes[i/2] = hexCharToByte[hexString[i]]*16 + hexCharToByte[hexString[i+1]]
	}

	return bytes
}

func byteArrayToBase64(byteArray []byte) string {

	// 3 bytes = 4 base64 symbols
	res := make([]byte, (len(byteArray)*4)/3)
	resultIndex := 0

	for i := 0; i+2 < len(byteArray); i += 3 {
		// 24 bits long so we need space
		var val uint32
		val = 0
		// Concatenate the three bytes together
		val = uint32(byteArray[i])<<16 + uint32(byteArray[i+1])<<8 + uint32(byteArray[i+2])

		// Break it down into 4 base64 symbols
		for j := 0; j < 4; j++ {
			res[resultIndex+3-j] = byteToBase64Symbol(byte(val % 64))
			val /= 64
		}

		resultIndex += 4

	}

	// Bytes were exactly divisble by 3
	if len(byteArray)%3 == 0 {
		return string(res)
	}

	// Have leftover bytes? Need to shift to make the bitcount divisible by 6

	// Need 4 extra base64 chars for our leftover bytes, intialize with padding
	ending := make([]byte, 0)

	var finalVal uint32

	// One byte left over
	if len(byteArray)%3 == 1 {
		// Make our final value 12 bit long

		finalVal = uint32(byteArray[len(byteArray)-1]) << 4
		res = res[0 : len(res)-1]

		// Two bytes left over
	} else {
		// Make our final value 18 bit long
		finalVal = uint32(byteArray[len(byteArray)-2])<<10 + uint32(byteArray[len(byteArray)-1])<<2

		res = res[0 : len(res)-2]

	}

	// Parse our last value to base64
	for finalVal > 0 {
		ending = prependbyte(ending, byteToBase64Symbol(byte(finalVal%64)))
		finalVal /= 64
	}

	// Pad it to 4 chars using "=" symbol
	for len(ending) < 4 {
		ending = append(ending, '=')
	}

	return string(res) + string(ending)

}

func byteToBase64Symbol(b byte) byte {
	// Capital letters
	if b <= 25 {
		return b + 65
	}
	// Small letters
	if b <= 51 {
		return b + 71
	}
	// Numbers
	if b <= 61 {
		return b - 4
	}
	// "+"" symbol
	if b == 62 {
		return 42
	}
	// "/" symbol
	return 47
}

// Taken from https://stackoverflow.com/questions/53737435/how-to-prepend-int-to-slice
func prependbyte(x []byte, y byte) []byte {
	x = append(x, y)
	copy(x[1:], x)
	x[0] = y
	return x
}
