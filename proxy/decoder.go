package proxy

import (
	"fmt"
	"io"
	"net"
	"reflect"
)

type NewDecoderFunc func(conn net.Conn) Decoder

type Decoder interface {
	Request(in io.Reader, out io.Writer)

	Response(in io.Reader, out io.Writer)
}

func NewNopDecoderFunc(_ net.Conn) Decoder {
	return &NopDecoder{}
}

type NopDecoder struct {
}

func (n NopDecoder) Request(in io.Reader, out io.Writer) {
	io.Copy(out, in)
}

func (n NopDecoder) Response(in io.Reader, out io.Writer) {
	io.Copy(out, in)
}

func newDecoder(rv reflect.Value) NewDecoderFunc {
	return func(conn net.Conn) Decoder {
		vv := rv.Call([]reflect.Value{reflect.ValueOf(conn)})
		if len(vv) != 1 {
			panic(fmt.Sprintf("newDecoder with wrong result length, got=%d, want=1", len(vv)))
		}
		decoder, ok := vv[0].Interface().(Decoder)
		if !ok {
			panic("not Decoder")
		}
		return decoder
	}
}

type writer struct {
	r *io.PipeReader
	w *io.PipeWriter
}

func (w writer) Write(p []byte) (n int, err error) {
	return w.w.Write(p)
}

func (w writer) Close() error {
	w.r.CloseWithError(io.EOF)
	w.w.CloseWithError(io.EOF)
	return nil
}

func NewDecoderWriter(out io.WriteCloser, decoder func(in io.Reader, out io.Writer)) io.WriteCloser {
	w := &writer{}
	w.r, w.w = io.Pipe()
	go decoder(w.r, out)
	return w
}
