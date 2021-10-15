package main

import (
	"bufio"
	"fmt"
	"net"

	"example.com/rygel/core"
	"example.com/rygel/servers"
	"example.com/rygel/application"
)

func buildConnectionHandler(store *core.Store, application application.Application) func(conn net.Conn) {
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

      result, store_was_updated := application.StatementExecutionService.Execute(store, string(buffer[:len(buffer)-1]))

      if store_was_updated {
        application.StorePersistenceService.PersistDataToDisk(store)
      }

      conn.Write([]byte(result))
    }
  }
}

func main() {
  store := core.BuildStore()

  application := application.New()

  application.StorePersistenceService.LoadDataFromDisk(&store)

  connectionHandler := buildConnectionHandler(
    &store,
    application,
  )

  servers.StartSocketServer(connectionHandler)
}
