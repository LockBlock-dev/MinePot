package packets

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type Handshake struct {
    ProtocolVersion int32
	ServerAddress string
	ServerPort uint16
	NextState int
}

func (p *Handshake) Read(data []byte) error {
    r := bytes.NewReader(data)

    // Read protocol version
    var version uint64
    version, err := binary.ReadUvarint(r)
    if err != nil {
        return fmt.Errorf("failed to read protocol version: %w", err)
    }
    p.ProtocolVersion = int32(version)

    // Read server address length
    serverAddressLen, err := binary.ReadUvarint(r)
    if err != nil {
        return fmt.Errorf("failed to read server address length: %w", err)
    }

    // Read server address
    serverAddressBytes := make([]byte, serverAddressLen)
    if _, err := io.ReadFull(r, serverAddressBytes); err != nil {
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
    var nextState uint64
    nextState, err = binary.ReadUvarint(r)
    if err != nil {
        return fmt.Errorf("failed to read next state: %w", err)
    }
    p.NextState = int(nextState)

    return nil
}
