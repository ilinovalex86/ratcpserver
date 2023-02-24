package ratcpserver

import (
	"fmt"
	cn "github.com/ilinovalex86/connection"
	"net"
	"time"
)

// Обрабатывает новое подключение
func tcpConnector(conn net.Conn) {
	conn.SetDeadline(time.Now().Add(time.Second * 5))
	cl, err := newClient(conn)
	if err != nil {
		conn.Close()
		return
	}
	conn.SetDeadline(time.Time{})
	if cl.Version != actualVersionClients {
		_ = cn.SendQuery(cn.Query{Method: "wrong version"}, cl.conn)
		conn.Close()
		return
	}
	if cl.Id == "----------------" {
		err = cl.isNewClient()
		if err != nil {
			conn.Close()
			return
		}
	} else {
		err = cl.isOldClient()
		if err != nil {
			conn.Close()
			return
		}
	}
	fmt.Println("connector: ", cl.Id, cl.BasePath)
}
