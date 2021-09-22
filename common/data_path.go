package common

import "strings"

type DataPath struct {
  RealPath []string
}

func (dp DataPath) SerializedPath() string {
  return strings.Join(dp.RealPath[:],",")
}

func (dp DataPath) Steps() []string {
  return dp.RealPath[:len(dp.RealPath) - 1] 
}

func (dp DataPath) Key() string {
  return dp.RealPath[len(dp.RealPath) - 1] 
}
