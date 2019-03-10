// client project main.go
package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	for {
		var source string
		fmt.Print("Введите слово: ")
		_, err := fmt.Scanln(&source)
		if err != nil {
			fmt.Println("Некорректный ввод", err)
			continue
		}
		t0 := time.Now()
		// отправляем сообщение серверу
		if n, err := conn.Write([]byte(source)); n == 0 || err != nil {
			fmt.Println(err)
			return
		}
		// получем ответ
		fmt.Print("Размер: ")
		buff := make([]byte, (1024 * 1024 * 500))
		n, err := conn.Read(buff)
		if err != nil {
			break
		}
		fmt.Println(n)

		fmt.Print("Последний символ: ")
		fmt.Println(buff[500*1024*1024-1])

		t1 := time.Now()
		fmt.Print("Время работы: ")
		fmt.Println((t1.Second() - t0.Second()))

		fmt.Println()
	}
}
