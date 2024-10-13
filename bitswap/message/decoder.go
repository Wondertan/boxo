package message

import "github.com/libp2p/go-msgio"

type Decoder interface {
	FromMsgReader(r msgio.Reader) (BitSwapMessage, error)
}

type DefaultDecoder struct{}

func (d *DefaultDecoder) FromMsgReader(r msgio.Reader) (BitSwapMessage, error) {
	return FromMsgReader(r)
}
