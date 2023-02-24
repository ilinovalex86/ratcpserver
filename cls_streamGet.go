package ratcpserver

import (
	"errors"
)

// StreamGet - возвращает указатель на стрим
func (cls *clients) StreamGet(id string) (*stream, error) {
	cls.mu.RLock()
	defer cls.mu.RUnlock()
	if !cls.m[id].streamStatus {
		return nil, errors.New("no stream")
	}
	return cls.m[id].streamLink, nil
}
