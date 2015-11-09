package models

import (
	//"github.com/cmingxu/monica/monica"
	"time"
)

type Player struct {
	Id          uint32
	Name        string
	Email       string
	LastLoginAt time.Time
	LastLoginIp string
	Gem         uint32
	CreatedAt   time.Time
	UpdateAt    time.Time
}

func NewPlayer() *Player {
	return &Player{}
}

func DbSave(p *Player) *Player {
	return p
}
func DbUpdate(p *Player) *Player {
	return p
}
func DbDestroy(*Player) bool {
	return true
}
