package store_repo

import (
  "fmt"
  "os"
  "rygel/common"
  "errors"
  str "rygel/core/store"
)

type FileSystemRepo struct {
  Dir string
  Stores []str.Store
}

func (sr *FileSystemRepo) appendStore(store str.Store) {
  sr.Stores = append(sr.Stores, store)
}

func (sr FileSystemRepo) FindByName(name string) (foundStore *str.Store, err error) {
  return findByName(sr.Stores, name)
}

func (sr FileSystemRepo) Create(name string) (store *str.Store, err error) {
  existingStore, _ := findByName(sr.Stores, name)

  if existingStore != nil {
    return nil, errors.New("Store already exists")
  }

  f, err := os.Create(sr.Dir + "/" + name + ".store")

  defer f.Close()

  common.HandleErr(err)

  err = os.Mkdir(sr.Dir + "/" + name, 0755)

  common.HandleErr(err)

  _, err = f.WriteString(name + "\n")

  common.HandleErr(err)

  builtStore := str.BuildStore(name, sr.Dir + "/" + name)

  sr.appendStore(builtStore)

  return &builtStore, nil
}

func InitializeFromDir(dir string) StoreRepo {
  stores := []str.Store{}

  if _, err := os.Stat(dir); os.IsNotExist(err) { 
    os.MkdirAll(dir, 0700) 
  }

  for _, name := range common.CollectPathsInDir(dir) {
    fmt.Println(name)

    stores = append(stores, str.BuildStore(name, dir + "/" + name))
  }

  return FileSystemRepo{Dir: dir, Stores: stores}
}
