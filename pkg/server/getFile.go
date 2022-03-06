package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

func getFile(conn net.Conn, name string, fileSize string) {

	fs, err := strconv.ParseInt(fileSize, 10, 64)
	if err != nil || fs == -1 {
		log.Println(err)
		return
	}

	outputFile, err := os.Create("store/server/" + name)
	if err != nil {
		log.Println(err)
		return
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, io.LimitReader(conn, fs))
	if err != nil {
		fmt.Fprintf(conn, "Ошибка при загрузки файла '%s': %s", name, err.Error())
	}

	fmt.Fprintf(conn, "Загрузка файла '%s' успешно завершена", name)
	log.Printf("Скачивание файла '%s' успешно завершена\n", name)

}
