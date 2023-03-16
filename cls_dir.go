package ratcpserver

import (
	"encoding/json"
	"errors"
	"fmt"
	cn "github.com/ilinovalex86/connection"
	ex "github.com/ilinovalex86/explorer"
)

// Dir - получает содержимое папки клиента по пути path
func (cls *clients) Dir(id string, path string) (map[string][]ex.LinkAndName, string, error) {
	const funcNameLog = "cls.Dir(): "
	conn, err := cls.clientAccessBlock(id)
	if err != nil {
		ToLog(id, funcNameLog+"cls.clientAccessBlock(id): "+fmt.Sprint(err), false)
		return nil, "", err
	}
	err = cn.SendQuery(cn.Query{Method: "dir", Query: path}, conn)
	if err != nil {
		cls.errConn(id)
		ToLog(id, funcNameLog+"cn.SendQuery(cn.Query{Method: \"dir\", Query: path}, conn)", false)
		return nil, "", errors.New("компьютер недоступен")
	}
	r, err := cn.ReadResponse(conn)
	if err != nil {
		cls.errConn(id)
		ToLog(id, funcNameLog+"cn.ReadResponse(conn)", false)
		return nil, "", errors.New("компьютер недоступен")
	}
	if r.Err != nil {
		cls.clientAccessFree(id)
		ToLog(id, funcNameLog+"r.Err: "+fmt.Sprint(r.Err), false)
		return nil, "", r.Err
	}
	cn.SendSync(conn)
	data, err := cn.ReadBytesByLen(r.DataLen, conn)
	if err != nil {
		cls.errConn(id)
		ToLog(id, funcNameLog+"cn.ReadBytesByLen(r.DataLen, conn)", false)
		return nil, "", errors.New("компьютер недоступен")
	}
	defer cls.clientAccessFree(id)
	res := make(map[string][]ex.LinkAndName)
	err = json.Unmarshal(data, &res)
	if err != nil {
		ToLog(id, funcNameLog+"json.Unmarshal(data, &res)", false)
		return nil, "", errors.New("ошибка данных")
	}
	return res, cls.sep(id), nil
}
