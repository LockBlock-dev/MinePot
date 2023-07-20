package handles

import (
	"fmt"
	"log"
	"time"

	"github.com/LockBlock-dev/MinePot/packets"
	"github.com/LockBlock-dev/MinePot/typings"
	"github.com/LockBlock-dev/MinePot/utils"
)

func HandleConnection(conn typings.ConnWrapper) {
    remoteAddrString := conn.RemoteAddr().String()
    
    defer func() {
        log.Println(remoteAddrString + " - Closing connection")

        // If the client has exceeded the packets threshold we can report it
        utils.HandleReport(conn, remoteAddrString)
        
        if err := conn.Close(); err != nil {
            log.Fatal(err)
        }
    }()

    log.Println(remoteAddrString + " - Client connected")

    for {
        packet := packets.MinecraftPacket{}
        if err := packet.Read(conn); err != nil {
            return
        }

        // "ghost" packet prevention
        if packet.Length == 0 {
            return
        }

        if conn.Config.Debug {
            log.Println(
                remoteAddrString +
                " - Received packet => Length: " +
                fmt.Sprint(packet.Length) +
                ", Id: " +
                fmt.Sprint(packet.PacketId) +
                ", Data: " +
                string(packet.Data),
            )
        }

        switch packet.PacketId {
            case 0x0:
                // Status Request packet has no data : https://wiki.vg/Protocol#Status_Request
                if len(packet.Data) == 0 && conn.Config.StatusResponse {
                    handleStatusRequest(&conn)
                } else {
                    handleHandshake(&conn, packet)
                }
            case 0x1:
                handlePing(&conn, packet)
            default:
                log.Println(remoteAddrString + " - Unknown packet received with ID: " + fmt.Sprint(packet.PacketId))
        }

        conn.PacketsReceived ++

        // We wait X ms until we try to read for another packet
        // Tested up to 1000ms (1s) with mcstatus and official MC Client
        time.Sleep(time.Duration(conn.Config.PollIntervalMs) * time.Millisecond)
    }
}
