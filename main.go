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

      if err != nil {
        fmt.Println(err.Error())
        conn.Close()
        return
      }

      conn.Write([]byte(command.execute(store)))
    }
  }
}

func main() {
  store := BuildStore()

  startSocketServer(buildConnectionHandler(&store))
}
