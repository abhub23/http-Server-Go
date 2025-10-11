package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/pkg/errors"
)

func handleConn(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	fmt.Printf("type is %T", conn)

	//Reads the response string line by line
	_, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("failed to read string ", err)
	}

	html, err := os.ReadFile("../index.html")
	if err != nil {
		fmt.Println("Error in reading HTML file ", err)
	}

	//Seperates the join by /n so the Buffer reader will read accordingly
	response := strings.Join([]string{
		"HTTP/1.1 200 OK",
		"Content-Type: text/html; charset=utf-8",
		fmt.Sprintf("Content-Length: %d", len(html)),
		"",
		string(html),
	}, "\n")

	conn.Write([]byte(response))

}

func main() {

	Port := ":8080"

	//Open a Tcp Conn on a Port
	addr, err := net.ResolveTCPAddr("tcp", Port)
	if err != nil {
		panic(errors.Wrap(err, "error in creating a tcp connection"))
	}

	//Tcp Listener
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(errors.Wrap(err, "error in tcp listener"))
	}

	defer listener.Close()
	fmt.Printf("tcp server listening on Port %s \n", Port)

	// The for loop is required to continuously accept and handle incoming connections.
	// Without it, the server would accept only a single connection and then exit.
	for {
		ln, err := listener.Accept()
		if err != nil {
			fmt.Println("error occured in accepting connections", err)
		}
		go handleConn(ln)
	}

}
