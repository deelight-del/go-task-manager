package main

import (
  "github.com/deelight-del/go-task-manager/internal/utils"
  "fmt"
  "strconv"
)

func addCmd(args []string) {
  if len(args) == 0 {
    fmt.Println("You need to pass at least one task to add command.")
    return
  }
  for _, arg := range args {
    utils.AddTask(arg)
  }
}

func listCmd(args []string) {
  if len(args) == 0 {
    utils.ListTasks()
  } else if len(args) == 1 {
    utils.ListTasks(args...)
  } else {
    fmt.Println("Too many arguments, chief.\nYou only need a single filter or no filter.")
  }
}

func completeCmd(args []string) {
  if len(args) < 1 {
    //TODO: Alternative is to create an error method that does some.
    //error printing.
    fmt.Println("You need to pass in at least one relevant task ID")
  } else {
    for _, arg := range args {
      taskID, err := strconv.Atoi(arg)
      if err != nil {
        fmt.Printf("Task ID \"%v\" irrelevant\n", arg)
      } else {
        utils.ModifyStatus(taskID, "completed")
      }
    }
  }
}


func deleteCmd(args []string) {
  if len(args) < 1 {
    //TODO: Alternative is to create an error method that does some.
    //error printing.
    fmt.Println("You need to pass in at least one relevant task ID")
  } else {
    for _, arg := range args {
      taskID, err := strconv.Atoi(arg)
      if err != nil {
        fmt.Printf("Task ID \"%v\" irrelevant\n", arg)
      } else {
        utils.DeleteTask(taskID)
      }
    }
  }
  //...
}
