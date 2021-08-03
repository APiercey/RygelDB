package main

import (
	// "encoding/gob"
	// "fmt"
	// "os"
)

type Item struct {
  Data string `json:"data"`
}

type Store struct {
  Items map[string]Item `json:"items"`
}

func (s Store) ReadByKey(key string) (Item, bool) {
  item, presence := s.Items[key]
  return item, presence
}

func (s *Store) InsertItem(key string, data string) bool {
  s.Items[key] = Item{data}

  return true
}

