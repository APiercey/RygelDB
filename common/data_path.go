package common

import (
	"fmt"
	"strings"
)

type DataPath struct {
  RealPath []string 
}

func (dp DataPath) SerializedPath() string {
  return strings.Join(dp.RealPath[:],",")
}

func (dp DataPath) Steps() []string {
  fmt.Println("DP IS HERE")
  fmt.Println(dp)
  return dp.RealPath[:len(dp.RealPath) - 1] 
}

func (dp DataPath) Key() string {
  fmt.Println(dp)
  return dp.RealPath[len(dp.RealPath) - 1] 
}
