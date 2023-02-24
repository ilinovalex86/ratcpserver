package ratcpserver

import (
	"encoding/json"
	cn "github.com/ilinovalex86/connection"
	"net"
	"time"
)

// Обрабатывает новое подключение
func streamConnector(conn net.Conn) {
	conn.SetDeadline(time.Now().Add(time.Second * 5))
	cl, err := newClient(conn)
	if err != nil {
		conn.Close()
		return
	}
	conn.SetDeadline(time.Time{})
	ok := Clients.Exist(cl.Id)
	if !ok || ok && cl.BasePath != Clients.basePath(cl.Id) {
		conn.Close()
		return
	}
	cn.SendSync(conn)
	jsonSS, err := cn.ReadByteByDelim(conn)
	if err != nil {
		conn.Close()
		return
	}
	var ss [2]int
	err = json.Unmarshal(jsonSS, &ss)
	if err != nil {
		conn.Close()
		return
	}
	s := Clients.streamNew(cl, ss)
	go streamWorker(s)
}
