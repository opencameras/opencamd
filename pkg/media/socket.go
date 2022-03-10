package media

import (
	"fmt"
	"net"
	"context"
	"github.com/pion/webrtc/v3/pkg/media/h264reader"
	"github.com/pion/webrtc/v3/pkg/media/oggreader"
)

var startCodeBuf = make ([]byte, 1)

type TCPSocketServer struct {
	ctx context.Context
	streaming bool
	aligned bool
	fd net.Conn
	addr string
	notify chan struct{}
}

func NewTcpSocketServer(ctx context.Context, addr string) (*TCPSocketServer, error) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	nt := &TCPSocketServer{addr: addr, notify: make(chan struct{})}
	go func(ln net.Listener, ctx context.Context) {
		select {
			case <- ctx.Done():
				fmt.Printf("media server is closed: %s\n", ctx.Err())
				ln.Close()
		}
	}(ln, ctx)

	go func(ln net.Listener, nt *TCPSocketServer) {
		var err error
		nt.fd, err = ln.Accept()
		if err != nil {
			fmt.Printf("Accept error: %s\n", err)
			return
		}

		go func (nt *TCPSocketServer) {
			buf := make([]byte, 512)
			for {
				if nt.streaming {
					// wait until streaming is closed
					<- nt.notify
				}
				_, err := nt.fd.Read(buf)
				if err != nil {
					fmt.Printf("Read error: %s\n", err)
					return
				}
			}
		} (nt)
	} (ln, nt)

	return nt, nil
}

func (ts *TCPSocketServer) GetCodecs() ([]string, error) {
	return []string{"h264"}, nil
}

func (ts *TCPSocketServer) OpenH264() (*h264reader.H264Reader, error) {
	ts.streaming = true
	ts.aligned = false
	return h264reader.NewReader(ts)
}

func (ts *TCPSocketServer) OpenOgg() (*oggreader.OggReader, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (ts *TCPSocketServer) Close() error {
	ts.streaming = false
	ts.aligned = false
	ts.notify <- struct{}{}
	return nil
}

func (ts *TCPSocketServer) String() string {
	return "TcpSocketServer"
}

func (ts *TCPSocketServer) Read(p []byte) (n int, err error){
	if ts.aligned {
		return ts.fd.Read(p)
	}

	// read buffer should be larger than 5: 32 bit start code + 1 byte (header)
	if len(p) <= 5 {
		return 0, fmt.Errorf("read buffer length should be larger than 5")
	}
	countOfConsecutiveZeroBytes := 0
	for {
		// read byte one by one
		n, err = ts.fd.Read(startCodeBuf)
		if err != nil {
			return
		}
		switch startCodeBuf[0] {
		case 0:
			countOfConsecutiveZeroBytes++
		case 1:
			if countOfConsecutiveZeroBytes < 2 {
				countOfConsecutiveZeroBytes = 0
				continue
			}
			ts.aligned = true
			// read one more byte, fill in buffer
			n, err = ts.fd.Read(startCodeBuf)
			if err != nil {
				return
			}
			p[0] = 0
			p[1] = 0
			p[2] = 0
			p[3] = 1
			p[4] = startCodeBuf[0]
			n = 5
			return
		default:
			countOfConsecutiveZeroBytes = 0
		}
	}
}