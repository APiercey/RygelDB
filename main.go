package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
  "example.com/rygel/store" 
  "example.com/rygel/commands" 
)

func buildConnectionHandler(currentStore *store.Store) func(conn net.Conn) {
  return func(conn net.Conn) {
    for {
      buffer, err := bufio.NewReader(conn).ReadBytes('\n')

      if err != nil {
        fmt.Println("Client left.")
        conn.Close()
        return
      }

      log.Println("Client message:", string(buffer[:len(buffer)-1]))

      command, err := commands.CommandParser(string(buffer[:len(buffer)-1]))

      if err != nil {
        fmt.Println(err.Error())
        conn.Close()
        return
      }

      result, store_was_updated := command.Execute(currentStore)

      if store_was_updated {
        currentStore.PersistToDisk()
      }

      conn.Write([]byte(result))
    }
  }
}

func main() {
  store := store.BuildStore("./store.db")

  startSocketServer(buildConnectionHandler(&store))
}
