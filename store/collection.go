package store

type Collection struct {
  Name string
  Items map[string]Item
  indices map[string]Index
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

func (c *Collection) AddIndex(index Index) {
  c.indices[index.path] = index
}

func BuildCollection(collectionName string) Collection {
  collection := Collection{Name: collectionName, Items: map[string]Item{}}
  // collection.AddIndex(BuildIndex("__primaryId"))

  return collection
}
