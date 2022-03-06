package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func HandleConnection(conn net.Conn) {
	var msg string
	err := json.NewDecoder(conn).Decode(&msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
	stdReader := bufio.NewReader(os.Stdin)

	printCommand()
	for {
		fmt.Printf("\033[0mftp> ")
		cmd, _ := stdReader.ReadString('\n')
		cmdArr := strings.Fields(strings.Trim(cmd, "\n"))
		if len(cmdArr) > 0 {
			switch strings.ToLower(cmdArr[0]) {

			case "u", "upload":
				if len(cmdArr) != 2 {
					fmt.Println("введите команду и имя файла")
					continue
				}
				sendFile(conn, cmdArr[1])

			case "d", "download":
				if len(cmdArr) != 2 {
					fmt.Println("введите команду и имя файла")
					continue
				}
				getFile(conn, cmdArr[1])

			case "ls":
				conn.Write([]byte(cmd))
				buffer := make([]byte, 4096)
				n, _ := conn.Read(buffer)
				fmt.Print(string(buffer[:n]))

			case "e", "exit":
				conn.Write([]byte("exit\n"))
				return

			default:
				fmt.Println("Неверная команда, используйте: upload u | download d | ls | exit e")
			}
		}
	}
}

func printCommand() {
	fmt.Println("Команды:")
	fmt.Printf("%-30s %s\n", "u, upload 'file name'", "загрузка файла")
	fmt.Printf("%-30s %s\n", "d, download 'file name'", "скачивание файла")
	fmt.Printf("%-30s %s\n", "ls", "файлы на сервере")
	fmt.Printf("%s%-30s %s\n", "\033[31m", "e, exit", "выход")
}
