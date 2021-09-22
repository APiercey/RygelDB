package core

import (
	"example.com/rygel/common"
)

type Index struct {
  dataPath common.DataPath
  referencedItems map[interface{}][]*Item
}

func (i *Index) indexItem(item *Item) {
  value, present := item.PluckValueOnPath(i.dataPath)

  if !present { return }

  i.ensureReferencedValue(value)

  i.referencedItems[value] = append(i.referencedItems[value], item)
}

func (i *Index) ensureReferencedValue(value interface{}) {
  if !i.containsValue(value) {
    i.referencedItems[value] = []*Item{}
  }
}

func (i Index) containsValue(value interface{}) bool {
  _, ok := i.referencedItems[value]

  return ok
}

func BuildIndex(path []string) Index {
  return Index{
    dataPath: common.DataPath{RealPath: path},
    referencedItems: map[interface{}][]*Item{},
  }
}
