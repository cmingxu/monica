package cmd_handler

import (
	"github.com/cmingxu/monica/monica"
	"github.com/cmingxu/monica/protogos/common"
	"github.com/golang/protobuf/proto"
)

type GdsHandler struct {
	session *monica.Session
}

func NewHandler(session *monica.Session) *GdsHandler {
	return &GdsHandler{
		session: session,
	}
}

func (*GdsHandler) HandlePackage(buf []byte) {
	gdsSync := new(common.GdsSync)
	proto.Unmarshal(buf, gdsSync)
	fmt.Println(gdsSync.version)
	fmt.Println(gdsSync.what)
}
