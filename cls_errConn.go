package ratcpserver

func (cls *clients) errConn(id string) {
	cls.mu.Lock()
	defer cls.mu.Unlock()
	cls.m[id].conn.Close()
	cls.m[id].status = false
}
