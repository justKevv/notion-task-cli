package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/justKevv/notion-task-cli/internal/chat"
	"github.com/justKevv/notion-task-cli/internal/notion"
)

var rootCmd = &cobra.Command{
	Use: "notion-task",
	Short: "A CLI tool to add tasks to Notion via chat",
	Long: `notion-task is a command-line interface tool that allows you
to quickly add tasks to a specified Notion database through an
interactive chat mode.`,
}

var chatCmd = &cobra.Command{
	Use: "chat",
	Short: "Enter interactive chat mode to add Notion tasks",
	Long:  `Starts an interactive session where you can type commands like 'add [task details]' to add tasks to your Notion database.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initializing the Notion Chat...")
		notionClient, err := notion.NewClient()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error initializing Notion client: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Client Initialize.")
		chat.StartChatMode(notionClient)
	},
}

func Execute()  {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init()  {
	rootCmd.AddCommand(chatCmd)
	rootCmd.AddCommand(addCmd)
}
