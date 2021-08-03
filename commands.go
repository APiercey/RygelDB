package main

import (
  "encoding/json"
  "fmt"
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

type FetchCommand struct {
  limit int
  collectionName string
}

type LookupCommand struct {
  collectionName string
  key string
}

type NoopCommand struct {

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

func (c FetchCommand) execute(s *Store) string {
  items := []string{}
  numFoundItems := 0

  fmt.Println(c)
  for _, item := range s.Collections[c.collectionName].Items {

    if numFoundItems == c.limit {
      break
    }

    items = append(items, item.Data)
    numFoundItems += 1
  }

  out, err := json.Marshal(items)

  if err != nil {
    panic (err)
  }

  return string(out)
}

func (c LookupCommand) execute(s *Store) string {
  item, presence := s.Collections[c.collectionName].ReadByKey(c.key)

  if presence {
    return item.Data
  } else {
    return ""
  }
}

func (c NoopCommand) execute(s *Store) string {
  return "NOOP"
}

