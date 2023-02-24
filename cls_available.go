package ratcpserver

// Available - возвращает список доступных клиентов
func (cls *clients) Available() []string {
	var res []string
	cls.mu.RLock()
	defer cls.mu.RUnlock()
	for id, cl := range cls.m {
		if cl.status {
			res = append(res, id)
		}
	}
	return res
}
