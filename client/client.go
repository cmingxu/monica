package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/cmingxu/monica/monica"
	"github.com/cmingxu/monica/protogos/common"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
)

type Client struct {
	Conn net.Conn
	Host string
	Port int
}

func NewClient() *Client {
	return &Client{
		Host: "localhost",
		Port: 8812,
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

		pingProto := new(common.Ping)
		proto.Unmarshal(buf[8:ByteRead], pingProto)
		fmt.Println("size", size)
		fmt.Println("protoType", protoType)
		fmt.Println("timeStamp", pingProto.GetTimestamp())
	}
	return c
}

func main() {
	log.Println("cliennt starting at ")
	client := NewClient()
	client.Init()
	client.BeginLoop()
}
