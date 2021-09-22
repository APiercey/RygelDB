package core

type Collection struct {
  Name string
  Items []Item
  Indices map[string]Index
}

func (c *Collection) InsertItem(item Item) bool {
  c.Items = append(c.Items, item)

  return true
}

func (c *Collection) AddIndex(index Index) {
  c.Indices[index.dataPath.SerializedPath()] = index
}

func (c Collection) IndexedPaths() []string {
  keys := make([]string, 0, len(c.Indices))

  for k := range c.Indices {
    keys = append(keys, k)
  }

  return keys
}

func BuildCollection(collectionName string) Collection {
  collection := Collection{Name: collectionName, Items: []Item{}}

  return collection
}
