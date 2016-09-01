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

func TestChatroom(t *testing.T) {

  cr := NewChatRoom()
  cliar := make([]*Client,0)

  // add ten clients.
  for i := 0; i < 10; i++ {
      cli := NewClient(NewMockConn(), cr.GetFreeId())
      cliar = append(cliar, cli)
      cr.AddClient(cli)
  }

  if len(cr.clients) != 10 {
    t.Error("Expected client count 1")
  }

  if cr.clients[0].id != 1 {
    t.Error("Expected client with id 1")
  }

  if cr.GetFreeId() != 11 {
    t.Error("Expected next client id to be 11")
  }

  cr.RemoveClient(cliar[5])

  if cr.GetFreeId() != 6 {
    t.Error("Expected next free id to be 6")
  }

  if len(cr.clients) != 9 {
    t.Error("Expected client count to be 9")
  }

  if cr.inClientArray(1) != true {
    t.Error("Expected client with id 1 to be in client array")
  }
}
