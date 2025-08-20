package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"task_tracker/task_cli/models"
)

var taskStore *models.TaskStore

var rootCmd = &cobra.Command{
	Use: "task-cli",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	getConfigPath, err := models.GetConfigPath()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Khởi tạo task store với file tasks.json
	taskStore = models.NewTaskStore(getConfigPath)

	// Load dữ liệu từ file
	if err := taskStore.Load(); err != nil {
		panic(err)
	}

	rootCmd.AddCommand(addCmd, updateCmd, deleteCmd, markTodoCmd, markInProgressCmd, markDoneCmd,
		listAllCmd)
}
