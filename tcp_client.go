package tcp_client

import (
	"bufio"
	"net"
)

type Connection struct {
	onOpenCallback    func()
	onMessageCallback func(message []byte)
	onErrorCallback   func(err error)

	Conn      net.Conn
	Address   string
	Connected bool
}

func (self *Connection) OnOpen(f func()) {
	self.onOpenCallback = f
}

func (self *Connection) OnMessage(f func(message []byte)) {
	self.onMessageCallback = f
}

func (self *Connection) OnError(f func(err error)) {
	self.onErrorCallback = f
}

func (self *Connection) Close() {
	self.Conn.Close()
}

func (self *Connection) Write(message []byte) {
	self.Conn.Write(message)
}

func (self *Connection) WriteString(message string) {
	self.Conn.Write([]byte(message))
}

func (self *Connection) Listen() {
	conexao, err := net.Dial("tcp", self.Address)

	if err != nil {
		self.onErrorCallback(err)
	} else {
		defer conexao.Close()
		self.Conn = conexao

		self.Connected = true
		self.onOpenCallback()
		self.read()
	}
}

func (self *Connection) read() {
	reader := bufio.NewReader(self.Conn)

	for {
		buf := make([]byte, 1024)
		num, err := reader.Read(buf)

		if err != nil {
			self.Close()
			self.onErrorCallback(err)
			return
		}

		mensagem := make([]byte, num)
		copy(mensagem, buf)

		self.onMessageCallback(mensagem)
	}
}

func New(address string) *Connection {
	conexao := &Connection{Address: address, Connected: false}

	conexao.OnOpen(func() {})
	conexao.OnError(func(err error) {})
	conexao.OnMessage(func(message []byte) {})

	return conexao
}