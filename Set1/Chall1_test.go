package Set1

import (
	"testing"
)

// Test for the provided sample by Cryptopals
func TestCorrectByteCount(t *testing.T) {
	hex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	want := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	res := hexToBase64(hex)

	if !(want == res) {
		t.Fatalf(`Got %q, expected %#q`, res, want)
	}

}

func TestTwoBytesTooMuch(t *testing.T) {
	hex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f"
	want := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb28="
	res := hexToBase64(hex)
	if !(want == res) {
		t.Fatalf(`Got %q, expected %#q`, res, want)
	}
}

func TestOneByteTooMuch(t *testing.T) {
	hex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f"
	want := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hybw=="
	res := hexToBase64(hex)
	if !(want == res) {
		t.Fatalf(`Got %q, expected %#q`, res, want)
	}
}

func TestHexConversion(t *testing.T) {
	want := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f"
	res := byteArrayToHex(hexToByteArray(want))
	if !(want == res) {
		t.Fatalf(`Got %q, expected %#q`, res, want)
	}

}
