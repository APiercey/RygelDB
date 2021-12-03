package store_repo

import (
  "errors"
  str "rygel/core/store"
)

type InMemoryRepo struct {
  Stores []str.Store
}

func (sr *InMemoryRepo) appendStore(store str.Store) {
  sr.Stores = append(sr.Stores, store)
}

func (sr InMemoryRepo) FindByName(name string) (foundStore *str.Store, err error) {
  return findByName(sr.Stores, name)
}

func (sr InMemoryRepo) Create(name string) (store *str.Store, err error) {
  _, err = findByName(sr.Stores, name)

  if err != nil {
    return nil, errors.New("Store already exists")
  }

  builtStore := str.BuildStore(name)
  sr.appendStore(builtStore)

  return &builtStore, nil
}
