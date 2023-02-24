package ratcpserver

import (
	"crypto/aes"
	cn "github.com/ilinovalex86/connection"
	"log"
)

//Валидация клиента
func (cl *client) valid() bool {
	var code = make([]byte, 16)
	s := generatorId()
	ver, err := cn.ReadString(cl.conn)
	if err != nil {
		return false
	}
	if _, errB := keys[ver]; !errB {
		return false
	}
	bc, err := aes.NewCipher([]byte(keys[ver]))
	if err != nil {
		log.Fatal(err)
	}
	bc.Encrypt(code, []byte(s))
	err = cn.SendBytes(code, cl.conn)
	if err != nil {
		return false
	}
	pass := s[len(s)/2:] + s[:len(s)/2]
	res, err := cn.ReadBytesByLen(16, cl.conn)
	if err != nil {
		return false
	}
	bc.Decrypt(code, res)
	if string(code) != pass {
		return false
	}
	err = cn.SendString("ok", cl.conn)
	if err != nil {
		return false
	}
	return true
}
