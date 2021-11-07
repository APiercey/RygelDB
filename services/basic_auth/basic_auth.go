package basic_auth

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
)

type BasicAuth struct {
  ConfiguredUsername string
  ConfiguredPassword string
}

func extractCreds(credsKV string) map[string]string {
  foundCreds := map[string]string{
    "user": "",
    "pass": "",
  }

  for _, cred := range strings.Split(credsKV, " ") {
    kv := strings.Split(cred, "=")

    switch kv[0] {
    case "user":
      foundCreds["user"] = kv[1]
    case "pass":
      foundCreds["pass"] = kv[1]
    }
  }

  return foundCreds
}

func (service BasicAuth) getCredentials(conn net.Conn) (username string, password string, err error) {
    buffer, err := bufio.NewReader(conn).ReadBytes('\n')

    if err != nil { return "", "", errors.New("Connection lost") }

    creds := extractCreds(string(buffer[:len(buffer)-1]))

    return creds["user"], creds["pass"], nil
}

func (service BasicAuth) Authenticate(conn net.Conn) bool {
  username, password, err := service.getCredentials(conn)

  if err != nil {
    fmt.Println(err.Error())
    conn.Write([]byte(err.Error()))
    return false
  }

  return service.ConfiguredUsername == username &&
         service.ConfiguredPassword == password
}

