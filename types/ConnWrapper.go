package types

import "github.com/Tnze/go-mc/net"

type ConnWrapper struct {
	net.Conn
	Config           *Config
	PacketsReceived  int
	ReceivedProtocol int
	DidHandshake     bool
	DidPing          bool
}
