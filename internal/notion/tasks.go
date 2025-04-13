package notion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Task struct {
	Name     string
	DueDate  *time.Time
	Priority string
	Status   string
}

func (c *Client) CreateTask(task Task) error {
	notionURL := c.baseURL + "/pages"

	properties := map[string]interface{}{}

	properties["Name"] = map[string]interface{}{
		"type": "title",
		"title": []map[string]interface{}{
			{
				"type": "text",
				"text": map[string]interface{}{
					"content": titleCase(task.Name),
				},
			},
		},
	}
	if task.DueDate != nil {
		properties["Due Date"] = map[string]interface{}{
			"type": "date",
			"date": map[string]interface{}{
				"start": task.DueDate.Format("2006-01-02"),
			},
		}
	}
	if task.Priority != "" {
		priority := strings.ToLower(task.Priority)
		if len(priority) > 0 {
			priority = strings.ToUpper(priority[:1]) + priority[1:]
		}
		properties["Priority"] = map[string]interface{}{
			"type": "select",
			"select": map[string]interface{}{
				"name": priority,
			},
		}
	}
	if task.Status != "" {
		properties["Status"] = map[string]interface{}{
			"type": "select",
			"select": map[string]interface{}{
				"name": task.Status,
			},
		}
	}
	payload := map[string]interface{}{
		"parent": map[string]interface{}{
			"database_id": c.databaseID,
		},
		"properties": properties,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	req, err := http.NewRequest("POST", notionURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Notion-Version", "2022-06-28")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if http.StatusOK != resp.StatusCode {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error creating task: %s", body)
	}

	return nil
}

func titleCase(s string) string  {
	words := strings.Fields(s)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(words, " ")
}
