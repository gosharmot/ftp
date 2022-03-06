package main

import (
	"ftp/pkg/server"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8085")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Сервер запущен")
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		conn.SetDeadline(time.Now().Add(time.Minute * 3))
		if err != nil {
			log.Println(err)
			continue
		}
		go server.HandleConnection(conn)
	}
}
