package ratcpserver

import (
	"errors"
	cn "github.com/ilinovalex86/connection"
)

// FileFromClient - скачивает файл с клиента
func (cls *clients) FileFromClient(id string, path string) (string, error) {
	conn, err := cls.clientAccessBlock(id)
	if err != nil {
		return "", err
	}
	err = cn.SendQuery(cn.Query{Method: "fileFromClient", Query: path}, conn)
	if err != nil {
		cls.errConn(id)
		return "", errors.New("компьютер недоступен")
	}
	r, err := cn.ReadResponse(conn)
	if err != nil {
		cls.errConn(id)
		return "", errors.New("компьютер недоступен")
	}
	if r.Err != nil {
		cls.clientAccessFree(id)
		return "", r.Err
	}
	cn.SendSync(conn)
	tempPath := tempFileName()
	err = cn.GetFile(tempPath, r.DataLen, conn)
	if err != nil {
		cls.errConn(id)
		return "", err
	}
	cls.clientAccessFree(id)
	return tempPath, nil
}
