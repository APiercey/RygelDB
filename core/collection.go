package core

type Collection struct {
  Name string
  Items []Item
}

func (c *Collection) InsertItem(item Item) bool {
  c.Items = append(c.Items, item)

  return true
}

func (c *Collection) ReplaceItems(items []Item) {
  c.Items = items
}

func BuildCollection(collectionName string) Collection {
  collection := Collection{Name: collectionName, Items: []Item{}}

  return collection
}
