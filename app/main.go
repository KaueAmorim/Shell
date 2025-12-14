package main

import (
	"bufio"
	"fmt"
	"os"
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
				fmt.Printf("%s: not found\n", args[1])
			}
		default:
			fmt.Println(command + ": command not found")
		}
	}
}
