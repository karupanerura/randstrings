# randstrings
--
    import "github.com/karupanerura/randstrings"


## Usage

```go
const (
	NumbersCharset          = "0123456789"
	LowerHexCharset         = NumbersCharset + "abcdef"
	UpperHexCharset         = NumbersCharset + "ABCDEF"
	LowerAlphabetCharset    = "abcdefghijklmnopqrstuvwxyz"
	UpperAlphabetCharset    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	BasicLatinSymbolCharset = " !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
)
```

#### type Reader

```go
type Reader struct {
	RandomSource io.Reader
	Charset      []byte
}
```


#### func (*Reader) Read

```go
func (r *Reader) Read(b []byte) (n int, err error)
```
