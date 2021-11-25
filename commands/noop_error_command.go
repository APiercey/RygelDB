package commands

type NoopErrorCommand struct {
  Err string
}

func (c NoopErrorCommand) Execute() (string, bool) {
  return c.Err, false
}

func (c NoopErrorCommand) Valid() bool {
  return true
}

