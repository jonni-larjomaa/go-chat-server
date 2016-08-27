package main

import (
  "bufio"
  "fmt"
  "net"
)

func clientListener(c net.Conn) {
  reader := bufio.NewReader(c)
  writer := bufio.NewWriter(c)

  for {
    msg, err := reader.ReadString('\n')

    if err != nil {
      fmt.Println("Connection error, client dropped")
      c.Close()
      return
    }

    writer.WriteString(msg)
    writer.Flush()
  }

}

func main() {

	listener, _ := net.Listen("tcp", ":9090")

	for {
		conn, _ := listener.Accept()

    if conn != nil {
      fmt.Println("new connection accepted")
    }

		go clientListener(conn)
	}
}
