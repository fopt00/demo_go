package proxy

import (
	"io"
	"log"
	"net"
)

type Server struct {
	DestAddr string

	OnNewConn func(conn net.Conn) (net.Conn, error)

	OnConnClose func(conn net.Conn)

	RequestDumpWriter io.WriteCloser

	ResponseDumpWriter io.WriteCloser

	NewDecoderFunc NewDecoderFunc
}

func (s *Server) Serve(l net.Listener) error {
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Accept error:", err.Error())
			continue
		}
		go s.handler(conn)
	}
}

func (s *Server) handler(inConn net.Conn) {
	var err error
	if s.OnNewConn != nil {
		inConn, err = s.OnNewConn(inConn)
	}

	if err != nil {
		return
	}

	if s.OnConnClose != nil {
		defer s.OnConnClose(inConn)
	}

	var decoder Decoder
	if s.NewDecoderFunc != nil {
		decoder = s.NewDecoderFunc(inConn)
	} else {
		decoder = NewNopDecoderFunc(inConn)
	}

	errchan := make(chan error, 2)

	copy := func(dst io.Writer, src io.Reader) {
		_, err := io.Copy(dst, src)
		errchan <- err
	}

	outConn, err := net.DialTimeout("tcp", s.DestAddr, Duration)
	if err != nil {
		log.Println(err.Error())
		return
	}

	defer outConn.Close()

	var wOut io.Writer = outConn
	if s.RequestDumpWriter != nil {
		w1 := NewDecoderWriter(s.RequestDumpWriter, decoder.Request)
		defer w1.Close()
		wOut = io.MultiWriter(outConn, w1)
	}
	go copy(wOut, inConn)

	var wIn io.Writer = inConn
	if s.ResponseDumpWriter != nil {
		w2 := NewDecoderWriter(s.ResponseDumpWriter, decoder.Response)
		defer w2.Close()
		wIn = io.MultiWriter(inConn, w2)
	}
	go copy(wIn, outConn)
	<-errchan
}
