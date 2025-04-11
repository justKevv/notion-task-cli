package chat

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/justKevv/notion-task-cli/internal/notion"
)

func StartChatMode(notionClient *notion.Client)  {
	fmt.Println("Entering Notion Task Chat Mode...")
	fmt.Println("Type 'exit' to quit.")
	fmt.Println("Example: add Finish report due tomorrow priority high status todo")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Println("Error reading input:", err)
			}
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if strings.ToLower(input) == "exit" {
			fmt.Println("Exiting Notion Task Chat Mode...")
			break
		}

		if strings.HasPrefix(strings.ToLower(input), "add ") {
			task, err := parseTaskInput(input)
			if err != nil {
				fmt.Println("Error parsing task:", err)
				continue
			}

			fmt.Println("Adding task to notion...")
			err = notionClient.CreateTask(task)
			if err!= nil {
				fmt.Println("Error adding task to notion:", err)
			} else {
				fmt.Println("Task added successfully!")
			}
		} else if input != "" {
			fmt.Println("Unknown command. Try starting with 'add'.")
		}
	}

	if err := scanner.Err(); err != nil  {
		fmt.Fprintln(os.Stderr, "Error reading from scanner:", err)
	}
}

func parseTaskInput(input string) (notion.Task, error) {
	// TODO: Implement parsing logic here
	// - Extract task name
	// - Look for keywords like "due", "priority", "status"
	// - Parse dates/times
	// - Return a populated notion.Task struct or an error
	task := strings.TrimSpace(strings.TrimPrefix(input, "add "))

	if task == "" {
		return notion.Task{}, fmt.Errorf("invalid task input")
	}
	taskDetails := map[string]string{}

	words := strings.Fields(task)

	var taskName []string

	keywords := []string{"due", "priority", "status"}

	i := 0
	for i < len(words) {
		word := strings.ToLower(words[i])

		isKeyword := false
		for _, keyword := range keywords {
			if word == keyword {
				isKeyword = true

				if i+1 < len(words) {
					taskDetails[word] = words[i+1]
					i += 2
				} else {
					i++
				}
				break
			}
		}

		if !isKeyword {
			taskName = append(taskName, words[i])
			i++
		}
	}

	name := strings.Join(taskName, " ")

	newTask := notion.Task{
		Name: name,
	}

	if dueDate, ok := taskDetails["due"]; ok {
		parsedDate, err := parseDueDate(dueDate)
		if err != nil {
			fmt.Println("Warning: Could not parse due date:", err)
		} else {
			newTask.DueDate = &parsedDate
		}
	}

	if priority, ok := taskDetails["priority"]; ok {
		newTask.Priority = priority
	}

	if status, ok := taskDetails["status"]; ok {
		newTask.Status = status
	}

	return newTask, nil
}

func parseDueDate(dateStr string) (time.Time, error)  {
	now := time.Now()
	dateStr = strings.ToLower(dateStr)

	switch dateStr {
	case "today":
		return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()), nil
	case "tomorrow":
		return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).AddDate(0, 0, 1), nil
	case "yesterday":
		return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).AddDate(0, 0, -1), nil
	}

	weekdays := map[string]time.Weekday{
		"sunday":    time.Sunday,
		"monday":    time.Monday,
		"tuesday":   time.Tuesday,
		"wednesday": time.Wednesday,
		"thursday":  time.Thursday,
		"friday":    time.Friday,
		"saturday":  time.Saturday,
	}

	if weekday, ok := weekdays[dateStr]; ok {
		daysUntil := int(weekday - now.Weekday())
		if daysUntil <= 0 {
			daysUntil += 7
		}
		return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).AddDate(0, 0, daysUntil), nil
	}

	for weekdayName, weekday := range weekdays {
		if dateStr == "next "+weekdayName {
			// Find the next occurrence of this weekday and add 7 days
			daysUntil := int(weekday - now.Weekday())
			if daysUntil <= 0 {
				daysUntil += 7
			}
			return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).AddDate(0, 0, daysUntil+7), nil
		}
	}

	if strings.HasSuffix(dateStr, " days from now") {
		daysStr := strings.TrimSuffix(dateStr, " days from now")
		days, err := strconv.Atoi(daysStr)
		if err == nil {
			return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).AddDate(0, 0, days), nil
		}
	}

	formats := []string {
		"2006-01-02",
		"01/02/2006",
		"Jan 2, 2006",
		"January 2, 2006",
	}

	for _, format := range formats {
		if t, err := time.Parse(format, dateStr); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("could not parse date: %s", dateStr)
}
