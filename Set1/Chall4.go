package Set1

func getHighestScoredCipherText(hexs []string) decryptionResultPretty {

	payload := ""
	key := ""
	lowestScore := uint32(0)

	for _, element := range hexs {
		res := decrypt(element)

		if score([]byte(res.payload)) > lowestScore {
			lowestScore = score([]byte(res.payload))
			payload = res.payload
			key = res.key
		}
	}

	return decryptionResultPretty{payload: payload, key: key}
}
