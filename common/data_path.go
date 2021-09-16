package common

import "strings"

type DataPath struct {
  RealPath []string
}

func (dp DataPath) SerializedPath() string {
  return strings.Join(dp.RealPath[:],",")
}

