package ratcpserver

// basePath - Возвращает basePath клиента
func (cls *clients) basePath(id string) string {
	cls.mu.RLock()
	defer cls.mu.RUnlock()
	return cls.m[id].BasePath
}
