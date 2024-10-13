package network

import (
	"github.com/ipfs/boxo/bitswap/message"
	"github.com/libp2p/go-libp2p/core/protocol"
)

type NetOpt func(*Settings)

type Settings struct {
	ProtocolPrefix     protocol.ID
	SupportedProtocols []protocol.ID
	Decoder            message.Decoder
}

func Prefix(prefix protocol.ID) NetOpt {
	return func(settings *Settings) {
		settings.ProtocolPrefix = prefix
	}
}

func SupportedProtocols(protos []protocol.ID) NetOpt {
	return func(settings *Settings) {
		settings.SupportedProtocols = protos
	}
}

func Decoder(bmsg message.Decoder) NetOpt {
	return func(settings *Settings) {
		settings.Decoder = bmsg
	}
}
