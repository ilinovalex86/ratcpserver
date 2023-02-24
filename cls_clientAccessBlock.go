package ratcpserver

import (
	"errors"
	"net"
	"time"
)

func (cls *clients) clientAccessBlock(id string) (net.Conn, error) {
	t := time.Now()
	for {
		if time.Since(t).Seconds() > 30 {
			return nil, errors.New("компьютер занят")
		}
		cls.mu.Lock()
		if !cls.m[id].status {
			cls.mu.Unlock()
			return nil, errors.New("компьютер недоступен")
		}
		if cls.m[id].busy {
			cls.mu.Unlock()
			time.Sleep(500 * time.Millisecond)
			continue
		}
		cls.m[id].busy = true
		conn := cls.m[id].conn
		cls.mu.Unlock()
		return conn, nil
	}
}
