package storage

type Storage interface {
  LogCommand(rawStatement string)
  LoadData()
  CreateSnapshot()
}
