package store

type Collection struct {
  Name string
  Items map[string]Item
}

func (c Collection) ReadByKey(key string) (Item, bool) {
  item, presence := c.Items[key]
  return item, presence
}

func (c *Collection) InsertItem(key string, data map[string]interface{}) bool {
  c.Items[key] = Item{Key: key, Data: data}
  return true
}

func BuildCollection(collectionName string) Collection {
  return Collection{Name: collectionName, Items: map[string]Item{}}
}
