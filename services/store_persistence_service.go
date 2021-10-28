package services

import (
	"encoding/json"
	"fmt"
	"os"

  "rygel/core" 
)

type StorePersistenceService struct {
  DiskLocation string
  Store *core.Store
}

func fileExists(diskLocation string) bool {
    info, err := os.Stat(diskLocation)

    if os.IsNotExist(err) { return false }

    return !info.IsDir()
}

func (service *StorePersistenceService) LoadDataFromDisk() {
  if !fileExists(service.DiskLocation) {
    return
  }
  
  var collections map[string]core.Collection

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

  // TODO: Perhaps a SetCollections() func would be good?
  store := service.Store
  store.Collections = collections
}

func (service StorePersistenceService) PersistDataToDisk(store *core.Store) {
  file, err := os.Create(service.DiskLocation)

  if err != nil {
    fmt.Println(err)
  }

  encoder := json.NewEncoder(file)
  err = encoder.Encode(store.Collections)

  if err != nil {
    fmt.Println(err)
  }
}

