package store

import (
	"encoding/gob"
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

func (s *Store) InsertItem(collectionName string, key string, data map[string]interface{}) bool {
  collectionRef, err := s.referenceCollection(collectionName) 

  if err != nil {
    fmt.Println(err)
    return false
  }

  collectionRef.InsertItem(key, data)

  return true
}

func (s *Store) PersistToDisk() {
  file, err := os.Create(s.diskLocation)

  if err != nil {
    fmt.Println(err)
  }

  encoder := gob.NewEncoder(file)
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

  file, err := os.Create(s.diskLocation)

  if err != nil {
    fmt.Println(err)
  }

  decoder := gob.NewDecoder(file)
  err = decoder.Decode(&collections)

  fmt.Println(collections)

  if err != nil {
    fmt.Println(err)
    panic(err)
  }

  s.Collections = collections
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