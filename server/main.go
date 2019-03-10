// server project main.go

package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":4545")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		go handleConnection(conn) // запускаем горутину для обработки запроса
	}
}

// обработка подключения
func handleConnection(conn net.Conn) {
	fmt.Println("Подключение установлено!")
	defer conn.Close()
	for {
		// считываем полученные в запросе данные
		input := make([]byte, (1024 * 1024 * 500))
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("Read error:", err)
			break
		}
		target := make([]byte, (1024 * 1024 * 500))
		for i := 0; i <= 1024*1024*10-1; i++ {
			target[i] = 250
		}
		target[500*1024*1024-1] = 251
		fmt.Println("Отправка")
		conn.Write([]byte(target))
	}
}

//var input string
//fmt.Scanf("%v", &input)
