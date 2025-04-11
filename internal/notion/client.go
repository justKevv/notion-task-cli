package notion

import (
	"fmt"
	"net/http"
	"time"

	"github.com/justKevv/notion-task-cli/internal/config"
)

type Client struct {
	apiKey    string
	baseURL   string
	httpClient *http.Client
}

func NewClient() (*Client, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return &Client{
		apiKey: cfg.NotionToken,
		baseURL: "https://api.notion.com/v1",
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}, nil
}
