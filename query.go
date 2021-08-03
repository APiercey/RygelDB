package main

import (
	"errors"
)

type Query interface {
  fetch(store Store) bool
}

type DataQuery struct {
  collectionName string
  key string
}

// type DefinitionQuery struct {
//   // collectionName string
//   // key string
// }

func (c DataQuery) fetch(s Store) (string, error) {
  item, presence := s.Collections[c.collectionName].ReadByKey(c.key)

  if presence {
    return item.Data, nil
  } else {
    return "", errors.New("Could not find object")
  }
}
