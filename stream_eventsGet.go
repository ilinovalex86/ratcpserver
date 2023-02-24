package ratcpserver

func (s *stream) eventsGet() []Event {
	s.mu.Lock()
	defer s.mu.Unlock()
	res := s.events
	s.events = []Event{}
	return res
}
