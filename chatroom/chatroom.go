package chatroom

import (
  "fmt"
  "strings"
  "encoding/json"
)

type ChatRoom struct {
	clients []*Client
}

func NewChatRoom() *ChatRoom {
	chatRoom := &ChatRoom{
		clients: make([]*Client, 0),
	}

	return chatRoom
}

func (cr *ChatRoom) AddClient(cli *Client) {
  cr.clients = append(cr.clients, cli)
}

func (cr *ChatRoom) GetFreeId() int {

  for id := 1; id <= len(cr.clients)+1; id++ {
    if !cr.inClientArray(id) {
        return id;
      }
  }
  return 1
}

func (cr *ChatRoom) RemoveClient(cli *Client) {
  for position, clia := range cr.clients {
    if clia.id == cli.id {
      cr.clients = append(cr.clients[:position], cr.clients[position+1:]...)
    }
  }
}

func (cr *ChatRoom) GetClientIds() []int {
  ar := make([]int,0)
  for _, clia := range cr.clients {
    ar = append(ar, clia.id)
  }

  return ar
}

func (cr *ChatRoom) ClientListener(cli *Client) {
  for{

    msg, err := cli.reader.ReadString('\n')

    if err != nil {
      cr.RemoveClient(cli)
      cli.conn.Close()
      return
    }

    var jsonmsg Message

    e := json.Unmarshal([]byte(strings.TrimSpace(msg)),&jsonmsg);

    if e != nil {
      fmt.Println("Error parsing received json")
      continue;
    }

    switch {
    case jsonmsg.Cmd == "getid":
        cli.EchoToClient(fmt.Sprintf("Client-id: %d\n", cli.id))
      case jsonmsg.Cmd == "list":
        cli.EchoToClient(fmt.Sprintf("Client-ids: %v\n", cr.GetClientIds()))
      case jsonmsg.Cmd == "send":
        cli.SendToNamedReceipents(fmt.Sprintf("Client: %d, Message: %s\n", cli.id, jsonmsg.Msg), jsonmsg.Receiv, cr.clients)
    }
  }
}

func (cr *ChatRoom) inClientArray(id int) bool {
  for _, cli := range cr.clients {
    if id == cli.id {
      return true
    }
  }

  return false
}
