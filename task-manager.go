package main

import (
  //"os"
  "fmt"
  "github.com/deelight-del/go-task-manager/internal/utils"
)

func main() {
  fmt.Println("Welcome to your favorite task manager")

  utils.AddTask("Washing Clothes")
  utils.AddTask("Ironing clothes")
  //addTask("Writing Codes.")
  utils.ModifyStatus(0, "completed")
  utils.DeleteTask(1)
  utils.AddTask("Writing Codes.")
  utils.ListTasks()
}
