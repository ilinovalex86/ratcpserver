package ratcpserver

func (cls *clients) errConn(id string) {
	const funcNameLog = "errConn(): "
	cls.mu.Lock()
	defer cls.mu.Unlock()
	cls.m[id].conn.Close()
	cls.m[id].status = false
	ToLog(id, funcNameLog+"errConn(): connClose", false)
}
