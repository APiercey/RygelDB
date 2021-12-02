package store

import (
	"fmt"
  "rygel/core"
)

type Store struct {
  Name string
  Collections map[string]core.Collection
}

func (s *Store) referenceCollection(collectionName string) (*core.Collection, error) {
  collection, ok := s.Collections[collectionName]

  if ok {
    return &collection, nil
  } else {
    return &core.Collection{}, fmt.Errorf("Could not find collection")
  }
}

func (s *Store) CreateCollection(collectionName string) bool {
  s.Collections[collectionName] = core.BuildCollection(collectionName)
  return true
}

func (s *Store) UndefineCollection(collectionName string) bool {
  _, ok := s.Collections[collectionName];

  if ok {
    delete(s.Collections, collectionName);
    return true
  } else {
    return false
  }
}

func (s *Store) InsertItem(collectionName string, item core.Item) bool {
  collection, present := s.Collections[collectionName]

  if !present {
    fmt.Println("Collection does not exist")
    return false
  }

  if !collection.InsertItem(item) {
    fmt.Println("Could not insert item into collection")
    return false
  }

  s.Collections[collectionName] = collection
  
  return true
}

func BuildStore(name string) Store {
  store := Store{Name: name, Collections: map[string]core.Collection{}}

  return store
}
