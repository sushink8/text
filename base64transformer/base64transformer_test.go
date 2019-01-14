package base64transformer

import (
	"strings"
	"testing"
)

// https://golang.org/src/encoding/base64/base64_test.go
type testpair []string

var pairs = []testpair{
	// RFC 3548 examples
	{"\x14\xfb\x9c\x03\xd9\x7e", "FPucA9l+"},
	{"\x14\xfb\x9c\x03\xd9", "FPucA9k="},
	{"\x14\xfb\x9c\x03", "FPucAw=="},

	// RFC 4648 examples
	{"", ""},
	{"f", "Zg=="},
	{"fo", "Zm8="},
	{"foo", "Zm9v"},
	{"foob", "Zm9vYg=="},
	{"fooba", "Zm9vYmE="},
	{"foobar", "Zm9vYmFy"},

	// Wikipedia examples
	{"sure.", "c3VyZS4="},
	{"sure", "c3VyZQ=="},
	{"sur", "c3Vy"},
	{"su", "c3U="},
	{"leasure.", "bGVhc3VyZS4="},
	{"easure.", "ZWFzdXJlLg=="},
	{"asure.", "YXN1cmUu"},
	{"sure.", "c3VyZS4="},
}

func TestEncode(t *testing.T) {
	for _, p := range pairs {
		src := []byte(p[1])
		dst := make([]byte, 100)
		Base64Decoder.Transform(dst, src, true)
		if p[0] == strings.Trim(string(dst), "\x00") {
			// OK
		} else {
			t.Errorf("Encode = %q, want %q", string(dst), string(p[0]))
		}
	}
}
