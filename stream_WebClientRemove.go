package ratcpserver

func (s *stream) WebClientRemove(id string) {
	s.mu.Lock()
	for i, v := range s.webClients {
		if v == id {
			s.webClients = append(s.webClients[:i], s.webClients[i+1:]...)
		}
	}
	ToLog(s.id, "WebClientRemove(): "+id, false)
	s.mu.Unlock()
}
