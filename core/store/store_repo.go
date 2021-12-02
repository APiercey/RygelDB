package store

import (
  "os"
  "rygel/common"
	"path/filepath"
  "errors"
  "io/ioutil"
  "strings"
)

type StoreRepo struct {
  Dir string
  Stores []Store
}

func (sr StoreRepo) FindByName(name string) (store *Store, err error) {
  for _, store := range sr.Stores {
    if store.Name == name {
      return &store, nil
    }
  }


  return nil, errors.New("Store not found")
}

func (sr *StoreRepo) Create(name string) (store *Store, err error) {
  for _, store := range sr.Stores {
    if store.Name == name {
      return nil, errors.New("Store already exists")
    }
  }

  f, err := os.Create(sr.Dir + "/" + name + ".store")

  defer f.Close()

  common.HandleErr(err)

  _, err2 := f.WriteString(name + "\n")

  common.HandleErr(err2)

  newStore := BuildStore(name)

  sr.Stores = append(sr.Stores, newStore)

  return &newStore, nil
}

func InitializeFromDir(dir string) StoreRepo {
  stores := []Store{}

  if _, err := os.Stat(dir); os.IsNotExist(err) { 
    os.MkdirAll(dir, 0700) 
  }

  err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
    if info.IsDir() {
      return nil
    }

    out, err := ioutil.ReadFile(path)

    common.HandleErr(err)

    name := strings.TrimRight(string(out), "\n")

    stores = append(stores, BuildStore(name))

    return nil
  })

  common.HandleErr(err)

  return StoreRepo{Dir: dir, Stores: stores}
}
