package common

import (
  "io/ioutil"
)

func CollectPathsInDir(dir string) []string {
  dirs := []string{}

  files, err := ioutil.ReadDir(dir)

  HandleErr(err)

  for _, file := range files {
    if file.IsDir() {
      dirs = append(dirs, file.Name())
    }
  }
  
  return dirs
}

