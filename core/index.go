package core

import (
	"example.com/rygel/common"
)

type Index struct {
  DataPath common.DataPath
  referencedItems map[interface{}][]*Item
}

func (i *Index) indexItem(item *Item) {
  value, present := item.PluckValueOnPath(i.DataPath)

  if !present { return }

  i.ensureReferencedValue(value)

  i.referencedItems[value] = append(i.referencedItems[value], item)
}

func (i *Index) ensureReferencedValue(value interface{}) {
  if !i.ContainsValue(value) {
    i.referencedItems[value] = []*Item{}
  }
}

func (i Index) ContainsValue(value interface{}) bool {
  _, ok := i.referencedItems[value]

  return ok
}

func (i Index) CopiedItems(value interface{}) []Item {
  items := []Item{}

  for _, item := range i.referencedItems[value] {
    items = append(items, *item)
  }

  return items
}

func BuildIndex(path []string) Index {
  return Index{
    DataPath: common.DataPath{RealPath: path},
    referencedItems: map[interface{}][]*Item{},
  }
}

