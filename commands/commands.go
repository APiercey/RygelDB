package commands

type Command interface {
  Execute() (result string, store_was_updated bool)
  Valid() bool
}
