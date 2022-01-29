package main

import (
  "fmt"
  "os/exec"
)

func main() {

  command := "uptime"
  output, error := exec.Command(command).Output()
  if error != nil {
    fmt.Println("Received error while executing the command")
  } else {
    fmt.Printf("Command %s output %s ", command, output)
  }
}
