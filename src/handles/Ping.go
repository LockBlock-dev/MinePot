package handles

import (
	"encoding/binary"
	"fmt"
	"log"

	"github.com/LockBlock-dev/MinePot/packets"
	"github.com/LockBlock-dev/MinePot/typings"
)

func handlePing(conn *typings.ConnWrapper, receivedPacket packets.MinecraftPacket) {
	remoteAddrString := conn.RemoteAddr().String()

	// Handle Ping Request packet : https://wiki.vg/Server_List_Ping#Ping_Request
    ping := packets.PingRequest{}
    if err := ping.Read(receivedPacket.Data); err != nil {
        log.Println("Failed to parse Ping Request packet:", err)
        return
    }

	if conn.Config.Debug {
		log.Println(remoteAddrString + " - Received Ping Request packet => Magic number: " + fmt.Sprint(ping.Magic))
	}

	magicBuf := make([]byte, 8)
	binary.BigEndian.PutUint64(magicBuf, uint64(ping.Magic))

	// Send Pong Response packet : https://wiki.vg/Server_List_Ping#Pong_Response
	responsePacket := packets.MinecraftPacket{
	    PacketID: 0x01, // Pong Response
	    Data: magicBuf, // Sent by the client
	}
	
	// Send the packet to the client
	if err := responsePacket.Write(conn); err != nil {
	    log.Println("Failed to send Pong Response packet to client:", err)
		return
	}

	if conn.Config.Debug {
		log.Println(remoteAddrString + " - Sent Pong Response packet")
	}

	conn.DidPing = true
}
