package ratcpserver

func (s *stream) EventAdd(e Event) {
	s.mu.Lock()
	s.events = append(s.events, e)
	s.mu.Unlock()
}
