package main

import (
  "bufio"
  "errors"
  "fmt"
  "os"
  "os/exec"
  "strings"
)

func main() {
  fmt.Printf("Hello Shell User! \n")

  reader := bufio.NewReader(os.Stdin)

  for {
    fmt.Print("$ ")

    input, err := reader.ReadString('\n')

    if err != nil {
      fmt.Fprintln(os.Stderr, err)
    }

    if err = execInput(input); err != nil {
      fmt.Fprintln(os.Stderr, err)
    }

  }
}


func execInput(input string) error{
  input = strings.TrimSuffix(input, "\n")
  input = strings.TrimSpace(input)

  args := strings.Split(input, " ")

  switch args[0] {
  case "cd":
    if len(args) < 2{
      return errors.New("Path Required!")
    }

    return os.Chdir(args[1])

  case "exit":
    os.Exit(0)
}

  cmd := exec.Command(args[0], args[1:]...)

  cmd.Stderr = os.Stderr
  cmd.Stdout = os.Stdout

  return cmd.Run()
}
