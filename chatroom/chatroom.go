package chatroom

import (
  "fmt"
  "strings"
  "strconv"
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

  for i := 1; i <= len(cr.clients)+1; i++ {
    if !inClientArray(i, cr.clients) {
        return i;
      }
  }
  return 1
}

func (cr *ChatRoom) RemoveClient(cli *Client) {
  for i, clia := range cr.clients {
    if clia.id == cli.id {
      cr.clients = append(cr.clients[:i], cr.clients[i+1:]...)
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

func inClientArray(i int, clia []*Client) bool {
  for _, cli := range clia {
    if i == cli.id {
      return true
    }
  }

  return false
}

func toIntSlice(strs []string) []int {
    list := make([]int, 0)

    for _, str := range strs {
      num, _ := strconv.ParseInt(str, 10, 0)
      list = append(list, int(num))
    }

    return list
}

func (cr *ChatRoom) ClientListener(cli *Client) {
  for{
    msg, err := cli.reader.ReadString('\n')

    if err != nil {
      cr.RemoveClient(cli)
      cli.conn.Close()
      return
    }

    fmt.Printf("Client: %d, Message: %s\n", cli.id, msg)

    sms := strings.Split(strings.TrimSpace(msg), " ")

    if sms[0] == "getid" {
      dmsg := fmt.Sprintf("Client-id: %d\n", cli.id)
      cli.EchoToClient(dmsg)
    }
    if sms[0] == "list" {
      dmsg := fmt.Sprintf("Client-ids: %v\n", cr.GetClientIds())
      cli.EchoToClient(dmsg)
    }
    if sms[0] == "send" {

      targets := toIntSlice(strings.Split(sms[1], ","))

      dmsg := fmt.Sprintf("Client: %d, Message: %s\n", cli.id, sms[2])

      cli.SendToNamedReceipents(dmsg, targets, cr)
    }
  }
}
