package packets

import (
	"fmt"
	"log"

	"github.com/LockBlock-dev/MinePot/typings"
)

type MinecraftPacket struct {
    Length    int
    PacketId  int
    Data      []byte
}

func (p *MinecraftPacket) Read(conn *typings.ConnWrapper) error {
    reader := typings.NewCustomReader(conn)

    // Read packet length
    length, err := reader.ReadVarInt()
    if err != nil {
        return fmt.Errorf("failed to read packet length: %w", err)
    }
    p.Length = length

    // Read packet id
    var packetId int
    packetId, err = reader.ReadVarInt()
    if err != nil {
        return fmt.Errorf("failed to read packet id: %w", err)
    }
    p.PacketId = packetId

    // Read packet data
    // Packet total length - id length
    data := make([]byte, p.Length - 1)
    if _, err := conn.Read(data); err != nil {
        return fmt.Errorf("failed to read packet data: %w", err)
    }
    p.Data = data

    if conn.Config.Debug {
        remoteAddrString := conn.RemoteAddr().String()
    
        logStr := ""

        for i, val := range data {
            logStr += fmt.Sprintf("%#x", val)

            if i != (len(data) - 1) {
                logStr += ", "
            }
        }

        log.Print(remoteAddrString + " - Raw packet data: [" + logStr + "]")
    }

    return nil
}

func (p *MinecraftPacket) Write(conn *typings.ConnWrapper) error {
    writer := typings.NewCustomWriter(conn)

    // Write packet length
    // Packet data length + id length
    if err := writer.WriteVarInt(len(p.Data) + 1); err != nil {
        return fmt.Errorf("failed to write packet length: %w", err)
    }

    // Write packet id
    if err := writer.WriteVarInt(p.PacketId); err != nil {
        return fmt.Errorf("failed to write packet id: %w", err)
    }

    // Write packet data
    if _, err := conn.Write(p.Data); err != nil {
        return fmt.Errorf("failed to write packet data: %w", err)
    }

    return nil
}
