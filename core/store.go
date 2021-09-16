package core

import (
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

// func (s *Store) RemoveItem(collectionName string, key string) bool {
//   collectionRef, err := s.referenceCollection(collectionName) 

//   if err != nil {
//     fmt.Println(err)
//     return false
//   }

//   return collectionRef.RemoveItem(key)
// }

// func (s *Store) AddIndex(collectionName string, index Index) bool {
//   collectionRef, err := s.referenceCollection(collectionName) 

//   if err != nil {
//     fmt.Println(err)
//     return false
//   }

//   collectionRef.AddIndex(index)

//   return true
// }


func BuildStore() Store {
  store := Store{Collections: map[string]Collection{}}

  return store
}
