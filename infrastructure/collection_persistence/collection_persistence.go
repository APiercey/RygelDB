package collection_persistence

import (
	"os"
  // "strings"
	"path/filepath"
	"rygel/core"
  "rygel/infrastructure/collection_persistence/page"
)

type CollectionPersistence struct {
  Name string
  collectionDir string
  pages []page.Page
}

func (cp CollectionPersistence) Enumerate(f page.EnumFunction, useLock bool) {
  for _, page := range cp.pages {
    page.Scan(f, useLock) 
  }
}

func (cp *CollectionPersistence) InsertItem(item core.Item) {
  page := cp.getLastPage()
  page.Append(item)
}

func (cp *CollectionPersistence) appendNewPage(filePath string) {
  file, _ := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
  page := page.New(file)

  cp.pages = append(cp.pages, page)
}

func (cp *CollectionPersistence) getLastPage() page.Page {
  if len(cp.pages) == 0 {
    filePath := cp.collectionDir + "/" + "page_0.sec"

    cp.appendNewPage(filePath)
  }

  return cp.pages[len(cp.pages)-1]
}

func buildPages(collectionDir string) []page.Page {
  pages := []page.Page{}

  err := filepath.Walk(collectionDir, func(path string, info os.FileInfo, err error) error {
    if path != collectionDir {
      pageFile, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)

      if err != nil { panic(err) }

      pages = append(pages, page.New(pageFile))
    }

    return nil
  })

  if err != nil { panic(err) }

  return pages
}

func New(name string, collectionDir string) CollectionPersistence {
  return CollectionPersistence{
    Name: name,
    collectionDir: collectionDir,
    pages: buildPages(collectionDir),
  }
}
