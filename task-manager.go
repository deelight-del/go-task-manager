package main

import (
  "os"
  "fmt"
  "github.com/deelight-del/go-task-manager/internal/utils"
)

func main() {
  fmt.Println("Welcome to your favorite task manager")

  utils.AddTask("Washing Clothes")
  utils.AddTask("Ironing clothes")
  //utils.ModifyStatus(0, "completed")

  if len(os.Args) < 2 {
    // TODO:Print to std error.
    // Or use fatih coloring
    fmt.Println("Specify a command.Add, list or something at all.")
    return
  }

  cmd := os.Args[1]
  args := os.Args[2:]

  switch cmd {
  case "add":
    addCmd(args)
  case "list":
    listCmd(args)
  case "complete":
    completeCmd(args)
  case "delete":
    deleteCmd(args)
  default:
    //TODO: handle unknown command.
    fmt.Println("Unknown command.")
  }
  //addTask("Writing Codes.")
  //utils.DeleteTask(1)
  //utils.AddTask("Writing Codes.")
  //completeCmd([]string{"0"})
  utils.ListTasks()
}
