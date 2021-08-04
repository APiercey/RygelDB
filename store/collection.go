package store

type Collection struct {
  Name string
  Items map[string]Item
}

func (c Collection) ReadByKey(key string) (Item, bool) {
  item, presence := c.Items[key]
  return item, presence
}

func (c *Collection) InsertItem(item Item) bool {
  c.Items[item.Key] = item
  return true
}

func (c *Collection) RemoveItem(key string) bool {
  _, ok := c.Items[key];

  if ok {
    delete(c.Items, key);
    return true
  } else {
    return false
  }
}

func BuildCollection(collectionName string) Collection {
  return Collection{Name: collectionName, Items: map[string]Item{}}
}
