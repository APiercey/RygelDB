package servers

import (
  "fmt"
  "net"
  "os"
)

const (
  connHost = "localhost"
  connPort = "8080"
  connType = "tcp"
)

func StartSocketServer(connectionHandler func (conn net.Conn)) {
  fmt.Println("Starting " + connType + " server on " + connHost + ":" + connPort)

  l, err := net.Listen(connType, connHost+":"+connPort)

  if err != nil {
    fmt.Println("Error listening:", err.Error())
    os.Exit(1)
  }

  defer l.Close()

  for {
    c, err := l.Accept()

    if err != nil {
      fmt.Println("Error connecting:", err.Error())
      return
    }

    fmt.Println("Client connected.")

    fmt.Println("Client " + c.RemoteAddr().String() + " connected.")

    go connectionHandler(c)
  }
}
