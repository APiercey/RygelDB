package main

import (
	"bufio"
	"fmt"
	"net"
  "example.com/rygel/core" 
  "example.com/rygel/commands" 
  "example.com/rygel/services" 
  "example.com/rygel/servers" 
)

var storePersistenceService = services.StorePersistenceService{
  DiskLocation: "./store.db",
}

func ExecuteStatementAgainstStore(store *core.Store, statement string) (result string, store_was_updated bool) {
  command, err := commands.CommandParser(statement)

  if err != nil {
    return err.Error(), false
  }

  if !command.Valid() {
    return "Command not valid", false
  }

  result, s := command.Execute(store)

  return result, s
}

func buildConnectionHandler(store *core.Store) func(conn net.Conn) {
  return func(conn net.Conn) {
    for {
      buffer, err := bufio.NewReader(conn).ReadBytes('\n')

      if err != nil {
        fmt.Println("Client left.")
        conn.Close()
        return
      }

      result, store_was_updated := ExecuteStatementAgainstStore(store, string(buffer[:len(buffer)-1]))

      if store_was_updated {
        storePersistenceService.PersistDataToDisk(store)
      }

      conn.Write([]byte(result))
    }
  }
}

func main() {
  store := core.BuildStore()
  storePersistenceService.LoadDataFromDisk(&store)

  servers.StartSocketServer(buildConnectionHandler(&store))
}
