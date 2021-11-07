package ledger

type InMemoryLedger struct {
  records []string
}

func (l *InMemoryLedger) AppendRecord(data string) {
  l.records = append(l.records, data)
}

func (l InMemoryLedger) ReplayRecords(fn replayFn) {

  for _, line := range l.records {
    fn(line)
  }
}
