package ledger

type replayFn func(string)

type Ledger interface {
  AppendRecord(data string)
  ReplayRecords(fn replayFn)
}

