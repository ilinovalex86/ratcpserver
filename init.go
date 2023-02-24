package ratcpserver

import (
	"encoding/json"
	"fmt"
	ex "github.com/ilinovalex86/explorer"
	"io/ioutil"
	"log"
	"net"
	"sync"
)

const configFile = "conf"
const keysFile = "keys"
const uploadDir = "files"
const clientsForLinux = "newClient"
const clientsForWindows = "newClient.exe"

var actualVersionClients string
var keys map[string]string
var Conf config

type config struct {
	UpdaterServer    string
	TcpServer        string
	StreamServer     string
	WebServerListner string
	WebServer        string
}

// Файл хранения клиентов
var clientsDB = db{file: "tcpClients.json"}

type db struct {
	sync.Mutex
	file string
}

var Clients = &clients{m: make(map[string]*client)}

type clients struct {
	mu sync.RWMutex
	m  map[string]*client
}

type client struct {
	BasePath      string
	Sep           string
	Id            string
	Version       string
	System        string
	status        bool
	busy          bool
	streamStatus  bool
	streamLink    *stream
	streamRunUser string
	conn          net.Conn
}

type stream struct {
	mu          sync.RWMutex
	id          string
	webClients  []string
	imgData     []byte
	imgIndex    int
	events      []Event
	ScreenSizeX int
	ScreenSizeY int
	conn        net.Conn
	err         error
}

type Event struct {
	Method string
	Event  string
	Key    string
	Code   string
	CorX   int
	CorY   int
	Ctrl   bool
	Shift  bool
}

//if err != nil -> log.Fatal(err)
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	if !ex.ExistFile(keysFile) {
		keys := map[string]string{
			"actualVersionClients": "0.0.0",
			"updaterKey":           "0000000000000000",
			"0.0.0":                "0000000000000000",
		}
		data, err := json.MarshalIndent(&keys, "", "  ")
		check(err)
		err = ioutil.WriteFile(keysFile, data, 0644)
		check(err)
		log.Fatal("Файл конфигурации не найден. Создан новый файл конфигурации.")
	} else {
		data, err := ex.ReadFileFull(keysFile)
		check(err)
		err = json.Unmarshal(data, &keys)
		check(err)
		var errB bool
		if actualVersionClients, errB = keys["actualVersionClients"]; !errB {
			log.Fatal("actualVersionClients does not exist in keys file")
		}
		delete(keys, "actualVersionClients")
		fmt.Println("ActualVersionClients: ", actualVersionClients)
	}
	if !ex.ExistFile(clientsForLinux) {
		log.Fatal("clientsForLinux err")
	}
	if !ex.ExistFile(clientsForWindows) {
		log.Fatal("clientsForWindows err")
	}
	if ex.ExistFile(configFile) {
		data, err := ex.ReadFileFull(configFile)
		check(err)
		err = json.Unmarshal(data, &Conf)
		check(err)
	} else {
		conf := config{
			UpdaterServer:    "127.0.0.1:50000",
			TcpServer:        "127.0.0.1:50001",
			StreamServer:     "127.0.0.1:50002",
			WebServerListner: "127.0.0.1:8080",
			WebServer:        "127.0.0.1:8080",
		}
		data, err := json.MarshalIndent(&conf, "", "  ")
		check(err)
		err = ioutil.WriteFile(configFile, data, 0644)
		check(err)
		log.Fatal("Файл конфигурации не найден. Создан новый файл конфигурации.")
	}
	if ex.ExistFile(clientsDB.file) {
		data, err := ex.ReadFileFull(clientsDB.file)
		check(err)
		err = json.Unmarshal(data, &Clients.m)
		check(err)
	}
	if ex.ExistDir(uploadDir) {
		err := ex.ClearDir(uploadDir)
		check(err)
	} else {
		err := ex.MakeDir(uploadDir)
		check(err)
	}
	go server(Conf.UpdaterServer, updaterConnector, "updaterServer")
	go server(Conf.TcpServer, tcpConnector, "tcpServer")
	go server(Conf.StreamServer, streamConnector, "streamServer")
}
