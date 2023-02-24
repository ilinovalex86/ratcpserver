package ratcpserver

func (s *stream) webClientAdd(webId string) {
	s.mu.Lock()
	s.webClients = append(s.webClients, webId)
	s.mu.Unlock()
}
