package main

import (
	"bufio"
	"fmt"
	"net"
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

func handleConn(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	fmt.Printf("type is %T", conn)

	reqLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("failed to read string ", err)
	}

	fmt.Print(reqLine)

	html := `<html>
	<title>go http server</title>
	<h1>Go Http Server</h1>
	<p>Local Http Server listening on Port 8080</p>
	</html>`

	response := strings.Join([]string{
		"HTTP/1.1 200 OK",
		"Content-Type: text/html; charset=utf-8",
		fmt.Sprintf("Content-Length: %d", len(html)),
		"",
		html,
	}, "\n")

	conn.Write([]byte(response))

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
