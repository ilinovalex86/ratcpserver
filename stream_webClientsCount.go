package ratcpserver

func (s *stream) webClientsCount() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.webClients)
}
