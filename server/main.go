package main

import (
	"fmt"
	"net"
)
import "net/http"

func main() {
	fmt.Println("Hello!")

	addr := net.TCPAddr{
		Port: 8999,
	}

	listener, err := net.ListenTCP("tcp", &addr)

	if err != nil {
		fmt.Println(err.Error())
	}

	err = http.Serve(
		listener,
		handler{},
	)

	if err != nil {
		fmt.Println(err.Error())
	}
}

type handler struct{}

func (h handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(200)
	rw.Write([]byte("Hella!"))
}
