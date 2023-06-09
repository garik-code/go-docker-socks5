package main

import (
	"log"
	"os"
	"fmt"

	"github.com/things-go/go-socks5"
)

func main() {
	// Create a SOCKS5 server
	server := socks5.NewServer(
		socks5.WithLogger(socks5.NewLogger(log.New(os.Stdout, "socks5: ", log.LstdFlags))),
	)

	fmt.Printf("Test %d\n", 23)

	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe("tcp", ":8000"); err != nil {
		panic(err)
	}
}