package client

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func getFile(conn net.Conn, fileName string) {
	conn.Write([]byte(fmt.Sprintf("download %s\n", fileName)))

	outputFile, err := os.Create("store/client/" + fileName)
	if err != nil {
		log.Println(err)
	}

	// получение размера файла
	r := bufio.NewReader(conn)
	s, _ := r.ReadString('\n')
	fs, _ := strconv.Atoi(strings.ReplaceAll(s, "\n", ""))
	_, err = io.Copy(outputFile, io.LimitReader(conn, int64(fs)))
	outputFile.Close()

	msg := make([]byte, 1024)
	n, err := conn.Read(msg)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(strings.Trim(string(msg[:n]), "\n"))

}
