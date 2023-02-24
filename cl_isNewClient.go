package ratcpserver

import (
	cn "github.com/ilinovalex86/connection"
)

//Создание и отправка id клиенту
func (cl *client) isNewClient() error {
	id := generatorId()
	err := cn.SendQuery(cn.Query{Method: "new id"}, cl.conn)
	if err != nil {
		return err
	}
	cn.ReadSync(cl.conn)
	err = cn.SendString(id, cl.conn)
	if err != nil {
		return err
	}
	cl.Id = id
	Clients.store(cl)
	err = Clients.dump()
	check(err)
	return nil
}
