package randstrings

import "io"

type Reader struct {
	RandomSource io.Reader
	Charset      []byte
}

const (
	NumbersCharset          = "0123456789"
	LowerHexCharset         = NumbersCharset + "abcdef"
	UpperHexCharset         = NumbersCharset + "ABCDEF"
	LowerAlphabetCharset    = "abcdefghijklmnopqrstuvwxyz"
	UpperAlphabetCharset    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	BasicLatinSymbolCharset = " !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
)

func (r *Reader) Read(b []byte) (n int, err error) {
	n, err = r.RandomSource.Read(b)
	if err != nil {
		return
	}

	for i := range b[:n] {
		b[i] = r.Charset[int(b[i])%len(r.Charset)]
	}
	return
}
