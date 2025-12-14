package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	for {
		fmt.Print("$ ")

		reader := bufio.NewReader(os.Stdin)
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		command = strings.TrimSpace(command)

		// Get the PATH environment variable and split it into directories
		pathEnv := os.Getenv("PATH")
		directories := strings.Split(pathEnv, string(os.PathListSeparator))

		args := strings.Fields(command)
		switch args[0] {
		case "exit":
			return
		case "echo":
			if len(args) > 1 {
				fmt.Println(strings.Join(args[1:], " "))
			} else {
				fmt.Println()
			}
		case "type":
			if args[1] == "echo" || args[1] == "type" || args[1] == "exit" {
				fmt.Printf("%s is a shell builtin\n", args[1])
			} else {
				arquive := args[1]
				found := false
				for _, dir := range directories {
					completePath := filepath.Join(dir, arquive)
					info, err := os.Stat(completePath)
					if err == nil {
						mode := info.Mode()
						// 0111 in octal checks if Owner(1) OR Group(1) OR Others(1) have execute permission.
						if mode.IsRegular() && mode.Perm()&0111 != 0 {
							fmt.Printf("%s is %s\n", arquive, completePath)
							found = true
							break
						}
					}
				}
				if !found {
					fmt.Printf("%s: not found\n", arquive)
				}
			}
		default:
			program := args[0]
			caminho, err := exec.LookPath(program)
			if err != nil {
				fmt.Println(command + ": command not found")
			} else {
				args := args[1:]
				cmd := exec.Command(caminho, args...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Stdin = os.Stdin
				cmd.Run()
			}
		}
	}
}
