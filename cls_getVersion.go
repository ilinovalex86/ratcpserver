package ratcpserver

// Добавляет\обновляет информацию о клиенте
func (cls *clients) getVersion(id string) string {
	cls.mu.Lock()
	defer cls.mu.Unlock()
	return cls.m[id].Version
}
