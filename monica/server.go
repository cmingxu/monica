package monica

import (
	"fmt"
	"io"
	"log"
	"net"
)

type ServerState int

const (
	IDLE ServerState = iota
	RUNNING
	STOP
)

type MonicaServer struct {
	Config *MonicaConfig
	State  ServerState
}

func (s *MonicaServer) Init(config *MonicaConfig) *MonicaServer {
	s.State = IDLE
	s.Config = config

	return s
}

func (s *MonicaServer) Start() *MonicaServer {
	s.Config.Log.Printf("server starting at %s:%d\n ", s.Config.Host, s.Config.Port)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port))
	if err != nil {
		s.Config.Log.Printf("server listening error")
		log.Panic(err)
	}

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			s.Config.Log.Printf("accepting error")
			log.Panic(err)
		}

		go handleClientConn(NewSession(clientConn))
	}
	return s
}

func handleClientConn(session *Session) {
	buf := make([]byte, 4096)
	log.Printf("start reading from %s\n", session.Conn.RemoteAddr().String())

	for {
		ByteRead, err := session.Conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Println("client exiting")
			} else {
				log.Println("client reading error")
			}

			break
		}

		fmt.Printf("%d\n", ByteRead)
	}
}
