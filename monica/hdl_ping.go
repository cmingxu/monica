package monica

import (
	"github.com/cmingxu/monica/protogos/common"
	"github.com/golang/protobuf/proto"
	"log"
)

type HdlPing struct {
	session *Session
}

func NewPingHandler(session *Session) *HdlPing {
	return &HdlPing{
		session: session,
	}
}

func (handler *HdlPing) HandlePackage(buf []byte) {
	pong := new(common.Pong)
	bytes, _ := proto.Marshal(pong)
	log.Println("xxxxx", len(bytes))
	handler.session.WriteToClient(ProtoPong, bytes)
}
