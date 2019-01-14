package echoTransformer

import "golang.org/x/text/transform"

type echoDecoder struct{}
type echoEncoder struct{}

var EchoDecoder echoDecoder = echoDecoder{}
var EchoEncoder echoEncoder = echoEncoder{}

// Decoder
func (e echoDecoder) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	return Transform(dst, src, atEOF)
}

func (e echoDecoder) Reset() {}

// Encoder
func (e echoEncoder) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	return Transform(dst, src, atEOF)
}

func (e echoEncoder) Reset() {}

//
func Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	nDst = 0
	nSrc = 0
	for {
		if len(src) == 0 {
			break
		}
		if nDst+1 > len(dst) {
			return nDst, nSrc, transform.ErrShortDst
		}
		dst[nDst] = src[0]
		nDst += 1
		nSrc += 1
		src = src[1:]
	}
	return nDst, nSrc, nil
}
