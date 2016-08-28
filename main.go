package main

import (
  "fmt"
  "net"

  "github.com/jonni-larjomaa/go-chat-server/chatroom"
)

var cr *chatroom.ChatRoom

func main() {

  cr := chatroom.NewChatRoom()

	listener, _ := net.Listen("tcp", ":9090")

	for {
		conn, _ := listener.Accept()

    if conn != nil {
      fmt.Println("new connection accepted")
    }

    client := chatroom.NewClient(conn, cr.GetFreeId())

    cr.AddClient(client)

		go cr.ClientListener(client)
	}
}
