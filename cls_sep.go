package ratcpserver

func (cls *clients) sep(id string) string {
	cls.mu.RLock()
	defer cls.mu.RUnlock()
	return cls.m[id].Sep
}
