package packets

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/LockBlock-dev/MinePot/typings"
)

type Handshake struct {
    ProtocolVersion int
	ServerAddress string
	ServerPort uint16
	NextState int
}

func (p *Handshake) Read(data []byte) error {
    r := bytes.NewReader(data)
    reader := typings.NewCustomReader(r)

    // Read protocol version
    version, err := reader.ReadVarInt()
    if err != nil {
        return fmt.Errorf("failed to read protocol version: %w", err)
    }
    p.ProtocolVersion = version

    // Read server address length
    serverAddressLen, err := reader.ReadVarInt()
    if err != nil {
        return fmt.Errorf("failed to read server address length: %w", err)
    }

    // Read server address
    serverAddressBytes := make([]byte, serverAddressLen)
    if _, err := r.Read(serverAddressBytes); err != nil {
        return fmt.Errorf("failed to read server address: %w", err)
    }
    p.ServerAddress = string(serverAddressBytes)

    // Read server port
    var port uint16
    if err := binary.Read(r, binary.BigEndian, &port); err != nil {
        return fmt.Errorf("failed to read server port: %w", err)
    }
    p.ServerPort = port

    // Read next state
    nextState, err := reader.ReadVarInt()
    if err != nil {
        return fmt.Errorf("failed to read next state: %w", err)
    }
    p.NextState = nextState

    return nil
}
