package ratcpserver

import (
	"fmt"
	"time"
)

func (s *stream) ImgNext(i int) (int, []byte, error) {
	for {
		s.mu.RLock()
		if s.err != nil {
			s.mu.RUnlock()
			ToLog(s.id, "ImgNext(): s.err: "+fmt.Sprint(s.err), false)
			return 0, nil, s.err
		}
		if s.imgIndex <= i {
			s.mu.RUnlock()
			time.Sleep(5 * time.Millisecond)
			continue
		}
		s.mu.RUnlock()
		return s.imgIndex, s.imgData, nil
	}
}
