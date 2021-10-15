package services

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
)

type BasicAuthService struct {
  ConfiguredUsername string
  ConfiguredPassword string
}

func (service BasicAuthService) getCredentials(conn net.Conn) (username string, password string, err error) {
    buffer, err := bufio.NewReader(conn).ReadBytes('\n')
    authDetails := string(buffer[:len(buffer)-1])

    if err != nil { return "", "", errors.New("Connection lost") }

    creds := strings.Split(authDetails, " ")

    if len(creds) == 0 { return "", "", nil }
    if len(creds) == 1 { return creds[0], "", nil }
    if len(creds) == 2 { return creds[0], creds[1], nil }

    return "", "", errors.New("Could not understand authentication credentials")
}

func (service BasicAuthService) Authenticate(conn net.Conn) bool {
  username, password, err := service.getCredentials(conn)

  if err != nil {
    fmt.Println(err.Error())
    conn.Write([]byte(err.Error()))
    return false
  }

  return service.ConfiguredUsername == username &&
         service.ConfiguredPassword == password
}

