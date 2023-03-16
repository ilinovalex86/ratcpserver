package ratcpserver

import (
	"errors"
	"fmt"
	cn "github.com/ilinovalex86/connection"
)

// FileToClient - отправляет файл клиенту
func (cls *clients) FileToClient(id string, data *[]byte, fileName string, fileSize int64) error {
	const funcNameLog = "cls.FileToClient(): "
	conn, err := cls.clientAccessBlock(id)
	if err != nil {
		ToLog(id, funcNameLog+"cls.clientAccessBlock(id): "+fmt.Sprint(err), false)
		return err
	}
	err = cn.SendQuery(cn.Query{Method: "fileToClient", Query: fileName, DataLen: int(fileSize)}, conn)
	if err != nil {
		cls.errConn(id)
		ToLog(id, funcNameLog+"cn.SendQuery(cn.Query{Method: \"fileToClient\", Query: fileName, DataLen: int(fileSize)}, conn)", false)
		return errors.New("компьютер недоступен")
	}
	cn.ReadSync(conn)
	err = cn.SendBytes(*data, conn)
	if err != nil {
		cls.errConn(id)
		ToLog(id, funcNameLog+"cn.SendBytes(*data, conn)", false)
		return errors.New("компьютер недоступен")
	}
	defer cls.clientAccessFree(id)
	r, err := cn.ReadResponse(conn)
	if r.Err != nil {
		ToLog(id, funcNameLog+"r.Err: "+fmt.Sprint(r.Err), false)
		return r.Err
	}
	return nil
}
