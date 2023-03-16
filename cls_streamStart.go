package ratcpserver

import (
	"errors"
	"fmt"
	cn "github.com/ilinovalex86/connection"
	"net"
	"time"
)

// StreamStart - показывает рабочий стол клиента клиента
func (cls *clients) StreamStart(id string, webId string) (*stream, error) {
	const funcNameLog = "cls.StreamStart(): "
	var conn net.Conn
	for {
		cls.mu.Lock()
		if !cls.m[id].status {
			cls.mu.Unlock()
			ToLog(id, funcNameLog+"компьютер недоступен", false)
			return nil, errors.New("компьютер недоступен")
		}
		if cls.m[id].streamStatus == true {
			s := cls.m[id].streamLink
			if s.webClientExist(webId) {
				cls.mu.Unlock()
				ToLog(id, funcNameLog+"вы уже подключены", false)
				return nil, errors.New("вы уже подключены")
			}
			if s.webClientsCount() >= 2 {
				cls.mu.Unlock()
				ToLog(id, funcNameLog+"превышен лимит подключений", false)
				return nil, errors.New("превышен лимит подключений")
			}
			s.webClientAdd(webId)
			cls.mu.Unlock()
			ToLog(id, funcNameLog+"подключен к существующему стриму", false)
			return s, nil
		}
		if cls.m[id].busy {
			cls.mu.Unlock()
			time.Sleep(500 * time.Millisecond)
			continue
		}
		cls.m[id].busy = true
		cls.m[id].streamRunUser = webId
		conn = cls.m[id].conn
		cls.mu.Unlock()
		break
	}
	err := cn.SendQuery(cn.Query{Method: "stream"}, conn)
	if err != nil {
		cls.errConn(id)
		ToLog(id, funcNameLog+"cn.SendQuery(cn.Query{Method: \"stream\"}, conn)", false)
		return nil, errors.New("компьютер недоступен")
	}
	r, err := cn.ReadResponse(conn)
	if err != nil {
		cls.errConn(id)
		ToLog(id, funcNameLog+"cn.ReadResponse(conn)", false)
		return nil, errors.New("компьютер недоступен")
	}
	if r.Err != nil {
		ToLog(id, funcNameLog+"r.Err: "+fmt.Sprint(r.Err), false)
		cls.clientAccessFree(id)
		return nil, r.Err
	}
	time.Sleep(2 * time.Second)
	cls.mu.Lock()
	defer cls.mu.Unlock()
	cls.m[id].busy = false
	s := cls.m[id].streamLink
	ToLog(id, funcNameLog+"done", false)
	return s, nil
}
