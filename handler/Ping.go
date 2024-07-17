package handler

import (
	"fmt"
	"log"
	"time"

	"github.com/LockBlock-dev/MinePot/types"
	"github.com/LockBlock-dev/MinePot/util"
	"github.com/Tnze/go-mc/net/packet"
)

func handlePing(conn *types.ConnWrapper, p *packet.Packet) {
	var Magic packet.Long

	remoteAddrString := conn.Conn.Socket.RemoteAddr().String()

	// Handle Ping Request packet : https://wiki.vg/Server_List_Ping#Ping_Request
	if err := p.Scan(&Magic); err != nil {
		log.Println("Failed to parse Ping data:", err)
		return
	}

	if conn.Config.Debug {
		log.Println(remoteAddrString + " - Received Ping Request packet => Magic number: " + fmt.Sprint(Magic))
	}

	// Artificial server ping
	time.Sleep(time.Duration(util.RandRange(conn.Config.PingDelayMinMs, conn.Config.PingDelayMaxMs)) * time.Millisecond)

	// Send Pong Response packet : https://wiki.vg/Server_List_Ping#Pong_Response
	if err := conn.WritePacket(*p); err != nil {
		log.Println("Failed to send Pong Response packet to client:", err)
		return
	}

	if conn.Config.Debug {
		log.Println(remoteAddrString + " - Sent Pong Response packet")
	}

	conn.DidPing = true
}
