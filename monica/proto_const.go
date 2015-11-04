package monica

type ProtoCode uint32

const (
	ProtoPing ProtoCode = iota
	ProtoPong
	ProtoHeartBeat
	ProtoHeartBeatRespons
	ProtoGdsSync
	ProtoBuildingsGds
)
