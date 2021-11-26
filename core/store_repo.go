package core

type StoreRepo struct {
  Stores []*Store
}

func (sr StoreRepo) FindByName(name string) *Store {
  for _, store := range sr.Stores {
    if store.Name == name {
      return store
    }
  }

  panic("Could not find store")
}
