package ratcpserver

import (
	"errors"
	"fmt"
	cn "github.com/ilinovalex86/connection"
)

//Создание и отправка id клиенту
func (cl *client) isNewClient() error {
	const funcNameLog = "cl.isNewClient(): "
	id := generatorId()
	err := cn.SendQuery(cn.Query{Method: "new id"}, cl.conn)
	if err != nil {
		return errors.New(funcNameLog + "cn.SendQuery(cn.Query{Method: \"new id\"}, cl.conn)")
	}
	cn.ReadSync(cl.conn)
	err = cn.SendString(id, cl.conn)
	if err != nil {
		return errors.New(funcNameLog + "cn.SendString(id, cl.conn)")
	}
	cl.Id = id
	Clients.store(cl)
	err = Clients.dump()
	if err != nil {
		ToLog(tcpServer, funcNameLog+fmt.Sprint(err), true)
	}
	return nil
}
