package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path"
	"time"

	"github.com/LockBlock-dev/MinePot/handles"
	"github.com/LockBlock-dev/MinePot/typings"
	"github.com/LockBlock-dev/MinePot/utils"
	"github.com/muesli/cache2go"
)

func main() {
    config, err := utils.GetConfig()
    if err != nil {
        log.Fatal(err)
    }

    var file *os.File

    if config.WriteLogs {
        // Open logs file
        file, err = os.OpenFile(config.LogFile, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644) // 644 = rw-,r--,r--
        if err != nil {
            log.Fatal(err)
        }
        defer file.Close()
    }

    if config.WriteHistory {
        // Open history file
        historyFile, err := os.OpenFile(config.HistoryFile, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644) // 644 = rw-,r--,r--
        if err != nil {
            log.Fatal(err)
        }
        defer historyFile.Close()

        _, err = historyFile.WriteString("datetime, ip, packets_count, reported, handshake, ping")
        if err != nil {
            log.Fatal("Failed to write history headers:", err)
        }
    }

    // Setup the cache
    _ = cache2go.Cache("MinePot")
    
    // Listen for incoming connections on TCP port X (see config.json)
    address := fmt.Sprintf(":%d", config.Port)
    listener, err := net.Listen("tcp", address)
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        listener.Close()
    }()

    log.Printf("Server listening on port %d\nYou can edit the config at /etc/minepot/config.json", config.Port)


    if config.WriteLogs {
        // Logs the logs file path
        cwd, err := os.Getwd()
        if err == nil {
            log.Println("Find the logs at: " + path.Join(cwd, config.LogFile))
        }

        // Setup logs to a file
        log.SetOutput(file)
    }

    for {
        // Wait for a client to connect
        conn, err := listener.Accept()
        if err != nil {
            log.Println(err)
            return
        }

        // Set a timeout of X seconds (see config.json)
        conn.SetDeadline(time.Now().Add(time.Duration(config.IdleTimeoutS) * time.Second))

        connWrapper := typings.ConnWrapper{
            Conn:   conn,
            Config: config,
        }

        // Start a new goroutine to handle the connection
        go handles.HandleConnection(connWrapper)
    }
}
