package services

import (
	"encoding/json"
	"fmt"
	"os"

  "example.com/rygel/store" 
)

type StorePersistenceService struct {
  DiskLocation string
}

func fileExists(diskLocation string) bool {
    info, err := os.Stat(diskLocation)

    if os.IsNotExist(err) { return false }

    return !info.IsDir()
}

func (service StorePersistenceService) LoadDataFromDisk(_store *store.Store) {
  if !fileExists(service.DiskLocation) {
    return
  }
  
  var collections map[string]store.Collection

  file, err := os.Open(service.DiskLocation)

  if err != nil {
    fmt.Println(err)
  }

  decoder := json.NewDecoder(file)
  err = decoder.Decode(&collections)

  if err != nil {
    fmt.Println(err)
    panic(err)
  }

  _store.Collections = collections
}

func (service StorePersistenceService) PersistDataToDisk(_store *store.Store) {
  file, err := os.Create(service.DiskLocation)

  if err != nil {
    fmt.Println(err)
  }

  encoder := json.NewEncoder(file)
  err = encoder.Encode(_store.Collections)

  if err != nil {
    fmt.Println(err)
  }
}

