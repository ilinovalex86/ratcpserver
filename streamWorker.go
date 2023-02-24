package ratcpserver

import (
	"encoding/json"
	"fmt"
	cn "github.com/ilinovalex86/connection"
	"time"
)

//Горутина работы со стримом клиента
func streamWorker(s *stream) {
	var events []Event
	for {
		t := time.Now()
		if s.webClientsCount() < 1 {
			s.stop()
			return
		}
		if len(events) > 0 {
			jsonEvents, err := json.Marshal(events)
			if err != nil {
				s.stop()
				fmt.Println("jsonMarshal")
				return
			}
			err = cn.SendBytesWithDelim(jsonEvents, s.conn)
			if err != nil {
				s.stop()
				return
			}
		} else {
			cn.SendSync(s.conn)
		}
		r, err := cn.ReadResponse(s.conn)
		if err != nil {
			s.stop()
			return
		}
		events = s.eventsGet()
		if len(events) > 0 {
			err = cn.SendString("yes", s.conn)
		} else {
			err = cn.SendString("no", s.conn)
		}
		data, err := cn.ReadBytesByLen(r.DataLen, s.conn)
		if err != nil {
			s.stop()
			return
		}
		s.imgUpdate(data)
		fmt.Printf("DataLen: %d kb, time: %d \n", r.DataLen/1000, time.Since(t).Milliseconds())
	}
}
