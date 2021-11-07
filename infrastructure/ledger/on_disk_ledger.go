package ledger

import (
	"bufio"
	"fmt"
	// "io/ioutil"
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
  fmt.Println("CALL REPLAY")
  l.LedgerFile.Seek(0, 0)
  scanner := bufio.NewScanner(l.LedgerFile)

  for scanner.Scan() {
    fmt.Println("CALL SCAN")
    text := scanner.Text()
    fmt.Println("Output" + text)
    f(text)
  }
}
