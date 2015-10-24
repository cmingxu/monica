package monica

import (
	"net"
	"time"
)

type Session struct {
	Conn              net.Conn
	BytesTransfered   int32
	PackageTransfered int32
	ConnectedAt       time.Time
	LastPingAt        time.Time
}

func NewSession(c net.Conn) *Session {
	return &Session{Conn: c}
}
