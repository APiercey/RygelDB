package store

import (
	"encoding/json"
	"fmt"
	"os"
)

type Store struct {
  diskLocation string
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

func (s *Store) CreateCollection(collectionName string) bool {
  s.Collections[collectionName] = BuildCollection(collectionName)
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

func (s *Store) InsertItem(collectionName string, item Item) bool {
  collectionRef, err := s.referenceCollection(collectionName) 

  if err != nil {
    fmt.Println(err)
    return false
  }

  return collectionRef.InsertItem(item)
}

func (s *Store) PersistToDisk() {
  file, err := os.Create(s.diskLocation)

  if err != nil {
    fmt.Println(err)
  }

  encoder := json.NewEncoder(file)
  err = encoder.Encode(s.Collections)

  if err != nil {
    fmt.Println(err)
  }
}

func (s *Store) loadFromDisk() {
  if !fileExists(s.diskLocation) {
    return
  }
  
  var collections map[string]Collection

  file, err := os.Open(s.diskLocation)

  if err != nil {
    fmt.Println(err)
  }

  decoder := json.NewDecoder(file)
  err = decoder.Decode(&collections)

  if err != nil {
    fmt.Println(err)
    panic(err)
  }

  s.Collections = collections
}

func (s *Store) RemoveItem(collectionName string, key string) bool {
  collectionRef, err := s.referenceCollection(collectionName) 

  if err != nil {
    fmt.Println(err)
    return false
  }

  return collectionRef.RemoveItem(key)
}

func (s *Store) AddIndex(collectionName string, index Index) bool {
  collectionRef, err := s.referenceCollection(collectionName) 

  if err != nil {
    fmt.Println(err)
    return false
  }

  collectionRef.AddIndex(index)

  return true
}

func fileExists(filename string) bool {
    info, err := os.Stat(filename)

    if os.IsNotExist(err) { return false }

    return !info.IsDir()
}

func BuildStore(diskLocation string) Store {
  store := Store{Collections: map[string]Collection{}, diskLocation: diskLocation}
  store.loadFromDisk()

  return store
}
