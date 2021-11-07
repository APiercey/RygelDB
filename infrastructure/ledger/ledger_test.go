package ledger

import (
	"testing"
	"github.com/stretchr/testify/assert"
  "os"
)

func assertAppendAndReplay(t *testing.T, l Ledger) {
  first_record := "First Record"
  second_record := "Second Record"
  third_record := "Third Record"

  l.AppendRecord(first_record)
  l.ReplayRecords(func(line string) {})
  l.AppendRecord(second_record)
  l.ReplayRecords(func(line string) {})
  l.AppendRecord(third_record)

  result := []string{}

  l.ReplayRecords(func(line string) { result = append(result, line) })

  expected_result := []string{first_record, second_record, third_record}
  assert.Equal(t, expected_result, result)
}

func TestAppendingToInMemoryLedger(t *testing.T) {
  ledger := InMemoryLedger{}

  assertAppendAndReplay(t, &ledger)
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

  assertAppendAndReplay(t, &ledger)
}

