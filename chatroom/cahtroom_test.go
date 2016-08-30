package chatroom

import (
  "testing"
  "net"
)

type MockConn struct {
  net.Conn
}

func NewMockConn() *MockConn {
  return &MockConn{}
}

var cr *ChatRoom
var cli *Client

func TestChatroom(t *testing.T) {

  cr := NewChatRoom()
  cli := NewClient(NewMockConn(), cr.GetFreeId())

  cr.AddClient(cli)

  if len(cr.clients) != 1 {
    t.Error("Expected client count 1")
  }

  if cr.clients[0].id != 1 {
    t.Error("Expected client with id 1")
  }

  if cr.GetFreeId() != 2 {
    t.Error("Expected next client id to be 2")
  }

  cr.RemoveClient(cli)

  if len(cr.clients) != 0 {
    t.Error("Expected client count to be 0")
  }
}
