package ratcpserver

import (
	"crypto/aes"
	"errors"
	cn "github.com/ilinovalex86/connection"
)

//Валидация клиента
func (cl *client) valid() error {
	const funcNameLog = "cl.valid(): "
	var code = make([]byte, 16)
	s := generatorId()
	ver, err := cn.ReadString(cl.conn)
	if err != nil {
		return errors.New(funcNameLog + "cn.ReadString(cl.conn)")
	}
	if _, errB := keys[ver]; !errB {
		return errors.New(funcNameLog + "if _, errB := keys[ver]; !errB")
	}
	bc, err := aes.NewCipher([]byte(keys[ver]))
	if err != nil {
		ToLog(tcpServer, funcNameLog+"aes.NewCipher([]byte(keys[ver]))", true)
	}
	bc.Encrypt(code, []byte(s))
	err = cn.SendBytes(code, cl.conn)
	if err != nil {
		return errors.New(funcNameLog + "cn.SendBytes(code, cl.conn)")
	}
	pass := s[len(s)/2:] + s[:len(s)/2]
	res, err := cn.ReadBytesByLen(16, cl.conn)
	if err != nil {
		return errors.New(funcNameLog + "cn.ReadBytesByLen(16, cl.conn)")
	}
	bc.Decrypt(code, res)
	if string(code) != pass {
		return errors.New(funcNameLog + "string(code) != pass")
	}
	err = cn.SendString("ok", cl.conn)
	if err != nil {
		return errors.New(funcNameLog + "cn.SendString(\"ok\", cl.conn)")
	}
	return nil
}
