package ratcpserver

import (
	"errors"
	cn "github.com/ilinovalex86/connection"
)

// FileToClient - отправляет файл клиенту
func (cls *clients) FileToClient(id string, data *[]byte, fileName string, fileSize int64) error {
	conn, err := cls.clientAccessBlock(id)
	if err != nil {
		return err
	}
	err = cn.SendQuery(cn.Query{Method: "fileToClient", Query: fileName, DataLen: int(fileSize)}, conn)
	if err != nil {
		cls.errConn(id)
		return errors.New("компьютер недоступен")
	}
	cn.ReadSync(conn)
	err = cn.SendBytes(*data, conn)
	if err != nil {
		cls.errConn(id)
		return errors.New("компьютер недоступен")
	}
	defer cls.clientAccessFree(id)
	r, err := cn.ReadResponse(conn)
	if r.Err != nil {
		return r.Err
	}
	return nil
}
