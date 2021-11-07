package ledger

import (
	"bufio"
	"os"
)

type OnDiskLedger struct {
  LedgerFile *os.File
}

func (l OnDiskLedger) appendToFile(data string) {
  _, err := l.LedgerFile.WriteString(data + "\n");

  if err != nil { panic(err) }
}

func (l OnDiskLedger) AppendRecord(data string) {
  l.appendToFile(data)
}

func (l OnDiskLedger) ReplayRecords(f replayFn) {
  l.LedgerFile.Seek(0, 0)
  scanner := bufio.NewScanner(l.LedgerFile)

  for scanner.Scan() {
    f(scanner.Text())
  }
}
