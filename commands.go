package main

import (
	"errors"
	"strings"
)

type Command interface {
  execute(store *Store) (*Item, error)
}

type InsertCommand struct {
  key string
  data string
}

type QueryCommand struct {
  key string
}

func (c InsertCommand) execute(s *Store) (*Item, error) {
  if s.InsertItem(c.key, c.data) {
    item, _ := s.ReadByKey(c.key)
    return &item, nil
  } else {
    return nil, errors.New("Could not store object.")
  }
}

func (c QueryCommand) execute(s *Store) (*Item, error) {
  item, presence := s.ReadByKey(c.key)

  if presence {
    return &item, nil
  } else {
    return nil, errors.New("Could not find object")
  }
}

func CommandParser(rawCommand string) (Command, error) {
  seperatedCommandStrings := strings.Split(rawCommand, " ")

  switch seperatedCommandStrings[0] {
  case "INSERT":
    return InsertCommand{key: seperatedCommandStrings[1], data: seperatedCommandStrings[2]}, nil
  case "QUERY":
    return QueryCommand{key: seperatedCommandStrings[1]}, nil
  default:
    return nil, errors.New("Unknown instruction.")
  }
}
