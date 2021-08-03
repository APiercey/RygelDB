package main

import(
  "fmt"
)

type Store struct {
  Collections map[string]Collection
}

func (s *Store) referenceCollection(collectionName string) (*Collection, error) {
  collection, ok := s.Collections[collectionName]

  if ok {
    return &collection, nil
  } else {
    return &Collection{}, fmt.Errorf("Could not find collection")
  }
}

func (s *Store) createCollection(collectionName string) bool {
  s.Collections[collectionName] = BuildCollection(collectionName)
  return true
}

func (s *Store) InsertItem(collectionName string, key string, data string) bool {
  collectionRef, err := s.referenceCollection(collectionName) 

  if err != nil {
    fmt.Println(err)
    return false
  }

  collectionRef.InsertItem(key, data)

  return true
}

func BuildStore() Store {
  return Store{Collections: map[string]Collection{}}
}
