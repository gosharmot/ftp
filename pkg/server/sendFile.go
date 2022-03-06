package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func sendFile(conn net.Conn, name string) {
	file, err := os.Open("store/server/" + name)
	if err != nil {
		log.Println(err.Error())
		conn.Write([]byte(err.Error()))
		return
	}
	defer file.Close()

	stat, _ := file.Stat()
	fmt.Fprintln(conn, stat.Size())

	io.Copy(conn, file)

	fmt.Fprintf(conn, "Загрузка файла '%s' завершена успешно", name)
	log.Printf("Отправка файла '%s' завершена успешно", name)
}
