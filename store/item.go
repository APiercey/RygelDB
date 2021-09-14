package store

type Item struct {
  Data map[string]interface{}
  IsStale bool
}

func (i *Item) MarkAsStale() {
  i.IsStale = true
}

func BuildItem(data map[string]interface{}) (Item, error) {
  return Item{Data: data, IsStale: false}, nil
}
