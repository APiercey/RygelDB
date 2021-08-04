package commands

import (
  "encoding/json"
  "example.com/kv_store/store" 
)

type LookupCommand struct {
  collectionName string
  key string
}

func (c LookupCommand) Execute(s *store.Store) (string, bool) {
  item, presence := s.Collections[c.collectionName].ReadByKey(c.key)

  if presence {
    out, err := json.Marshal(item.Data)

    if err != nil { panic (err) }

    return string(out), false
  } else {
    return "", false
  }
}
