package core

import (
  "os"
  "rygel/common"
	"path/filepath"
)

type StoreRepo struct {
  Dir string
  Stores []Store
}

func (sr StoreRepo) FindByName(name string) *Store {
  for _, store := range sr.Stores {
    if store.Name == name {
      return &store
    }
  }

  panic("Could not find store")
}

func (sr *StoreRepo) Create(name string) {
  for _, store := range sr.Stores {
    if store.Name == name {
      panic("Store already exists")
    }
  }

  f, err := os.Create(sr.Dir + name + ".store")

  defer f.Close()

  common.HandleErr(err)

  _, err2 := f.WriteString(name + "\n")

  common.HandleErr(err2)

  store := BuildStore(name)

  sr.Stores = append(sr.Stores, store)
}

func InitializeFromDir(dir string) StoreRepo {
  stores := []Store{}

  err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
    storeFile, err := os.OpenFile(path, os.O_RDWR, 0644)

    var out []byte
    storeFile.Read(out)

    stores = append(stores, BuildStore(string(out)))

    return nil
  })

  common.HandleErr(err)

  return StoreRepo{Dir: dir, Stores: stores}
}
