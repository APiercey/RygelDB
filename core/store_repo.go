package core

import (
)

type StoreRepo struct {
  Store *Store
}

func (sr StoreRepo) FindByName(name string) *Store {
  return sr.Store
}
