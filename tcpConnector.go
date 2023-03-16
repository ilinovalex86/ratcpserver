package ratcpserver

import (
	"fmt"
	cn "github.com/ilinovalex86/connection"
	"net"
	"time"
)

// Обрабатывает новое подключение
func tcpConnector(conn net.Conn) {
	const funcNameLog = "tcpConnector(): "
	conn.SetDeadline(time.Now().Add(time.Second * 5))
	cl, err := newClient(conn)
	if err != nil {
		ToLog(tcpServer, funcNameLog+fmt.Sprint(err), false)
		conn.Close()
		return
	}
	conn.SetDeadline(time.Time{})
	if cl.Version != actualVersionClients {
		_ = cn.SendQuery(cn.Query{Method: "wrong version"}, cl.conn)
		conn.Close()
		ToLog(tcpServer, funcNameLog+"wrong version", false)
		return
	}
	if cl.Id == "----------------" {
		err = cl.isNewClient()
		if err != nil {
			ToLog(tcpServer, funcNameLog+fmt.Sprint(err), false)
			conn.Close()
			return
		}
	} else {
		err = cl.isOldClient()
		if err != nil {
			ToLog(tcpServer, funcNameLog+fmt.Sprint(err), false)
			conn.Close()
			return
		}
	}
	ToLog(tcpServer, fmt.Sprintf("%s %s %s", funcNameLog, cl.Id, cl.BasePath), false)
	fmt.Println("connector: ", cl.Id, cl.BasePath)
}
