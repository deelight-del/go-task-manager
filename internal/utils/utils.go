package utils

import (
  "fmt"
  "github.com/fatih/color"
  "os"
  "encoding/json"
)

//var count int = 0

type task struct {
  Name, Status string
}

var tasks = make([]*task, 0)

func AddTask(name string) {
  //Create task struct.

  t := task{
    name,
    "pending",
  }
  tasks = append(tasks, &t)
  //count++
}

func DeleteTask(id int) {
  //delete(tasks, id) 
  if id >= len(tasks) {
    //TODO: Throw and handle some error.
    return
  }
  ret := make([]*task, 0)
  ret = append(ret, tasks[:id]...)
  tasks = append(ret, tasks[id+1:]...)
}

func ModifyStatus(id int, newStatus string) {
  if newStatus != "completed" && newStatus != "pending" {
    fmt.Println("Unrecognized status")
    return
  }

  if id >= len(tasks) {
    //TODO: Throw and handle some error.
    return
  }

  (tasks[id]).Status = newStatus
}

func ListTasks(filter ...string) {
  c := color.New(color.FgCyan, color.Bold)

  c.Printf("%-3s ------>    %-20s --- %10s\n", "ID", "Task", "Status")

  if len(filter) == 0 {
    for i, v := range(tasks) {
      fmt.Printf("%-3d ------>    %-20s --- %10s\n", i, v.Name, v.Status)
    }
  } else {
    if filter[0] == "pending" {
      for i, v := range(tasks) {
        if v.Status == "pending" {
          fmt.Printf("%-3d ------>    %-20s --- %10s\n", i, v.Name, v.Status)
        }
      }
    } else if filter[0] == "completed" {
      for i, v := range(tasks) {
        if v.Status == "completed" {
          fmt.Printf("%-3d ------>    %-20s --- %10s\n", i, v.Name, v.Status)
        }
      }
    } else {
      fmt.Println("Unrecognized filter. Use pending or completed")
    }
    //fmt.Println("filter applied.")
  }
}

func SaveTasks() error {
  f, err := os.Create("save.json")
  if err != nil {
    return err
  }
  defer f.Close()

  if err := json.NewEncoder(f).Encode(&tasks); err != nil {
    return err
  }

  return nil
}

func LoadTasks() error {
  f, err := os.Open("save.json")
  if err != nil {
    if os.IsNotExist(err) {
      fmt.Println("No file is saved at the moment.")
      return nil
    }
    return err
  }
  defer f.Close()

  if err := json.NewDecoder(f).Decode(&tasks); err != nil {
    return err
  }
  return nil
}
