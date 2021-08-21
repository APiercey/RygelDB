package store

// import (
//   "errors"
// )

type Item struct {
  Key string
  Data map[string]interface{}
}

func BuildItem(key string, data map[string]interface{}) (Item, error) {
  return Item{Key: key, Data: data}, nil
}
