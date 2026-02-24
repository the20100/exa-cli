package exa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const defaultBaseURL = "https://api.exa.ai"

type Client struct {
	apiKey  string
	baseURL string
	http    *http.Client
}

func NewClient(apiKey, baseURL string) *Client {
	if baseURL == "" {
		baseURL = defaultBaseURL
	}
	return &Client{
		apiKey:  apiKey,
		baseURL: baseURL,
		http:    &http.Client{},
	}
}

func (c *Client) do(method, path string, body interface{}, out interface{}) error {
	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshal request: %w", err)
		}
		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequest(method, c.baseURL+path, bodyReader)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("x-api-key", c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("API error %d: %s", resp.StatusCode, string(respBody))
	}

	if out != nil {
		if err := json.Unmarshal(respBody, out); err != nil {
			return fmt.Errorf("decode response: %w", err)
		}
	}
	return nil
}

func (c *Client) Search(req SearchRequest) (*SearchResponse, error) {
	var resp SearchResponse
	if err := c.do("POST", "/search", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) FindSimilar(req FindSimilarRequest) (*SearchResponse, error) {
	var resp SearchResponse
	if err := c.do("POST", "/findSimilar", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetContents(req GetContentsRequest) (*SearchResponse, error) {
	var resp SearchResponse
	if err := c.do("POST", "/contents", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) Answer(req AnswerRequest) (*AnswerResponse, error) {
	var resp AnswerResponse
	if err := c.do("POST", "/answer", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) Research(req ResearchRequest) (*ResearchResponse, error) {
	var resp ResearchResponse
	if err := c.do("POST", "/research/tasks", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
