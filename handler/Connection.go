package handler

import (
	"log"

	"github.com/LockBlock-dev/MinePot/types"
	"github.com/LockBlock-dev/MinePot/util"
)

func HandleConnection(conn types.ConnWrapper) {
	remoteAddrString := conn.Conn.Socket.RemoteAddr().String()

	defer func() {
		log.Println(remoteAddrString + " - Closing connection")

		// If the client has exceeded the packets threshold we can report it
		util.HandleReport(conn, remoteAddrString)

		if err := conn.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println(remoteAddrString + " - Client connected")

	nextState := handleHandshake(&conn)
	if nextState == -1 {
		return
	}

	switch nextState {
	case 1:
		handleServerListPing(&conn)
	}

}
