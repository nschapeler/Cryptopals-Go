package Set1

import "errors"

func xorBytes(a []byte, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, errors.New("lengths need to be the same")
	}

	res := make([]byte, len(a))

	for i := 0; i < len(res); i++ {
		res[i] = a[i] ^ b[i]
	}

	return res, nil
}

func xorByteWise(a []byte, b byte) []byte {
	res := make([]byte, len(a))

	for i := 0; i < len(res); i++ {
		res[i] = a[i] ^ b
	}

	return res
}
