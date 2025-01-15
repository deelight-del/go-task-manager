package main

import (
  "encoding/json"
  "os"
  "fmt"
)

func main() {
  f, err := os.Create("file.json")
  if err != nil {
    fmt.Println(err)
    return
  }
  defer f.Close()
  enc := json.NewEncoder(f)
  type career struct {
    Name, Career string
  }
  task1 := career{"david", "cen"}
  if err := enc.Encode(&task1); err != nil {
    fmt.Println(err)
    return
  }
  var v career
  newF, err := os.Open("file.json")
  dec := json.NewDecoder(newF)
  if err := dec.Decode(&v); err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(v)
}
