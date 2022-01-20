package Set1

import "testing"

// Test for the provided sample by Cryptopals
func TestByteWiseXor(t *testing.T) {
	hex1 := "1c0111001f010100061a024b53535009181c"
	hex2 := "686974207468652062756c6c277320657965"
	want := "746865206b696420646f6e277420706c6179"

	xored, err := xorBytes(hexToByteArray(hex1), hexToByteArray(hex2))

	if err != nil || !(want == byteArrayToHex(xored)) {
		t.Fatalf(`Got %q with error %s, expected %#q`, byteArrayToHex(xored), err, want)
	}

}

func TestByteWiseXorDifferentLength(t *testing.T) {
	hex1 := "1c0111001f010100061a024b53535009"
	hex2 := "686974207468652062756c6c277320657965"

	xored, err := xorBytes(hexToByteArray(hex1), hexToByteArray(hex2))

	if err == nil || xored != nil {
		t.Fatalf(`Excepted an error to be returned`)
	}

}
