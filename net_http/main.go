package main

import (
	"fmt"
	"net"
	"github.com/pkg/errors"
	"reflect"
)

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

}
