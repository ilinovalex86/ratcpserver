package ratcpserver

import (
	"encoding/json"
	"errors"
	cn "github.com/ilinovalex86/connection"
	"net"
)

//Создание и получение данных о новом клиенте
func newClient(conn net.Conn) (*client, error) {
	cl := &client{conn: conn, status: true}
	if !cl.valid() {
		return nil, errors.New("valid error")
	}
	var err error
	cl.Id, err = cn.ReadString(conn)
	if err != nil {
		return nil, err
	}
	cn.SendSync(conn)
	data, err := cn.ReadByteByDelim(cl.conn)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &cl)
	if err != nil {
		return nil, err
	}
	return cl, nil
}
