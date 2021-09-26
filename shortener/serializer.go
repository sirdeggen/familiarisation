package shortener

// RedirectSerializer is ...
type RedirectSerializer interface {
	Decode(input []byte) (*Redirect, error)
	Encode(input *Redirect) ([]byte, error)
}
