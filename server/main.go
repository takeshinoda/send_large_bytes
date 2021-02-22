package main

import (
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)


	ln, err := net.Listen("tcp4", "0.0.0.0:23000")
	if err != nil {
		log.Fatal(err)
	}
	defer func () {
		if err := ln.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	log.Printf("Starting the server %s", ln.Addr().String())

	go func () {
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Fatal(err)
			}

			go readConn(conn)
		}
	}()

	log.Printf("Received signal: %v", <-ch)
}

func readConn(conn net.Conn) {
	log.Printf("Connected")
	defer log.Printf("Finished to read the conn")

	const bufSize = 10 * 1024

	readLength := 0
	for {
		buf := make([]byte, bufSize)// 10 KB

		if err := conn.SetReadDeadline(time.Now().Add(5 * time.Second)); err !=nil {
			log.Fatal(err)
		}
		n, err := conn.Read(buf[:bufSize])
		if err != nil {
			if err != io.EOF {
				log.Printf("unexpected error: %v, read lengsh is %d", err, readLength)
				return
			}
			readLength += n
			// EOF
			log.Printf("received EOF, read length is %d", readLength)
			break
		}
		readLength += n
	}
}
