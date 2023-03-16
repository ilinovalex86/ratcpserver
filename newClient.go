package ratcpserver

import (
	"encoding/json"
	"errors"
	"fmt"
	cn "github.com/ilinovalex86/connection"
	"net"
)

//Создание и получение данных о новом клиенте
func newClient(conn net.Conn) (*client, error) {
	const funcNameLog = "newClient(): "
	cl := &client{conn: conn, status: true}
	err := cl.valid()
	if err != nil {
		return nil, errors.New(funcNameLog + fmt.Sprint(err))
	}
	cl.Id, err = cn.ReadString(conn)
	if err != nil {
		return nil, errors.New(funcNameLog + "cn.ReadString(conn)")
	}
	cn.SendSync(conn)
	data, err := cn.ReadByteByDelim(cl.conn)
	if err != nil {
		return nil, errors.New(funcNameLog + "cn.ReadByteByDelim(cl.conn)")
	}
	err = json.Unmarshal(data, &cl)
	if err != nil {
		return nil, errors.New(funcNameLog + "json.Unmarshal(data, &cl)")
	}
	return cl, nil
}
