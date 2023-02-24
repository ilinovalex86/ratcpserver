package ratcpserver

import (
	"math/rand"
	"strconv"
	"time"
)

//Генерирует уникальный id для клиентов
func generatorId() string {
	rand.Seed(time.Now().UnixNano())
	for {
		id := strconv.Itoa(rand.Int())
		if len(id) < 16 {
			continue
		}
		id = id[:16]
		if Clients.Exist(id) {
			continue
		} else {
			return id
		}
	}
}
