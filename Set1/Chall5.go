package Set1

func repeatingXOR(payload []byte, key []byte) []byte {
	for i := 0; i < len(payload); i++ {
		payload[i] = payload[i] ^ (key[i%len(key)])
	}

	return payload

}
