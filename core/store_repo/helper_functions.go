package store_repo

import (
  "errors"
  str "rygel/core/store"
)

func findByName(stores []str.Store, name string) (foundStore *str.Store, err error) {
  for _, store := range stores {
    if store.Name == name {
      return &store, nil
    }
  }

  return nil, errors.New("Store not found")
}
