package main

import (
  "fmt"
  "github.com/faith/color"
)

var count int = 0

type task struct {
  name, status string
}

type taskMap map[int]*task

var tasks taskMap = make(taskMap)

func addTask(name string) {
  //Create task struct.

  t := task{
    name,
    "pending",
  }
  tasks[count] = &t
  count++
}

func deleteTask(id int) {
  delete(tasks, id)
}

func modifyStatus(id int, newStatus string) {
  (tasks[id]).status = newStatus
}

func listTasks() {
  c := color.New(color.FgCyan, color.Bold)
  c.Printf("%-3s ------>    %-20s --- %10s\n", "ID", "Task", "Status")
  for k, v := range(tasks) {
    fmt.Printf("%-3d ------>    %-20s --- %10s\n", k, v.name, v.status)
  }
}

func main() {
  fmt.Println("Welcome to your favorite task manager")

  addTask("Washing Clothes")
  addTask("Ironing clothes")
  addTask("Writing Codes.")
  modifyStatus(0, "done")
  //deleteTask(1)
  listTasks()
}
