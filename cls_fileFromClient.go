package ratcpserver

import (
	"errors"
	"fmt"
	cn "github.com/ilinovalex86/connection"
)

// FileFromClient - скачивает файл с клиента
func (cls *clients) FileFromClient(id string, path string) (string, error) {
	const funcNameLog = "cls.FileFromClient(): "
	conn, err := cls.clientAccessBlock(id)
	if err != nil {
		ToLog(id, funcNameLog+"cls.clientAccessBlock(id): "+fmt.Sprint(err), false)
		return "", err
	}
	err = cn.SendQuery(cn.Query{Method: "fileFromClient", Query: path}, conn)
	if err != nil {
		cls.errConn(id)
		ToLog(id, funcNameLog+"cn.SendQuery(cn.Query{Method: \"fileFromClient\", Query: path}, conn)", false)
		return "", errors.New("компьютер недоступен")
	}
	r, err := cn.ReadResponse(conn)
	if err != nil {
		cls.errConn(id)
		ToLog(id, funcNameLog+"cn.ReadResponse(conn)", false)
		return "", errors.New("компьютер недоступен")
	}
	if r.Err != nil {
		cls.clientAccessFree(id)
		ToLog(id, funcNameLog+"r.Err: "+fmt.Sprint(r.Err), false)
		return "", r.Err
	}
	cn.SendSync(conn)
	tempPath := tempFileName()
	err = cn.GetFile(tempPath, r.DataLen, conn)
	if err != nil {
		cls.errConn(id)
		ToLog(id, funcNameLog+fmt.Sprintf("cn.GetFile(%s, %d, conn) %s", tempPath, r.DataLen, err), false)
		return "", err
	}
	cls.clientAccessFree(id)
	return tempPath, nil
}
