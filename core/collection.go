package core

type Collection struct {
  Name string
  Items []Item
  indices map[string]Index
}

func (c *Collection) InsertItem(item Item) bool {
  c.Items = append(c.Items, item)

  return true
}

// TODO: Make this work with WHERE clause
// func (c *Collection) RemoveItem(key string) bool {
//   _, ok := c.Items[key];

//   if ok {
//     delete(c.Items, key);
//     return true
//   } else {
//     return false
//   }
// }

// func (c *Collection) AddIndex(index Index) {
//   c.indices[index.path] = index
// }

func BuildCollection(collectionName string) Collection {
  collection := Collection{Name: collectionName, Items: []Item{}}
  // collection.AddIndex(BuildIndex("__primaryId"))

  return collection
}
