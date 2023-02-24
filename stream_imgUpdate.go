package ratcpserver

func (s *stream) imgUpdate(data []byte) {
	s.mu.Lock()
	s.imgData = data
	s.imgIndex++
	s.mu.Unlock()
}
