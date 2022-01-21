package Set1

import (
	"testing"
)

// Test for the provided sample by Cryptopals
func TestCorrectDecryption(t *testing.T) {
	cipherText := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	want := "Cooking MC's like a pound of bacon"

	res := decrypt(cipherText)

	if !(want == res.payload) {
		t.Fatalf(`Got %q, expected %#q`, res.payload, want)
	}

}
