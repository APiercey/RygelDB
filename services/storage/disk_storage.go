package storage

import (
  "encoding/json"
  "fmt"
  "os"

  "rygel/core"
  "rygel/services/ledger"
)

type DiskStorage struct {
  Store *core.Store
  Ledger ledger.Ledger
}

func (service *DiskStorage) LoadData() {
  if !fileExists(service.ledgerPath()) {
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

func (service DiskStorage) CreateSnapshot() {
  data, err := json.Marshal(service.Store.Collections)

  if err != nil {
    panic(err)
  }

  service.Ledger.Append(string(data))
}

