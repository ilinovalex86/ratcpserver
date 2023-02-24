package ratcpserver

import (
	"encoding/json"
	"io/ioutil"
)

// Сохраняет информацию о клиентах в json файл
func (cls *clients) dump() error {
	cls.mu.RLock()
	jsonData, err := json.MarshalIndent(&cls.m, "", "  ")
	cls.mu.RUnlock()
	if err != nil {
		return err
	}
	clientsDB.Lock()
	defer clientsDB.Unlock()
	err = ioutil.WriteFile(clientsDB.file, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}
