package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/cmingxu/monica/monica"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var config = new(monica.MonicaConfig)
var configFilePath string
var Server *monica.MonicaServer

func main() {
	SetParseFlag()
	flag.Parse()

	Initialization()
}

func Initialization() {
	InitConfig()
	InitLog()

	Server = new(monica.MonicaServer)
	Server.Init(config).Start()
}

func InitConfig() {
	configFile, err := os.Open(configFilePath)
	if err != nil {
		log.Panic(err)
	}
	defer configFile.Close()

	scanner := bufio.NewScanner(configFile)
	for scanner.Scan() {
		var key, value string
		c := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(c, "#") {
			key = strings.Split(c, "=")[0]
			value = strings.Split(c, "=")[1]

			switch key {
			case "port":
				port, _ := strconv.Atoi(value)
				config.Port = port
			case "host":
				config.Host = value
			case "log_level":
				config.LogLevel, _ = strconv.Atoi(value)
			case "env":
				config.Env = value
			case "mysqlschema":
				config.MysqlSchema = value
			case "redisschema":
				config.RedisSchema = value
			case "log_path":
				if strings.HasPrefix(value, "/") {
					config.LogPath = value
				} else {
					config.LogPath = fmt.Sprintf("%s/%s", getwd(), value)
				}
			default:
			}
		}
	}
	config.GdsPath = fmt.Sprintf("%s/%s", getwd(), "/gds")
}

func InitLog() {
	logfile, err := os.Create(filepath.Join(config.LogPath, "log.log"))
	if err != nil {
		fmt.Println("logger initializing error")
	}
	if config.Env == "dev" {
		config.Log = log.New(os.Stdout, "", log.Lshortfile)
	} else {
		config.Log = log.New(logfile, "", log.Lshortfile)
	}
}

func SetParseFlag() {
	defaultEtcPath := fmt.Sprintf("%s/etc/config", getwd())
	flag.StringVar(&configFilePath, "config", defaultEtcPath, "config path ")
	flag.StringVar(&configFilePath, "c", defaultEtcPath, "config path ")
}

func getwd() (wd string) {
	wd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}
	return wd
}
