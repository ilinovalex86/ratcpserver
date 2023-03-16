package ratcpserver

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// Сохраняет информацию о клиентах в json файл
func (cls *clients) dump() error {
	const funcNameLog = "cls.dump(): "
	cls.mu.RLock()
	jsonData, err := json.MarshalIndent(&cls.m, "", "  ")
	cls.mu.RUnlock()
	if err != nil {
		return errors.New(funcNameLog + "json.MarshalIndent(&cls.m, \"\", \"  \")")
	}
	clientsDB.Lock()
	defer clientsDB.Unlock()
	err = ioutil.WriteFile(clientsDB.file, jsonData, 0644)
	if err != nil {
		return errors.New(funcNameLog + "ioutil.WriteFile(clientsDB.file, jsonData, 0644)")
	}
	return nil
}
