package msgpack

import (
	"bitbucket.stressedsharks.com/d.kellenschwiler/familiarisation/shortener"
	"github.com/vmihailenco/msgpack"
	"github.com/pkg/errors"
)

type Redirect struct {}

func (r *Redirect)) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	if err := msgpack.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Decode")
	}
	return redirect, nil
}

func (r *Redirect) Encode(input *shortener.redirect) ([]byte, error) {
	rawMsg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Encode")
	}
	return rawMsg, nil
}