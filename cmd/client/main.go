package main

import (
	"ftp/pkg/client"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8085")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client.HandleConnection(conn)
}
