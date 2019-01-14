package base64transformer

import (
	"encoding/base64"
	"io/ioutil"
	"strings"

	"golang.org/x/text/transform"
)

type base64Decoder struct {
	Encoding *base64.Encoding
}

const base64EncodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func (b base64Decoder) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	nDst = 0
	nSrc = 0
	for {
		if len(src) == 0 {
			break
		}
		if len(src) < 4 {
			return nDst, nSrc, transform.ErrShortSrc
		}
		if nDst+1 > len(dst) {
			return nDst, nSrc, transform.ErrShortDst
		}
		if strings.Index(base64EncodeStd, string(src[0])) == -1 {
			dst[nDst] = src[0]
			nDst += 1
			nSrc += 1
			src = src[1:]
			continue
		}
		b := src[0:4]
		decode := base64.NewDecoder(base64.StdEncoding, strings.NewReader(string(b)))
		str, ioerr := ioutil.ReadAll(decode)
		if ioerr != nil {
			return nDst, nSrc, ioerr
		}
		copy(dst[nDst:], []byte(str))
		nDst += 3
		nSrc += 4
		src = src[4:]

	}
	return nDst, nSrc, nil
}

func (b base64Decoder) Reset() {}

var Base64Decoder base64Decoder

func init() {
	Base64Decoder = base64Decoder{
		Encoding: base64.StdEncoding,
	}
	// decode := base64.NewDecoder(base64.StdEncoding, strings.NewReader(encodedString))
}
