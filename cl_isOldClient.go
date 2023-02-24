package ratcpserver

import (
	"errors"
	cn "github.com/ilinovalex86/connection"
)

//Обработка Client с id
func (cl *client) isOldClient() error {
	ok := Clients.Exist(cl.Id)
	if !ok || ok && cl.BasePath != Clients.basePath(cl.Id) {
		err := cl.isNewClient()
		return err
	}
	if ok && Clients.testConnect(cl.Id) {
		_ = cn.SendQuery(cn.Query{Method: "already exist"}, cl.conn)
		return errors.New("already exist")
	}
	ver := Clients.getVersion(cl.Id)
	Clients.store(cl)
	if ver != cl.Version {
		err := Clients.dump()
		check(err)
	}
	err := cn.SendQuery(cn.Query{Method: "connect"}, cl.conn)
	if err != nil {
		return err
	}
	return nil
}
