package ratcpserver

func (cls *clients) clientAccessFree(id string) {
	cls.mu.Lock()
	defer cls.mu.Unlock()
	cls.m[id].busy = false
}
