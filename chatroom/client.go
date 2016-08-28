package chatroom

import (
  "bufio"
  "net"
  //"strings"
  //"fmt"
)

type Client struct {
  id int
  conn net.Conn
  reader *bufio.Reader
  writer *bufio.Writer
}

func NewClient(c net.Conn, id int) *Client {

  client := &Client{
    id: id,
    conn: c,
    reader: bufio.NewReader(c),
    writer: bufio.NewWriter(c),
  }

  return client
}

func (cli *Client) SendToNamedReceipents(msg string, others []int, hub *ChatRoom){
  for _, cli := range hub.clients {
    if inArray(cli.id, others) {
      cli.writer.WriteString(msg)
      cli.writer.Flush()
    }
  }
}

func (cli *Client) EchoToClient(msg string) {
  cli.writer.WriteString(msg)
  cli.writer.Flush()
}

func inArray(a int, list []int) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}
