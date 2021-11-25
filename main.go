package main

import (
	"bufio"
	"fmt"
	"net"

	"rygel/infrastructure/socket_server"
	"rygel/application"
  "rygel/context"
)

func buildConnectionHandler(application *application.Application) func(conn net.Conn) {
  return func(conn net.Conn) {
    if !application.BasicAuth.Authenticate(conn) {
      conn.Write([]byte("Could not authenticate"))
      conn.Close()
      return
    }

    ctx := context.Context{SelectedStore: "prod"}

    for {
      buffer, err := bufio.NewReader(conn).ReadBytes('\n')

      if err != nil {
        fmt.Println("Client left.")
        conn.Close()
        return
      }

      result := application.StatementExecutor.Execute(ctx, string(buffer[:len(buffer)-1]))

      conn.Write([]byte(result))
    }
  }
}

func main() {
  application := application.New()

  go func() { for { application.CommandExecutor.Process() } }()

  application.StatementReplay.Replay(context.Context{SelectedStore: "prod"})

  connectionHandler := buildConnectionHandler(
    &application,
  )

  socket_server.StartSocketServer(connectionHandler)
}
