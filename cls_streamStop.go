package ratcpserver

func (cls *clients) streamStop(id string) {
	cls.mu.Lock()
	defer cls.mu.Unlock()
	cls.m[id].streamStatus = false
	cls.m[id].streamRunUser = ""
	ToLog(id, "streamStop", false)
}
