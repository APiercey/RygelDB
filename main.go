package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func buildConnectionHandler(store *Store) func(conn net.Conn) {
  return func(conn net.Conn) {
    for {
      buffer, err := bufio.NewReader(conn).ReadBytes('\n')

      if err != nil {
        fmt.Println("Client left.")
        conn.Close()
        return
      }

      log.Println("Client message:", string(buffer[:len(buffer)-1]))

      command, err := CommandParser(string(buffer[:len(buffer)-1]))

      if err == nil {
        item, err := command.execute(store)

        if err == nil {
          conn.Write([]byte(item.Data))
        } else {
          fmt.Println(err.Error())
        }
      } else {
        fmt.Println(err.Error())
      }
    }
  }
}

func main() {
  store := Store{Items: map[string]Item{}}
  startSocketServer(buildConnectionHandler(&store))
}
