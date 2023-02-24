package ratcpserver

import (
	cn "github.com/ilinovalex86/connection"
	"net"
	"time"
)

//Проверяет подключение
func (cls *clients) testConnect(id string) bool {
	var conn net.Conn
	for {
		cls.mu.Lock()
		if !cls.m[id].status {
			cls.mu.Unlock()
			return false
		}
		if cls.m[id].busy {
			cls.mu.Unlock()
			time.Sleep(500 * time.Millisecond)
			continue
		}
		cls.m[id].busy = true
		conn = cls.m[id].conn
		cls.mu.Unlock()
		break
	}
	err := cn.SendQuery(cn.Query{Method: "testConnect"}, conn)
	if err != nil {
		cls.errConn(id)
		return false
	}
	_, err = cn.ReadResponse(conn)
	if err != nil {
		cls.errConn(id)
		return false
	}
	cls.clientAccessFree(id)
	return true
}
