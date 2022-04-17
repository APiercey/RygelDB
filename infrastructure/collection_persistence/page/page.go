package page

import (
	"os"
	"bufio"
	"rygel/common"
	"rygel/core"
	// "encoding/json"
)

type EnumFunction func(*core.Item) bool

type Page struct {
  file *os.File
  locked bool
}

// TODO: Create a "Persistable" interface
func (p *Page) Scan(f EnumFunction, lockPage bool) {
  action := func() {
    dataUpdated := false
    keepItems := []core.Item{}

    p.file.Seek(0, 0)
    scanner := bufio.NewScanner(p.file)

    for scanner.Scan() {
      item := toItem(scanner.Bytes())

      shouldContinue := f(&item)

      if item.WasUpdated() {
        dataUpdated = true
      }

      if item.ShouldBeRemoved() {
        dataUpdated = true
      } else {
        keepItems = append(keepItems, item)
      }

      if !shouldContinue {
        break;
      }
    }

    if dataUpdated {
      p.write(keepItems)
    }
  }

  p.lockedAction(action, lockPage)
}

func (p *Page) Append(item core.Item) {
  action := func() {
    _, err := p.file.WriteString(string(common.EncodeData(item.Data)) + "\n")

    if err != nil { panic(err) }
  }

  p.lockedAction(action, true)
}

func (p *Page) lockedAction(f func(), useLock bool) {
  p.waitForLock()

  if useLock { p.lock() }

  f()

  if p.IsLocked() { p.unlock() }
}

func (p Page) write(items []core.Item) {
  datas := [][]byte{}

  for _, item := range items {
    datas = append(datas, common.EncodeData(item.Data))
  }

  p.file.Truncate(0)
  p.file.Seek(0, 0)
  
  for _, data := range datas {
    p.file.WriteString(string(data) + "\n")
  }

}

func (p *Page) lock() {
  p.locked = true
}

func (p *Page) unlock() {
  p.locked = false
}

func (p Page) IsLocked() bool {
  return p.locked
}

func (p Page) waitForLock() {
  for p.IsLocked() {}
}

func New(file *os.File) Page {
  return Page{file: file, locked: false}
}

func toItem(rawData []byte) core.Item {
  data := common.DecodeData(rawData)

  item, _ := core.BuildItem(data)

  return item
}

