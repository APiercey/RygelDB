package commands

import (
	"rygel/core"
)

type noopErrorCommand struct {
  err string
}

func (c noopErrorCommand) RawStatement() string {
  return "";
}

func (c noopErrorCommand) Execute(s *core.Store) (string, bool) {
  return c.err, false
}

func (c noopErrorCommand) Valid() bool {
  return true
}

