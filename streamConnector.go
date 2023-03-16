package ratcpserver

import (
	"encoding/json"
	"fmt"
	cn "github.com/ilinovalex86/connection"
	"net"
	"time"
)

// Обрабатывает новое подключение
func streamConnector(conn net.Conn) {
	const funcNameLog = "streamConnector(): "
	conn.SetDeadline(time.Now().Add(time.Second * 5))
	cl, err := newClient(conn)
	if err != nil {
		ToLog(tcpServer, funcNameLog+fmt.Sprint(err), false)
		conn.Close()
		return
	}
	conn.SetDeadline(time.Time{})
	ok := Clients.Exist(cl.Id)
	if !ok || ok && cl.BasePath != Clients.basePath(cl.Id) {
		ToLog(tcpServer, funcNameLog+"!ok || ok && cl.BasePath != Clients.basePath(cl.Id)", false)
		conn.Close()
		return
	}
	cn.SendSync(conn)
	jsonSS, err := cn.ReadByteByDelim(conn)
	if err != nil {
		ToLog(tcpServer, funcNameLog+"cn.ReadByteByDelim(conn)", false)
		conn.Close()
		return
	}
	var ss [2]int
	err = json.Unmarshal(jsonSS, &ss)
	if err != nil {
		ToLog(tcpServer, funcNameLog+"json.Unmarshal(jsonSS, &ss)", false)
		conn.Close()
		return
	}
	s := Clients.streamNew(cl, ss)
	go streamWorker(s)
}
