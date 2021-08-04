package store

type Item struct {
  Key string
  Data map[string]interface{}
}

func BuildItem(key string, data map[string]interface{}) Item {
  return Item{Key: key, Data: data}
}
