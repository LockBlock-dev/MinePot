package handles

import (
	"fmt"
	"log"

	"github.com/LockBlock-dev/MinePot/packets"
	"github.com/LockBlock-dev/MinePot/typings"
)

func handleHandshake(conn *typings.ConnWrapper, receivedPacket packets.MinecraftPacket) {
    // Handle Handshake packet : https://wiki.vg/Server_List_Ping#Handshake
    handshake := packets.Handshake{}
    if err := handshake.Read(receivedPacket.Data); err != nil {
        log.Println("Failed to parse Handshake packet:", err)
        return
    }

    if conn.Config.Debug {
        log.Println(
            conn.RemoteAddr().String() +
            " - Received Handshake packet => Protocol version: " +
            fmt.Sprint(handshake.ProtocolVersion) +
            ", Server address: " +
            string(handshake.ServerAddress) +
            ", Server port: " +
            fmt.Sprint(handshake.ServerPort) +
            ", Next state: " +
            fmt.Sprint(handshake.NextState),
        )
    }

    conn.ReceivedProtocol = int(handshake.ProtocolVersion)
    conn.DidHandshake = true

    if handshake.NextState == 1 {
        handleStatusRequest(conn)
    }
}
