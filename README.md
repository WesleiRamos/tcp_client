# tcp_client
Lib to create tcp clients fater

## Usage:

``` go
package main

import client "github.com/WesleiRamos/tcp_client"

func main() {
	conexao := client.New("127.0.0.1:1232")

	conexao.OnOpen(func() {
		conexao.WriteString("EAE MAN")
		println("Conectou-se")
	})

	conexao.OnMessage(func(message string) {
		println("Messagem: " + message)
	})

	conexao.OnError(func(err error) {
		if !conexao.Connected {
			println(err.Error())
		}
	})

	conexao.Listen()
}
```