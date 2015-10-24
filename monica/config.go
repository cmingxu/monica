package monica

import (
	"fmt"
	"log"
)

type MonicaConfig struct {
	Port int
	Host string
	Env  string

	Log      *log.Logger
	LogLevel int
	LogPath  string
}

func (s *MonicaConfig) ToString() string {
	return fmt.Sprintf("MonicConfig port %d Host %s LogLevel %s LogPath %s", s.Port, s.Host, s.LogLevel, s.LogPath)
}
