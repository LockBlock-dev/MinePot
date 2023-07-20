package handles

import (
	"encoding/binary"
	"encoding/json"
	"log"
	"net"
	"strings"

	"github.com/LockBlock-dev/MinePot/packets"
	"github.com/LockBlock-dev/MinePot/typings"
	"github.com/LockBlock-dev/MinePot/utils"
)

func handleStatusRequest(conn *typings.ConnWrapper) {
	remoteAddrString := conn.RemoteAddr().String()

	// Handle Status Request packet : https://wiki.vg/Server_List_Ping#Status_Request
	if conn.Config.Debug {
		log.Println(remoteAddrString + " - Received Status Request packet")
	}

	if conn.Config.StatusResponseData.Version.Protocol == -1 {
		conn.Config.StatusResponseData.Version.Protocol = conn.ReceivedProtocol
	}

	ipSubstr := "%IP%"
	desc := conn.Config.StatusResponseData.Description
	oldDesc := desc

	if strings.Contains(desc, ipSubstr) {
		conn.Config.StatusResponseData.Description = strings.Replace(
			desc,
			ipSubstr,
			conn.RemoteAddr().(*net.TCPAddr).IP.String(),
			1,
		)
	}

	// Add the favicon from its file
	err := utils.GetFavicon(conn.Config); if err != nil {
		conn.Config.StatusResponseData.Favicon = ""
	}

	// Sending Status Response packet : https://wiki.vg/Server_List_Ping#Status_Response
	status := conn.Config.StatusResponseData

	jsonData, err := json.Marshal(status)

	if err != nil {
		log.Println("Failed to transform JSON Status Response:", err)
	}

	// Encoding the json message length as a Varint
	dataLengthBuf := make([]byte, binary.MaxVarintLen32)
	dataLengthLen := binary.PutUvarint(dataLengthBuf, uint64(len(jsonData)))

	// We prefix the json message with its length
	data := make([]byte, dataLengthLen + len(jsonData))
	offset := copy(data, dataLengthBuf[:dataLengthLen])
	copy(data[offset:], jsonData)

	responsePacket := packets.MinecraftPacket{
		PacketId: 0x00, // Status Response
		Data: data,
	}
	
	if err := responsePacket.Write(conn); err != nil {
		log.Println("Failed to send Status Response packet to client:", err)
	}

	if conn.Config.Debug {
		log.Println(remoteAddrString + " - Sent Status Response packet")
	}

	conn.Config.StatusResponseData.Description = oldDesc
}
