package packets

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

type MinecraftPacket struct {
    Length    int32
    PacketID  int32
    Data      []byte
}

func (p *MinecraftPacket) Read(conn net.Conn) error {
    r := bufio.NewReader(conn)
    
    // Read packet length
    var length uint64
    length, err := binary.ReadUvarint(r)
    if err != nil {
        return fmt.Errorf("failed to read packet length: %w", err)
    }
    p.Length = int32(length)

    // Read packet data
    data := make([]byte, p.Length)
    if _, err := io.ReadFull(r, data); err != nil {
        return fmt.Errorf("failed to read packet data: %w", err)
    }

    // Extract packet ID
    var packetID uint64
    var n int
    packetID, n = binary.Uvarint(data)
    if n <= 0 {
        return fmt.Errorf("failed to read packet ID")
    }
    p.PacketID = int32(packetID)

    // Extract packet data
    p.Data = data[n:]

    return nil
}

func (p *MinecraftPacket) Write(conn net.Conn) error {
    // Calculate packet ID length
    idBuf := make([]byte, binary.MaxVarintLen32)
    idLen := binary.PutUvarint(idBuf, uint64(p.PacketID))

    // Calculate packet length
    length := len(p.Data) + idLen
    lengthBuf := make([]byte, binary.MaxVarintLen32)
    lengthLen := binary.PutUvarint(lengthBuf, uint64(length))
    
    // Write packet length
    if _, err := conn.Write(lengthBuf[:lengthLen]); err != nil {
        return fmt.Errorf("failed to write packet length: %w", err)
    }

    // Write packed ID
    if _, err := conn.Write(idBuf[:idLen]); err != nil {
        return fmt.Errorf("failed to write packet ID: %w", err)
    }

    // Write packet data
    if _, err := conn.Write(p.Data); err != nil {
        return fmt.Errorf("failed to write packet data: %w", err)
    }

    return nil
}
