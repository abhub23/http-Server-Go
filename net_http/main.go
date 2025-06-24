package main

import (
	"bufio"
	"fmt"
	"net"
	"reflect"

	"github.com/pkg/errors"
)

func handleConn(conn net.Conn) {
	defer conn.Close().Error()

	reader := bufio.NewReader(conn)

	reqLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("failed to read string ", err)
	}

	fmt.Print(reqLine)

	body := `<html>
	<h1>Go Http Server</h1>
	<p>Local Http Server listening on Port 8080</p>
	</html>`

	response := `HTTP/1.1 200 OK\r\n` + fmt.Sprintf()

	
}

func main() {

	Port := ":8080"

	addr, err := net.ResolveTCPAddr("tcp", Port)
	if err != nil {
		panic(errors.Wrap(err, "error in creating a tcp connection"))
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(errors.Wrap(err, "error in tcp listener"))
	}
	defer listener.Close()
	fmt.Printf("tcp server listening on Port %s \n", Port)

	fmt.Println(listener)
	fmt.Println(reflect.TypeOf(listener))

	for {
		ln, err := listener.Accept()
		if err != nil {
			fmt.Println("error occured in accepting connections", err)
		}
		go handleConn(ln)
	}

}
