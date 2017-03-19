package tcp_client

import "fmt"
import "net"
import "time"
import "bufio"
import "testing"

func TestClientTCP(t *testing.T) {
	go CreateServer(t)

	time.Sleep(10 * time.Millisecond)

	client := New("127.0.0.1:1232")

	client.OnOpen(func() {
		fmt.Printf("Conneted on %s\n\n", client.Address)
		client.WriteString("OK, close")
	})

	client.OnMessage(func(message []byte) {
		println(string(message))
	})

	client.OnError(func(err error) {
		if !client.Connected {
			t.Fatal(err.Error())
		}
	})

	client.Listen()
}

func CreateServer(t *testing.T) {
	server, err := net.Listen("tcp", ":1232")

	if err != nil {
		t.Fatal(err.Error())
		return
	} else {
		for {
			connections, _ := server.Accept()

			go NewClient(server, connections)
		}
	}
}

func NewClient(s net.Listener, con net.Conn) {
	reader := bufio.NewReader(con)

	for {
		buf := make([]byte, 1024)
		num, _ := reader.Read(buf)

		message := make([]byte, num)
		copy(message, buf)

		if string(message) == "OK, close" {

			con.Write([]byte("OK"))

			con.Close()
			s.Close()
		}
	}
}
