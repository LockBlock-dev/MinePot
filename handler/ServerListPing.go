package handler

import (
	"fmt"
	"log"

	"github.com/LockBlock-dev/MinePot/types"
	"github.com/Tnze/go-mc/data/packetid"
	"github.com/Tnze/go-mc/net/packet"
)

func handleServerListPing(conn *types.ConnWrapper) {
	var p packet.Packet
	remoteAddrString := conn.Conn.Socket.RemoteAddr().String()

	for i := 0; i < 2; i++ {

		conn.PacketsReceived++

		// Handle Server List Ping following packet : https://wiki.vg/Server_List_Ping
		if err := conn.ReadPacket(&p); err != nil {
			log.Println("Failed to parse Server List Ping packet:", err)
			return
		}

		if conn.Config.Debug {
			log.Println(
				remoteAddrString +
					" - Received packet => Length: " +
					fmt.Sprint(len(p.Data)) +
					", Id: " +
					fmt.Sprint(p.ID) +
					", Data: " +
					string(p.Data),
			)
		}

		switch packetid.ClientboundPacketID(p.ID) {
		case packetid.ClientboundStatusResponse:
			if conn.Config.StatusResponse {
				handleStatusRequest(conn)
			}
		case packetid.ClientboundStatusPongResponse:
			handlePing(conn, &p)
		}
	}
}
