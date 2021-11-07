package ledger

type InMemoryLedger struct {
  records []string
}

func (l *InMemoryLedger) Append(data string) {
  l.records = append(l.records, data)
}

func (l InMemoryLedger) Read() []string {
  return l.records
}
