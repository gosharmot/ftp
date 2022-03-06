package server

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
	"strings"
	"time"
)

func HandleConnection(conn net.Conn) {
	log.Printf("%s Подключился", conn.RemoteAddr().String())
	err := json.NewEncoder(conn).Encode("Соединение с сервером установленно")
	if err != nil {
		log.Fatal(err)
	}
	buf := bufio.NewScanner(conn)
	for buf.Scan() {

		commandArr := strings.Fields(strings.Trim(buf.Text(), "\n"))

		conn.SetDeadline(time.Now().Add(time.Minute * 5))

		switch strings.ToLower(commandArr[0]) {

		case "download":
			log.Printf("Скачивание файла '%s'\n", commandArr[1])
			sendFile(conn, commandArr[1])

		case "upload":
			log.Printf("Загрузка файла '%s'\n", commandArr[1])
			getFile(conn, commandArr[1], commandArr[2])
		case "ls":
			log.Println("ls")
			getListFiles(conn)

		case "exit":
			log.Printf("%s Отключился\n", conn.RemoteAddr().String())
			return
		}
	}
}
