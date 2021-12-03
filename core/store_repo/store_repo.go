package store_repo

import (
  "rygel/core/store"
)

type StoreRepo interface {
  FindByName(name string) (store *store.Store, err error)
  Create(name string) (store *store.Store, err error)
}

