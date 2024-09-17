package handler

import (
	"encoding/json"
	"log"
	"net"
	"strings"

	"github.com/LockBlock-dev/MinePot/internal/minecraft"
	"github.com/LockBlock-dev/MinePot/types"
	"github.com/LockBlock-dev/MinePot/util"
	"github.com/Tnze/go-mc/net/packet"
)

func handleStatusRequest(conn *types.ConnWrapper) {
	remoteAddrString := conn.SrcAddr.String()

	// Handle Status Request packet : https://wiki.vg/Server_List_Ping#Status_Request
	if conn.Config.Debug {
		log.Println(remoteAddrString + " - Received Status Request packet")
	}

	statusResponseData := conn.Config.StatusResponseData
	versions := minecraft.GetAllVersions()
	protocolMapping := minecraft.GetAllVersionsMapping()

	if conn.Config.RandomVersion {
		key := util.RandRange(0, len(versions)-1)
		version := versions[key]
		protocol := protocolMapping[version]

		statusResponseData.Version.Name = version
		statusResponseData.Version.Protocol = protocol
	} else {
		if conn.Config.StatusResponseData.Version.Protocol == -1 {
			for k, v := range protocolMapping {
				if v == conn.ReceivedProtocol {
					statusResponseData.Version.Name = k
				}
			}

			statusResponseData.Version.Protocol = conn.ReceivedProtocol
		}
	}

	ipSubstr := "%IP%"

	statusResponseData.Description = strings.Replace(
		statusResponseData.Description,
		ipSubstr,
		conn.SrcAddr.(*net.TCPAddr).IP.String(),
		1,
	)

	// Add the favicon from its file
	err := util.GetFavicon(conn.Config)
	if err != nil {
		statusResponseData.Favicon = ""
	}

	// Sending Status Response packet : https://wiki.vg/Server_List_Ping#Status_Response
	status := statusResponseData

	jsonData, err := json.Marshal(status)
	if err != nil {
		log.Println("Failed to transform JSON Status Response:", err)
	}

	if err := conn.WritePacket(packet.Marshal(0x00, packet.String(jsonData))); err != nil {
		log.Println("Failed to send Status Response packet to client:", err)
	}

	if conn.Config.Debug {
		log.Println(remoteAddrString + " - Sent Status Response packet")
	}
}
