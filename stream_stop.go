package ratcpserver

import "errors"

func (s *stream) stop() {
	s.mu.Lock()
	s.conn.Close()
	s.err = errors.New("stop")
	ToLog(s.id, "(s *stream) stop()", false)
	s.mu.Unlock()
	Clients.streamStop(s.id)
}
