package ratcpserver

import (
	"errors"
	"fmt"
	cn "github.com/ilinovalex86/connection"
)

//Обработка Client с id
func (cl *client) isOldClient() error {
	const funcNameLog = "cl.isOldClient(): "
	ok := Clients.Exist(cl.Id)
	if !ok || ok && cl.BasePath != Clients.basePath(cl.Id) {
		err := cl.isNewClient()
		return errors.New(funcNameLog + fmt.Sprint(err))
	}
	if ok && Clients.testConnect(cl.Id) {
		_ = cn.SendQuery(cn.Query{Method: "already exist"}, cl.conn)
		return errors.New(funcNameLog + "already exist")
	}
	ver := Clients.getVersion(cl.Id)
	Clients.store(cl)
	if ver != cl.Version {
		err := Clients.dump()
		if err != nil {
			ToLog(tcpServer, funcNameLog+"Clients.dump()", true)
		}
	}
	err := cn.SendQuery(cn.Query{Method: "connect"}, cl.conn)
	if err != nil {
		return errors.New(funcNameLog + "cn.SendQuery(cn.Query{Method: \"connect\"}, cl.conn)")
	}
	return nil
}
