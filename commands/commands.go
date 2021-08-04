package store

import (
  "encoding/json"
)

type WherePredicate struct {
  path string
  value string
}

func (wp WherePredicate) filter(item Item) bool {
  return item.Data[wp.path] == wp.value
}

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
  wherePredicates []WherePredicate
}

type LookupCommand struct {
  collectionName string
  key string
}

type RemoveCommand struct {
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
  item := BuildItem(c.key, c.data)

  if s.InsertItem(c.collectionName, item) {
    return "OK", true
  } else {
    return "ERR", false
  }
}

func (c FetchCommand) Execute(s *Store) (string, bool) {
  items := []string{}
  numFoundItems := 0

  for _, item := range s.Collections[c.collectionName].Items {

    if numFoundItems == c.limit {
      break
    }

    meetsPredicates := true

    Predicates:
    for _, wp := range c.wherePredicates {
      meetsPredicates = wp.filter(item)

      if !meetsPredicates { break Predicates }
    }

    if meetsPredicates {
      out, err := json.Marshal(item.Data)

      if err != nil { panic (err) }

      items = append(items, string(out))
      numFoundItems += 1
    }
  }

  out, err := json.Marshal(items)

  if err != nil { panic (err) }

  return string(out), false
}

func (c LookupCommand) Execute(s *Store) (string, bool) {
  item, presence := s.Collections[c.collectionName].ReadByKey(c.key)

  if presence {
    out, err := json.Marshal(item.Data)

    if err != nil { panic (err) }

    return string(out), false
  } else {
    return "", false
  }
}

func (c RemoveCommand) Execute(s *Store) (string, bool) {
  if s.RemoveItem(c.collectionName, c.key) {
    return "OK", true
  } else {
    return "ERR", false
  }
}
