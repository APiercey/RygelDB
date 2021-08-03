package store

import (
  "encoding/json"
  "fmt"
)

type Command interface {
  Execute(store *Store) (result string, store_was_updated bool)
}

type CreateCollectionCommand struct {
  collectionName string
}

type InsertCommand struct {
  collectionName string
  key string
  data map[string]interface{}
}

type FetchCommand struct {
  limit int
  collectionName string
}

type LookupCommand struct {
  collectionName string
  key string
}

func (c CreateCollectionCommand) Execute(s *Store) (string, bool) {
  if s.CreateCollection(c.collectionName) {
    return "OK", true
  } else {
    return "ERR", false
  }
}

func (c InsertCommand) Execute(s *Store) (string, bool) {
  if s.InsertItem(c.collectionName, c.key, c.data) {
    return "OK", true
  } else {
    return "ERR", false
  }
}

func (c FetchCommand) Execute(s *Store) (string, bool) {
  items := []string{}
  numFoundItems := 0

  fmt.Println(c)
  for _, item := range s.Collections[c.collectionName].Items {

    if numFoundItems == c.limit {
      break
    }

    out, err := json.Marshal(item.Data)

    if err != nil { panic (err) }

    items = append(items, string(out))
    numFoundItems += 1
  }

  out, err := json.Marshal(items)

  if err != nil { panic (err) }

  return string(out), false
}

func (c LookupCommand) Execute(s *Store) (string, bool) {
  item, presence := s.Collections[c.collectionName].ReadByKey(c.key)

  fmt.Println(item)

  if presence {
    out, err := json.Marshal(item.Data)

    if err != nil { panic (err) }

    return string(out), false
  } else {
    return "", false
  }
}
