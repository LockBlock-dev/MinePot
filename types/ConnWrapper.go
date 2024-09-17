package types

import (
	"net"

	mcNet "github.com/Tnze/go-mc/net"
)

type ConnWrapper struct {
	mcNet.Conn
	SrcAddr          net.Addr
	DestAddr         net.Addr
	Config           *Config
	PacketsReceived  int
	ReceivedProtocol int
	DidHandshake     bool
	DidPing          bool
}
