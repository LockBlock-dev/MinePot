package handler

import (
	"fmt"
	"log"

	"github.com/LockBlock-dev/MinePot/types"
	"github.com/Tnze/go-mc/net/packet"
)

func handleHandshake(conn *types.ConnWrapper) int {
	var (
		Protocol, Intention packet.VarInt
		ServerAddress       packet.String
		ServerPort          packet.UnsignedShort
	)
	var p packet.Packet

	conn.PacketsReceived++

	// Handle Handshake packet : https://wiki.vg/Server_List_Ping#Handshake
	if err := conn.ReadPacket(&p); err != nil {
		log.Println("Failed to parse Handshake packet:", err)
		return -1
	}

	if err := p.Scan(&Protocol, &ServerAddress, &ServerPort, &Intention); err != nil {
		log.Println("Failed to parse Handshake data:", err)
		return -1
	}

	if conn.Config.Debug {
		log.Println(
			conn.SrcAddr.String() +
				" - Received Handshake packet => Protocol version: " +
				fmt.Sprint(Protocol) +
				", Server address: " +
				string(ServerAddress) +
				", Server port: " +
				fmt.Sprint(ServerPort) +
				", Next state: " +
				fmt.Sprint(Intention),
		)
	}

	conn.ReceivedProtocol = int(Protocol)
	conn.DidHandshake = true

	return int(Intention)
}
