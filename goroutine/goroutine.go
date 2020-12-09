package goroutine

import (
	"io"
	"log"
	"net"
)

func RoutineEntry() {
	echoServer()
}

func echoServer() {
	listener, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleFunc(conn)
	}
}

func handleFunc(conn net.Conn) {
	io.Copy(conn, conn)
	conn.Close()
}
