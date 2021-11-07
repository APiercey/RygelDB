package storage

import (
  "rygel/core"
)

type MemoryStorage struct {
  LoggedStatements []string
}

func (service *MemoryStorage) LogCommand(rawStatement string) {
  service.LoggedStatements = append(service.LoggedStatements, rawStatement)
}

func (service *MemoryStorage) LoadData() {
  
}

func (service MemoryStorage) PersistData(store *core.Store) {

}

