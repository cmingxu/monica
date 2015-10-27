package main

import (
	"fmt"
	//"github.com/cmingxu/monica/monica/proto/common"
	//"github.com/golang/protobuf/proto"
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
		Port: 8080,
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
	return c
}

func main() {
	log.Println("cliennt starting at ")
	client := NewClient()
	client.Init()
	client.BeginLoop()
}
