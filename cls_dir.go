package ratcpserver

import (
	"encoding/json"
	"errors"
	cn "github.com/ilinovalex86/connection"
	ex "github.com/ilinovalex86/explorer"
)

// Dir - получает содержимое папки клиента по пути path
func (cls *clients) Dir(id string, path string) (map[string][]ex.LinkAndName, string, error) {
	conn, err := cls.clientAccessBlock(id)
	if err != nil {
		return nil, "", err
	}
	err = cn.SendQuery(cn.Query{Method: "dir", Query: path}, conn)
	if err != nil {
		cls.errConn(id)
		return nil, "", errors.New("компьютер недоступен")
	}
	r, err := cn.ReadResponse(conn)
	if err != nil {
		cls.errConn(id)
		return nil, "", errors.New("компьютер недоступен")
	}
	if r.Err != nil {
		cls.clientAccessFree(id)
		return nil, "", r.Err
	}
	cn.SendSync(conn)
	data, err := cn.ReadBytesByLen(r.DataLen, conn)
	if err != nil {
		cls.errConn(id)
		return nil, "", errors.New("компьютер недоступен")
	}
	defer cls.clientAccessFree(id)
	res := make(map[string][]ex.LinkAndName)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, "", errors.New("ошибка данных")
	}
	return res, cls.sep(id), nil
}
