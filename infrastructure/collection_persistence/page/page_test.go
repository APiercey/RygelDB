package page

import (
	"os"
	"rygel/common"
	"rygel/core"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func useTempDataFile(f func(*os.File)) {
  file, err := os.OpenFile("./test.page", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
  if err != nil { panic(err) }

  defer func() {
    file.Close()
    os.Remove("./test.page")
  }()

  f(file)
}

func TestAppend(t *testing.T) {
  useTempDataFile(func(f *os.File) {
    page := New(f)
    item, _ := core.BuildItem(common.Data{"Hello": "World"})

    before, _ := f.Stat()
    assert.True(t, before.Size() == 0)

    page.Append(item)

    after, _ := f.Stat()
    assert.True(t, after.Size() > 0)
  })
}

func TestReadingValues(t *testing.T) {
  useTempDataFile(func(f *os.File) {
    page := New(f)
    item, _ := core.BuildItem(common.Data{"Hello": "World"})
    page.Append(item)

    page.Scan(func(scannedItem *core.Item) bool {
      assert.Equal(t, item.Data, scannedItem.Data)
      return true
    }, false)
  })
}

func TestLockingPage(t *testing.T) {
  useTempDataFile(func(f *os.File) {
    page := New(f)
    item, _ := core.BuildItem(common.Data{"Hello": "World"})
    page.Append(item)

    page.Scan(func(scannedItem *core.Item) bool {
      assert.True(t, page.IsLocked())
      return true
    }, true)

    assert.False(t, page.IsLocked())
  })
}

func TestLockingPageSecondProcessWaits(t *testing.T) {
  useTempDataFile(func(f *os.File) {
    page := New(f)
    item, _ := core.BuildItem(common.Data{"Hello": "World"})
    page.Append(item)

    before := time.Now()
    waitTime := 50 * time.Millisecond

    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
      page.Scan(func(scannedItem *core.Item) bool {
        time.Sleep(waitTime)
        return true
      }, true)
      wg.Done();
    }()

    go func() {
      time.Sleep(1 * time.Microsecond)
      page.Scan(func(scannedItem *core.Item) bool { return true }, false)
      wg.Done();
    }()

    wg.Wait();

    diff := time.Now().Sub(before)
    assert.GreaterOrEqual(t, diff, waitTime)
  })
}
