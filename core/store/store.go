package store

import (
	"fmt"
	"os"
	"rygel/core"
  "io/ioutil"
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

func collectionPaths(dir string) []string {
  dirs := []string{}

  files, err := ioutil.ReadDir(dir)

  common.HandleErr(err)

  for _, file := range files {
    if file.IsDir() {
      dirs = append(dirs, file.Name())
    }
  }
  
  return dirs
}

func buildExistingCollections(storeDir string) map[string]coll.CollectionPersistence {
  collections := map[string]coll.CollectionPersistence{}

  fmt.Println("####")
  fmt.Println(storeDir)
  fmt.Println("####")

  collectionPaths := collectionPaths(storeDir)

  for _, name := range collectionPaths {
    collections[name] = coll.New(name, storeDir + "/" + name)
  }

  // err := filepath.Walk(storeDir, func(path string, info os.FileInfo, err error) error {

  //   // splits := strings.Split(path, "/")
  //   // name := splits[len(splits)-1]

  //   collections[name] = coll.New(name, path)

  //   return nil
  // })

  // if err != nil { panic(err) }

  return collections
}

func BuildStore(name string, locationOnDisk string) Store {
  return Store{
    Name: name,
    Collections: buildExistingCollections(locationOnDisk),
    locationOnDisk: locationOnDisk,
  }
}
