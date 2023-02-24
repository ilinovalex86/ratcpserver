package ratcpserver

func (s *stream) webClientExist(id string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, v := range s.webClients {
		if v == id {
			return true
		}
	}
	return false
}
