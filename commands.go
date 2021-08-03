package main

import (
	"errors"
	"strings"
)

type Command interface {
  execute(store *Store) string
}

type CreateCollectionCommand struct {
  collectionName string
}

type InsertCommand struct {
  collectionName string
  key string
  data string
}

type QueryCommand struct {
  collectionName string
  key string
}

func (c CreateCollectionCommand) execute(s *Store) string {
  if s.createCollection(c.collectionName) {
    return "OK"
  } else {
    return "ERR"
  }
}

func (c InsertCommand) execute(s *Store) string {
  if s.InsertItem(c.collectionName, c.key, c.data) {
    return "OK"
  } else {
    return "ERR"
  }
}

func (c QueryCommand) execute(s *Store) string {
  item, presence := s.Collections[c.collectionName].ReadByKey(c.key)

  if presence {
    return item.Data
  } else {
    return ""
  }
}

func CommandParser(rawCommand string) (Command, error) {
  seperatedCommandStrings := strings.Split(rawCommand, " ")

  switch seperatedCommandStrings[0] {
  case "CREATE":
    return CreateCollectionCommand{collectionName: seperatedCommandStrings[1]}, nil
  case "INSERT":
    return InsertCommand{collectionName: seperatedCommandStrings[1], key: seperatedCommandStrings[2], data: seperatedCommandStrings[3]}, nil
  case "QUERY":
    return QueryCommand{collectionName: seperatedCommandStrings[1], key: seperatedCommandStrings[2]}, nil
  default:
    return nil, errors.New("Unknown instruction.")
  }
}
