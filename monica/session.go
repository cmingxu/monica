package monica

import (
	"bytes"
	"encoding/binary"
	"github.com/cmingxu/monica/cmd_handler"
	"github.com/cmingxu/monica/protogos/common"
	"github.com/golang/protobuf/proto"
	"io"
	"log"
	"net"
	"time"
)

type Session struct {
	Conn              net.Conn
	BytesTransfered   int32
	PackageTransfered int32
	ConnectedAt       time.Time
	LastPingAt        time.Time
	PingInterval      time.Duration
}

func NewSession(c net.Conn) *Session {
	return &Session{Conn: c,
		ConnectedAt:       time.Now(),
		LastPingAt:        time.Now(),
		BytesTransfered:   0,
		PackageTransfered: 0,
		PingInterval:      1 * time.Second,
	}
}

func (s *Session) Loop() {
	buf := make([]byte, 4096)
	log.Printf("start reading from %s\n", s.Conn.RemoteAddr().String())
	go s.AsyncLoop()
	for {
		ByteRead, err := s.Conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Println("client exiting")
			} else {
				log.Fatalf("client reading error")
			}

			break
		}
		var size uint32
		var protoType ProtoCode

		bufReader := bytes.NewReader(buf[:4])
		binary.Read(bufReader, binary.LittleEndian, &size)

		abufReader := bytes.NewReader(buf[4:8])
		binary.Read(abufReader, binary.LittleEndian, &protoType)

		switch protoType {
		case ProtoPing:
		case GdsSync:
			go cmd_handler.NewHandler(s).HandlePackage(buf[8:ByteRead])
		}
	}
}

func (s *Session) WriteToClient(protoCode ProtoCode, content []byte) {
	protoPackageSizeBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(protoPackageSizeBytes, uint32(4+4+len(content)))

	protoCodeBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(protoCodeBytes, uint32(protoCode))

	buffer := bytes.NewBufferString("")
	buffer.Write(protoPackageSizeBytes)
	buffer.Write(protoCodeBytes)
	buffer.Write(content)

	s.Conn.Write(buffer.Bytes())
}

func (s *Session) AsyncLoop() {
	ticker := time.NewTicker(s.PingInterval)
	for {
		select {
		case <-ticker.C:
			ping := &common.Ping{
				Timestamp: proto.Int64(1),
				Header:    &common.Header{Code: proto.Int32(2)},
			}
			bytes, _ := proto.Marshal(ping)
			s.WriteToClient(ProtoPing, bytes)
		}
	}
}
