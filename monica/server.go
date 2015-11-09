package monica

import (
	"fmt"
	"log"
	"net"
)

const (
	MAX_CLIENTS_HOLD = 12000 // maximum clinet connection that server suppose to support
)

type ServerState int     // state of server
type Sessions []*Session // client list

const (
	IDLE    ServerState = iota // idle
	RUNNING                    // running
	STOP                       // stop
)

type MonicaServer struct {
	Config         *MonicaConfig // game server config
	State          ServerState   // current state of the game server
	ClientSessions Sessions      // list of client connections
	MySQL          *MonicaMySQL
	Redis          *MonicaRedis
}

// Game server initialization
func (s *MonicaServer) Init(config *MonicaConfig) *MonicaServer {
	s.State = IDLE
	s.Config = config
	s.ClientSessions = make(Sessions, MAX_CLIENTS_HOLD)

	return s
}

// 1 Game server listening
// 2, Game server config loading
// 3, Log initialization
// 4, Tcp start listening
// 5, Signal handler
// 6, MySQL
// 7, Redis
func (s *MonicaServer) Start() *MonicaServer {
	s.Config.Log.Printf("server starting at %s:%d\n ", s.Config.Host, s.Config.Port)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port))
	if err != nil {
		s.Config.Log.Printf("server listening error")
		log.Panic(err)
	}

	// mysql connection here
	s.MySQL = NewMonicaMySQL(s)

	// redis connection local server
	s.Redis = NewRedisClient(s)

	for {
		// accpeting new connections
		clientConn, err := listener.Accept()
		if err != nil {
			s.Config.Log.Printf("accepting error")
			log.Panic(err)
		}

		// package tcp connection into sessions and added to servers client's list
		clientSession := NewSession(clientConn)
		clientSession.Server = s
		s.ClientSessions = append(s.ClientSessions, clientSession)

		// handle client conenctions here
		go handleClientConn(clientSession)
	}

	return s
}

func handleClientConn(session *Session) {
	session.Loop()
}
