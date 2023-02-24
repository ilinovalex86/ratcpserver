package ratcpserver

import (
	ex "github.com/ilinovalex86/explorer"
	"math/rand"
	"strconv"
	"time"
)

//Генерирует уникальный id для файлов клиентов
func tempFileName() string {
	rand.Seed(time.Now().UnixNano())
	for {
		tFN := strconv.Itoa(rand.Int())
		if ex.ExistFile(uploadDir + "/" + tFN) {
			continue
		} else {
			return uploadDir + "/" + tFN
		}
	}
}
