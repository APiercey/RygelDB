package store

import (
	"fmt"
	"os"
	"rygel/core"
	"rygel/common"

	coll "rygel/infrastructure/collection_persistence"
)

type Store struct {
  Name string
  Collections map[string]coll.CollectionPersistence
  locationOnDisk string
}

func (s *Store) referenceCollection(collectionName string) (*coll.CollectionPersistence, error) {
  collection, ok := s.Collections[collectionName]

  if ok {
    return &collection, nil
  } else {
    return &coll.CollectionPersistence{}, fmt.Errorf("Could not find collection")
  }
}

func (s *Store) CreateCollection(collectionName string) bool {
  filePath := s.locationOnDisk + "/" + collectionName

  os.Mkdir(filePath, 0755)
  s.Collections[collectionName] = coll.New(collectionName, filePath)

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

  collection.InsertItem(item)

  s.Collections[collectionName] = collection
  
  return true
}

func buildExistingCollections(storeDir string) map[string]coll.CollectionPersistence {
  collections := map[string]coll.CollectionPersistence{}

  for _, name := range common.CollectPathsInDir(storeDir) {
    collections[name] = coll.New(name, storeDir + "/" + name)
  }

  return collections
}

func BuildStore(name string, locationOnDisk string) Store {
  return Store{
    Name: name,
    Collections: buildExistingCollections(locationOnDisk),
    locationOnDisk: locationOnDisk,
  }
}
