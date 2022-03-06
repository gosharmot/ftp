package client

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func getFile(conn net.Conn, fileName string) {
	conn.Write([]byte(fmt.Sprintf("download %s\n", fileName)))

	outputFile, err := os.Create("store/client/" + fileName)
	if err != nil {
		log.Println(err)
	}

	buf := new(bytes.Buffer)
	io.Copy(outputFile, bytes.NewReader(buf.Bytes()))
	outputFile.Close()

	msg := make([]byte, 1024)
	n, err := conn.Read(msg)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(strings.Trim(string(msg[:n]), "\n"))

}
