package packets

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type PingRequest struct {
    Magic int64
}

func (p *PingRequest) Read(data []byte) error {
    r := bytes.NewReader(data)

	// Read Ping Request magic number
	var magic int64
	if err := binary.Read(r, binary.BigEndian, &magic); err != nil {
		return fmt.Errorf("failed to read Ping Request magic number: %w", err)
	}
	p.Magic = magic

    return nil
}
