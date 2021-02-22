package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {
	sendSizeKB := flag.Int("size", 1, "KB")
	host := flag.String("host", "127.0.0.1", "Server host")
	port := flag.Int("port", 23000, "Server port")
	flag.Parse()

	sendSize := *sendSizeKB * 1024

	bytes := make([]byte, sendSize)
	hostPort := fmt.Sprintf("%s:%d", *host, *port)
	conn, err := net.Dial("tcp4", hostPort)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connect to %s", hostPort)

	n, err := conn.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
	if n != sendSize {
		log.Printf("unexpected write length: %d", n)
	}

	log.Printf("Finished to write the %d bytes", n)
}
