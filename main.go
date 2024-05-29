package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func main() {
  green := color.New(color.FgGreen).Add(color.Bold)
  neonYellow := color.New(color.FgHiYellow)
  neonYellow.Print(`
 ___  ___  _______   ___       ___       ________          ________  ___  ___  _______   ___       ___               ___  ___  ________  _______   ________  ___       
|\  \|\  \|\  ___ \ |\  \     |\  \     |\   __  \        |\   ____\|\  \|\  \|\  ___ \ |\  \     |\  \             |\  \|\  \|\   ____\|\  ___ \ |\   __  \|\  \      
\ \  \\\  \ \   __/|\ \  \    \ \  \    \ \  \|\  \       \ \  \___|\ \  \\\  \ \   __/|\ \  \    \ \  \            \ \  \\\  \ \  \___|\ \   __/|\ \  \|\  \ \  \     
 \ \   __  \ \  \_|/_\ \  \    \ \  \    \ \  \\\  \       \ \_____  \ \   __  \ \  \_|/_\ \  \    \ \  \            \ \  \\\  \ \_____  \ \  \_|/_\ \   _  _\ \  \    
  \ \  \ \  \ \  \_|\ \ \  \____\ \  \____\ \  \\\  \       \|____|\  \ \  \ \  \ \  \_|\ \ \  \____\ \  \____        \ \  \\\  \|____|\  \ \  \_|\ \ \  \\  \\ \__\   
   \ \__\ \__\ \_______\ \_______\ \_______\ \_______\        ____\_\  \ \__\ \__\ \_______\ \_______\ \_______\       \ \_______\____\_\  \ \_______\ \__\\ _\\|__|   
    \|__|\|__|\|_______|\|_______|\|_______|\|_______|       |\_________\|__|\|__|\|_______|\|_______|\|_______|        \|_______|\_________\|_______|\|__|\|__|   ___ 
                                                             \|_________|                                                        \|_________|                     |\__\
                                                                                                                                                                  \|__|                                                                                                                                                                       
`)

  reader := bufio.NewReader(os.Stdin)

  for {
    green.Print("$ ")

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
