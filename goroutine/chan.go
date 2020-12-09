package goroutine

import (
	"io"
	"log"
	"net"
	"os"
)

func ChanEntry() {
	co, err := net.Dial("tcp", "localhost:8001")
	if err != nil {
		log.Println(err)
		return
	}

	// 使用类型断言，拿到 *net.TCPConn 的值
	conn := co.(*net.TCPConn)

	done := make(chan struct{})
	go func() {
		mustCopy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	// 调用写关闭。只有 *net.TCPConn 才有这个方法，这是类型断言的目的。
	conn.CloseWrite()
	<-done
	log.Println("exit")
}

func mustCopy(w io.Writer, r io.Reader)  {
	buf := [64]byte{}
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			log.Println(err)
			break
		}

		_, err = w.Write(buf[0:n])
		if err != nil {
			log.Println(err)
			break
		}
	}

}
