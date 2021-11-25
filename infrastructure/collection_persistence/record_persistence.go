package record_persistence

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"rygel/common"
	"rygel/core"
)

type EnumFunction func(*core.Item)

type CollectionPersistence struct {
  collectionDir string
}

func (cp CollectionPersistence) Enumerate(f EnumFunction) {
  collectionUpdated := false

  err := filepath.Walk(cp.collectionDir, func(path string, info os.FileInfo, err error) error {
    sectorFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
    keepItems := []core.Item{}

    if err != nil {
      panic(err)
    }

    scanner := bufio.NewScanner(sectorFile)

    for scanner.Scan() {
      item := toItem(scanner.Bytes())

      f(&item)

      if item.WasUpdated() {
        collectionUpdated = true
      }

      if item.ShouldRemove() {
        collectionUpdated = true
      } else {
        keepItems = append(keepItems, item)
      }
    }

    if collectionUpdated {
      persist(sectorFile, keepItems)
    }

    return nil
  })

  if err != nil {
      panic(err)
  }
}

func (cp CollectionPersistence) Store(item core.Item) {
  f, err := os.OpenFile(cp.getLastSectorFile(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

  if err != nil {
    panic(err)
  }

  defer f.Close()

  out := common.EncodeData(item.Data)

  if _, err = f.Write(out); err != nil {
    panic(err)
  }
}

func (cp CollectionPersistence) getLastSectorFile() string {
  var writeToFile string;
  filePaths := []string{}

  err := filepath.Walk(cp.collectionDir, func(path string, info os.FileInfo, err error) error {
    filePaths = append(filePaths, path)
    return nil
  })

  if err != nil {
    panic(err)
  }

  if len(filePaths) > 0 {
    writeToFile = filePaths[len(filePaths)-1]
  } else {
    writeToFile = "part_0.sec"
  }

  return writeToFile
}


func toItem(rawData []byte) core.Item {
  data := common.DecodeData(rawData)

  item, _ := core.BuildItem(data)

  return item
}

func persist(f *os.File, items []core.Item) {
  datas := [][]byte{}

  for _, item := range items {
    datas = append(datas, common.EncodeData(item.Data))
  }

  out, err := json.Marshal(datas)

  if err != nil {
    panic(err)
  }

  f.Seek(0, 0)
  f.Write(out)
}
