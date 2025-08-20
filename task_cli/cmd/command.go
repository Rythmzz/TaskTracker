package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"task_tracker/task_cli/models"
)

var listAllCmd = &cobra.Command{
	Use:  "list [status]",
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var tasksToShow []models.Task

		if len(args) > 0 {
			status := args[0]
			tasksToShow = make([]models.Task, 0)

			for _, t := range taskStore.Tasks {
				if string(t.Status) == status {
					tasksToShow = append(tasksToShow, t)
				}
			}
		} else {
			tasksToShow = taskStore.Tasks
		}

		taskStore.PrintTask(tasksToShow)
	},
}

var markTodoCmd = &cobra.Command{
	Use:  "mark-todo [id]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		idConvert, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			fmt.Println("parse id err:", err)
			return
		}
		taskStore.UpdateStatusTask(idConvert, 0)

		err = taskStore.Save()
		if err != nil {
			fmt.Println("save task store err:", err)
			return
		}
	},
}

var markInProgressCmd = &cobra.Command{
	Use:  "mark-in-progress [id]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		idConvert, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			fmt.Println("parse id err:", err)
			return
		}
		taskStore.UpdateStatusTask(idConvert, 1)

		err = taskStore.Save()
		if err != nil {
			fmt.Println("save task store err:", err)
			return
		}
	},
}

var markDoneCmd = &cobra.Command{
	Use:  "mark-done [id]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		idConvert, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			fmt.Println("parse id err:", err)
			return
		}
		taskStore.UpdateStatusTask(idConvert, 2)

		err = taskStore.Save()
		if err != nil {
			fmt.Println("save task store err:", err)
			return
		}
	},
}

var deleteCmd = &cobra.Command{
	Use:  "delete [id]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		idConvert, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			fmt.Println("parse id err:", err)
			return
		}
		newTasks := taskStore.DeleteTask(idConvert)
		taskStore.Tasks = newTasks

		err = taskStore.Save()
		if err != nil {
			fmt.Println("save task store err:", err)
			return
		}
	},
}

var updateCmd = &cobra.Command{
	Use:  "update [id] [description]",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		description := args[1]
		idConvert, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			fmt.Println("parse id err:", err)
			return
		}
		taskStore.UpdateTask(idConvert, description)

		err = taskStore.Save()
		if err != nil {
			fmt.Println("save task store err:", err)
			return
		}
	},
}

var addCmd = &cobra.Command{
	Use:  "add [description]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description := args[0]
		task, err := taskStore.AddTask(description)
		if err != nil {
			fmt.Println("task add err:", err)
			return
		}
		err = taskStore.Save()
		if err != nil {
			fmt.Println("save task store err:", err)
			return
		}
		addNotifyTask(task.ID)

	},
}

func addNotifyTask(id int64) {
	s := fmt.Sprintf("Task added successfully (ID:%d)", id)
	fmt.Println(s)
}
