package main

import (
	"fmt"
	"os"
	"task_tracker/task_cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
