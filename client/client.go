package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/cmingxu/monica/monica"
	"github.com/cmingxu/monica/protogos/common"
	"github.com/cmingxu/monica/protogos/gds"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
)

type Client struct {
	Conn net.Conn
	Host string
	Port int
	Stop chan bool
}

func NewClient() *Client {
	return &Client{
		Host: "localhost",
		Port: 8812,
		Stop: make(chan bool),
	}
}

func (c *Client) EntryPoint() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c *Client) Init() {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.Host, c.Port))
	if err != nil {
		log.Fatal("connection error")
	}
	c.Conn = conn
}

func (c *Client) BeginLoop() *Client {
	buf := make([]byte, 4096)
	for {
		ByteRead, err := c.Conn.Read(buf)
		fmt.Println("ByteRead", ByteRead)
		if err != nil {
			log.Fatal("xx")
		}

		var size uint32
		var protoType monica.ProtoCode

		bufReader := bytes.NewReader(buf[:4])
		binary.Read(bufReader, binary.LittleEndian, &size)

		abufReader := bytes.NewReader(buf[4:8])
		binary.Read(abufReader, binary.LittleEndian, &protoType)

		switch protoType {
		case monica.ProtoPing:
			pingProto := new(common.Ping)
			proto.Unmarshal(buf[8:ByteRead], pingProto)
			fmt.Println("size", size)
			fmt.Println("protoType", protoType)
			fmt.Println("timeStamp", pingProto.GetTimestamp())
		case monica.ProtoBuildingsGds:
			gds := new(gds.BuildingGds)
			proto.Unmarshal(buf[8:ByteRead], gds)
			log.Println(len(gds.GetBuildings()))
			for _, building := range gds.GetBuildings() {
				fmt.Println("name ", building.GetName())
				fmt.Println("level ", building.GetLevel())
			}
		}

	}
	return c
}

func (c *Client) SendGdsSync() *Client {
	version := "xxxxx"
	what := "what"
	gdsSync := &common.GdsSync{
		Version: &version,
		What:    &what}
	bytes, _ := proto.Marshal(gdsSync)
	c.WriteToClient(monica.ProtoGdsSync, bytes)
	return c
}

func (s *Client) WriteToClient(protoCode monica.ProtoCode, content []byte) {
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

func main() {
	log.Println("cliennt starting at ")
	client := NewClient()
	client.Init()
	go client.BeginLoop()
	client.SendGdsSync()
	<-client.Stop
}
