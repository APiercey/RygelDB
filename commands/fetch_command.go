package commands

import (
	"encoding/json"
	"fmt"
	"strings"
	"example.com/kv_store/store"
)

type FetchCommand struct {
  limit int
  collectionName string
  wherePredicates []wherePredicate
}

type wherePredicate struct {
  path string
  value string
}

func (wp wherePredicate) filter(item store.Item) bool {
  split := strings.Split(wp.path, ".")
  steps := split[:len(split) - 1]
  key := split[len(split) - 1]

  structure := item.Data

  for _, step := range steps {
    traversedStructure, presence := structure[step]

    if presence {
      structure = traversedStructure.(map[string]interface{})
    } else {
      // It would be possible to traverse into arrays
      // but I wont implement this yet
      return false
    }
  }

  return structure[key] == wp.value
}

func (c FetchCommand) Execute(s *store.Store) (string, bool) {
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


