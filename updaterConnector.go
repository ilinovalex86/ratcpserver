package ratcpserver

import (
	cn "github.com/ilinovalex86/connection"
	ex "github.com/ilinovalex86/explorer"
	"net"
	"time"
)

// Обрабатывает новое подключение
func updaterConnector(conn net.Conn) {
	conn.SetDeadline(time.Now().Add(time.Second * 5))
	cl, err := newClient(conn)
	if err != nil {
		conn.Close()
		return
	}
	conn.SetDeadline(time.Time{})
	ok := Clients.Exist(cl.Id)
	if ok && Clients.testConnect(cl.Id) {
		_ = cn.SendQuery(cn.Query{Method: "already exist"}, cl.conn)
		conn.Close()
		return
	}
	file := clientsForLinux
	if cl.System == "windows" {
		file = clientsForWindows
	}
	data, err := ex.ReadFileFull(file)
	check(err)
	if cl.Version != actualVersionClients {
		err = cn.SendQuery(cn.Query{Method: "downloadNewClient", Query: file, DataLen: len(data)}, cl.conn)
		if err != nil {
			conn.Close()
			return
		}
		cn.ReadSync(cl.conn)
		err = cn.SendBytes(data, cl.conn)
		conn.Close()
		return
	} else {
		err = cn.SendQuery(cn.Query{Method: "lenClient", DataLen: len(data)}, cl.conn)
		if err != nil {
			conn.Close()
			return
		}
	}
}
