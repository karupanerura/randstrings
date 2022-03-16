package randstrings_test

import (
	"fmt"
	"io"
	"math/rand"
	"testing"

	"github.com/karupanerura/randstrings"
)

func TestReader(t *testing.T) {
	t.Run("NumbersCharset", func(t *testing.T) {
		r := newRandomStringReader([]byte(randstrings.NumbersCharset))

		var b [64]byte
		_, err := r.Read(b[:])
		if err != nil {
			t.Fatal(t)
		}
		if s := string(b[:]); s != "1834072251584150099568582294451714352784737013521076636510508748" {
			t.Errorf("unexpected string: %s", s)
		}
	})
	t.Run("LowerHexCharset", func(t *testing.T) {
		r := newRandomStringReader([]byte(randstrings.LowerHexCharset))

		var b [64]byte
		_, err := r.Read(b[:])
		if err != nil {
			t.Fatal(t)
		}
		if s := string(b[:]); s != "14d2afc013f24b38ef9f625e8ad4adb5b8df87e01d9615f4189e25c11458856e" {
			t.Errorf("unexpected string: %s", s)
		}
	})
	t.Run("UpperHexCharset", func(t *testing.T) {
		r := newRandomStringReader([]byte(randstrings.UpperHexCharset))

		var b [64]byte
		_, err := r.Read(b[:])
		if err != nil {
			t.Fatal(t)
		}
		if s := string(b[:]); s != "14D2AFC013F24B38EF9F625E8AD4ADB5B8DF87E01D9615F4189E25C11458856E" {
			t.Errorf("unexpected string: %s", s)
		}
	})
	t.Run("LowerAlphabetCharset", func(t *testing.T) {
		r := newRandomStringReader([]byte(randstrings.LowerAlphabetCharset))

		var b [64]byte
		_, err := r.Read(b[:])
		if err != nil {
			t.Fatal(t)
		}
		if s := string(b[:]); s != "bstmqvskndvsenlsgbprmujeyqhkwtlnryntulcqvnfsddpefqhusdcrduhqinga" {
			t.Errorf("unexpected string: %s", s)
		}
	})
	t.Run("UpperAlphabetCharset", func(t *testing.T) {
		r := newRandomStringReader([]byte(randstrings.UpperAlphabetCharset))

		var b [64]byte
		_, err := r.Read(b[:])
		if err != nil {
			t.Fatal(t)
		}
		if s := string(b[:]); s != "BSTMQVSKNDVSENLSGBPRMUJEYQHKWTLNRYNTULCQVNFSDDPEFQHUSDCRDUHQINGA" {
			t.Errorf("unexpected string: %s", s)
		}
	})
	t.Run("BasicLatinSymbolCharset", func(t *testing.T) {
		r := newRandomStringReader([]byte(randstrings.BasicLatinSymbolCharset))

		var b [64]byte
		_, err := r.Read(b[:])
		if err != nil {
			t.Fatal(t)
		}
		if s := string(b[:]); s != "!:@{=.?_~-\\<$]:\"+-<{/~ '!)\\-(,+<>\\-*&\"^^,-\\$/!/.{?%-`!%_/>=?(<:," {
			t.Errorf("unexpected string: %s", s)
		}
	})
}

func ExampleReader() {
	r := &randstrings.Reader{
		RandomSource: rand.New(rand.NewSource(0)),
		Charset:      []byte(randstrings.LowerHexCharset),
	}

	var b [64]byte
	_, err := io.ReadFull(r, b[:])
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b[:]))
	// Output: 14d2afc013f24b38ef9f625e8ad4adb5b8df87e01d9615f4189e25c11458856e
}

func ExampleReader_WithMultipleCharset() {
	r := &randstrings.Reader{
		RandomSource: rand.New(rand.NewSource(0)),
		Charset:      []byte(randstrings.BasicLatinSymbolCharset + randstrings.UpperAlphabetCharset + randstrings.LowerAlphabetCharset),
	}

	var b [64]byte
	_, err := io.ReadFull(r, b[:])
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b[:]))
	// Output: !ey\vOx@gI <$&||]uu*r-vjdJ=ppM+~w\-%nxoV'-YlwP/JrCTNXPhWw> C(~Lt
}

func newRandomStringReader(cs []byte) io.Reader {
	return &randstrings.Reader{
		RandomSource: rand.New(rand.NewSource(0)),
		Charset:      cs,
	}
}
