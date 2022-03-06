package client

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"strings"
)

func sendFile(conn net.Conn, fileName string) {

	content, err := ioutil.ReadFile("store/client/" + fileName)
	if err != nil {
		log.Println(err)
		return
	}

	conn.Write([]byte(fmt.Sprintf("upload %s %d\n", fileName, len(content))))

	io.Copy(conn, bytes.NewReader(content))

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(strings.Trim(string(buf[:n]), "\n"))
}
