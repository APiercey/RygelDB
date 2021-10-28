package main

import (
	"bufio"
	"fmt"
	"net"

	"rygel/servers"
	"rygel/application"
)

func buildConnectionHandler(application *application.Application) func(conn net.Conn) {
  return func(conn net.Conn) {
    if !application.BasicAuthService.Authenticate(conn) {
      conn.Write([]byte("Could not authenticate"))
      conn.Close()
      return
    }

    for {
      buffer, err := bufio.NewReader(conn).ReadBytes('\n')

      if err != nil {
        fmt.Println("Client left.")
        conn.Close()
        return
      }

      result, store_was_updated := application.StatementExecutionService.Execute(string(buffer[:len(buffer)-1]))

      if store_was_updated {
        application.StorePersistenceService.PersistDataToDisk(&application.Store)
      }

      conn.Write([]byte(result))
    }
  }
}

func main() {
  application := application.New()

  application.StorePersistenceService.LoadDataFromDisk()

  go func() { for { application.CommandExecutor.Process() } }()

  connectionHandler := buildConnectionHandler(
    &application,
  )

  servers.StartSocketServer(connectionHandler)
}
