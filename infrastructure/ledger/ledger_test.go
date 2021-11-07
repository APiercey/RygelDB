package ledger

import (
	"testing"
	"github.com/stretchr/testify/assert"
  "os"
)

func assertAppendRecord(t *testing.T, l Ledger) {
  record := "this is a test record"
  expected_result := []string{record}

  l.AppendRecord(record)

  result := []string{}

  l.ReplayRecords(func(line string) { result = append(result, line) })

  assert.Equal(t, expected_result, result)
}

func TestAppendingToInMemoryLedger(t *testing.T) {
  ledger := InMemoryLedger{}

  assertAppendRecord(t, &ledger)
}

func TestAppendingToOnDiskLedger(t *testing.T) {
  f, err := os.OpenFile("./test.ledger", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)

  if err != nil { panic(err) }

  defer func() {
    f.Close()
    os.Remove("./test.ledger")
  }()

  ledger := OnDiskLedger{
    LedgerFile: f,
  }

  assertAppendRecord(t, &ledger)
}

