package main

import (
  "fmt"
  "github.com/fatih/color"
  //"os"
)

//var count int = 0

type task struct {
  name, status string
}

var tasks = make([]*task, 0)

func addTask(name string) {
  //Create task struct.

  t := task{
    name,
    "pending",
  }
  tasks = append(tasks, &t)
  //count++
}

func deleteTask(id int) {
  //delete(tasks, id) 
  if id >= len(tasks) {
    //TODO: Throw and handle some error.
    return
  }
  ret := make([]*task, 0)
  ret = append(ret, tasks[:id]...)
  tasks = append(ret, tasks[id+1:]...)
}

func modifyStatus(id int, newStatus string) {
  if newStatus != "completed" && newStatus != "pending" {
    fmt.Println("Unrecognized status")
    return
  }
  (tasks[id]).status = newStatus
}

func listTasks(filter ...string) {
  c := color.New(color.FgCyan, color.Bold)
  c.Printf("%-3s ------>    %-20s --- %10s\n", "ID", "Task", "Status")
  if len(filter) == 0 {
    for i, v := range(tasks) {
      fmt.Printf("%-3d ------>    %-20s --- %10s\n", i, v.name, v.status)
    }
  } else {
    if filter[0] == "pending" {
      for i, v := range(tasks) {
        if v.status == "pending" {
          fmt.Printf("%-3d ------>    %-20s --- %10s\n", i, v.name, v.status)
        }
      }
    } else if filter[0] == "completed" {
      for i, v := range(tasks) {
        if v.status == "completed" {
          fmt.Printf("%-3d ------>    %-20s --- %10s\n", i, v.name, v.status)
        }
      }
    }
    fmt.Println("filter applied.")
  }
}

func main() {
  fmt.Println("Welcome to your favorite task manager")

  addTask("Washing Clothes")
  addTask("Ironing clothes")
  //addTask("Writing Codes.")
  modifyStatus(0, "completed")
  deleteTask(1)
  addTask("Writing Codes.")
  listTasks()
}
