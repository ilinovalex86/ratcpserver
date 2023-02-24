package ratcpserver

func (cls *clients) streamNew(cl *client, ss [2]int) *stream {
	cls.mu.Lock()
	defer cls.mu.Unlock()
	s := &stream{id: cl.Id, conn: cl.conn, webClients: []string{cls.m[cl.Id].streamRunUser}, ScreenSizeX: ss[0], ScreenSizeY: ss[1]}
	cls.m[cl.Id].streamStatus = true
	cls.m[cl.Id].streamLink = s
	return s
}
