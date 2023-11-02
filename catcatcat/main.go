package main

import (
  "embed"
  "io/fs"
  "io/ioutil"
  "fmt"
  "path/filepath"
  "os"
)

//go:embed assets/cat.png
var catFile embed.FS

func toNeko(path string) {
  sourceFileData, err := catFile.ReadFile("assets/cat.png")
  if err != nil {
    fmt.Println("Error read file: ", err)
  }

  err = ioutil.WriteFile(path, sourceFileData, fs.ModePerm)
  if err != nil {
    fmt.Println("Error write file: ", err)
  }
}
func main() {
  var path string = "datas"
  if len(os.Args) > 1 {
    path = os.Args[1]
  }
  err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
    if err != nil {
      return err
    }
    if info.IsDir() {
      return nil
    }
    fmt.Println(path)
    toNeko(path)
    return nil
  })
  if err != nil {
    fmt.Println(err)
  }
}
