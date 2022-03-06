package client

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func sendFile(conn net.Conn, fileName string) {
	file, err := os.Open("store/client/" + fileName)
	if err != nil {
		log.Println(err)
		return
	}
	stats, _ := file.Stat()

	conn.Write([]byte(fmt.Sprintf("upload %s %d\n", fileName, stats.Size())))

	io.Copy(conn, file)

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(strings.Trim(string(buf[:n]), "\n"))
}
