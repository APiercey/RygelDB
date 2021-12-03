package commands

import (
  cs "rygel/core/store" 
)

type DefineStoreCommand struct {
  StoreRepo cs.StoreRepo
  StoreName string
}

func (c DefineStoreCommand) Execute() (string, bool) {
  _, err := c.StoreRepo.Create(c.StoreName)
  
  if err != nil {
    return err.Error(), false
  }

  return "Store " + c.StoreName + " created", true
}

func (c DefineStoreCommand) Valid() bool {
  if c.StoreName == "" {
    return false
  }

  return true
}

