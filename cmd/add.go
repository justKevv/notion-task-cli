package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/justKevv/notion-task-cli/internal/chat"
	"github.com/justKevv/notion-task-cli/internal/notion"
)

var addCmd = &cobra.Command{
	Use: "add [task details...]",
	Short: "Add a new task to notion using natural language",
	Long: `Adds a new task to the configured Notion database using natural language keywords.
	Example: notion-task add Finish the report due next friday priority high status todo`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fullInput := "add " + strings.Join(args, " ")

		task, err  := chat.ParseTaskInput(fullInput)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing task input: %v\n", err)
			os.Exit(1)
		}

		notionClient, err := notion.NewClient()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error initializing Notion client: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Adding task '%s' to Notion...\n", task.Name)
		err = notionClient.CreateTask(task)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error adding task to Notion: %v\n", err)
			os.Exit(1) // Exit with error on failure
		}

		fmt.Println("Task Sucesfully added!")
	},
}
