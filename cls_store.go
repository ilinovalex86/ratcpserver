package ratcpserver

// Добавляет\обновляет информацию о клиенте
func (cls *clients) store(cl *client) {
	cls.mu.Lock()
	defer cls.mu.Unlock()
	cls.m[cl.Id] = cl
}
