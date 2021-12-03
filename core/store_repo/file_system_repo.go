package store_repo

import (
  "os"
  "rygel/common"
	"path/filepath"
  "errors"
  "io/ioutil"
  "strings"
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
  _, err = findByName(sr.Stores, name)

  if err != nil {
    return nil, errors.New("Store already exists")
  }

  f, err := os.Create(sr.Dir + "/" + name + ".store")

  defer f.Close()

  common.HandleErr(err)

  _, err2 := f.WriteString(name + "\n")
  common.HandleErr(err2)
  builtStore := str.BuildStore(name)
  sr.appendStore(builtStore)

  return &builtStore, nil
}

func InitializeFromDir(dir string) StoreRepo {
  stores := []str.Store{}

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
    stores = append(stores, str.BuildStore(name))

    return nil
  })

  common.HandleErr(err)

  return FileSystemRepo{Dir: dir, Stores: stores}
}
