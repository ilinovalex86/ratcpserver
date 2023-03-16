package ratcpserver

import (
	"fmt"
	cn "github.com/ilinovalex86/connection"
	ex "github.com/ilinovalex86/explorer"
	"net"
	"time"
)

// Обрабатывает новое подключение
func updaterConnector(conn net.Conn) {
	const funcNameLog = "updaterConnector(): "
	conn.SetDeadline(time.Now().Add(time.Second * 5))
	cl, err := newClient(conn)
	if err != nil {
		ToLog(tcpServer, funcNameLog+fmt.Sprint(err), false)
		conn.Close()
		return
	}
	conn.SetDeadline(time.Time{})
	ok := Clients.Exist(cl.Id)
	if ok && Clients.testConnect(cl.Id) {
		_ = cn.SendQuery(cn.Query{Method: "already exist"}, cl.conn)
		ToLog(tcpServer, funcNameLog+"already exist", false)
		conn.Close()
		return
	}
	file := clientsForLinux
	if cl.System == "windows" {
		file = clientsForWindows
	}
	data, err := ex.ReadFileFull(file)
	if err != nil {
		ToLog(tcpServer, funcNameLog+"ex.ReadFileFull(file)", false)
	}
	if cl.Version != actualVersionClients {
		err = cn.SendQuery(cn.Query{Method: "downloadNewClient", Query: file, DataLen: len(data)}, cl.conn)
		if err != nil {
			ToLog(tcpServer, funcNameLog+"cn.SendQuery(cn.Query{Method: \"downloadNewClient\", Query: file, DataLen: len(data)}, cl.conn)", false)
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
			ToLog(tcpServer, funcNameLog+"cn.SendQuery(cn.Query{Method: \"lenClient\", DataLen: len(data)}, cl.conn)", false)
			conn.Close()
			return
		}
	}
}
