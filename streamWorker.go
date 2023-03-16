package ratcpserver

import (
	"encoding/json"
	"fmt"
	cn "github.com/ilinovalex86/connection"
	"time"
)

//Горутина работы со стримом клиента
func streamWorker(s *stream) {
	const funcNameLog = "streamWorker(): "
	id := s.id
	var events []Event
	for {
		t := time.Now()
		if s.webClientsCount() < 1 {
			s.stop()
			ToLog(id, "s.webClientsCount() < 1", false)
			return
		}
		if len(events) > 0 {
			jsonEvents, err := json.Marshal(events)
			if err != nil {
				s.stop()
				ToLog(id, "json.Marshal(events)", false)
				return
			}
			err = cn.SendBytesWithDelim(jsonEvents, s.conn)
			if err != nil {
				s.stop()
				ToLog(id, "cn.SendBytesWithDelim(jsonEvents, s.conn)", false)
				return
			}
		} else {
			cn.SendSync(s.conn)
		}
		r, err := cn.ReadResponse(s.conn)
		if err != nil {
			ToLog(id, "cn.ReadResponse(s.conn)", false)
			s.stop()
			return
		}
		events = s.eventsGet()
		if len(events) > 0 {
			err = cn.SendString("yes", s.conn)
		} else {
			err = cn.SendString("no", s.conn)
		}
		if err != nil {
			ToLog(id, "cn.SendString(yes/no, s.conn)", false)
			s.stop()
			return
		}
		data, err := cn.ReadBytesByLen(r.DataLen, s.conn)
		if err != nil {
			ToLog(id, "cn.ReadBytesByLen(r.DataLen, s.conn)", false)
			s.stop()
			return
		}
		s.imgUpdate(data)
		fmt.Printf("DataLen: %d kb, time: %d \n", r.DataLen/1000, time.Since(t).Milliseconds())
	}
}
