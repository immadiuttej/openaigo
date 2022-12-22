package openaigo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	HTTPClient *http.Client
	APIKey string
	Endpoint string
}

type CompletionRequest struct {
	Prompt string
	Model string
	MaxTokens int
	Temperature float64
}

type CompletionResponse struct {
	ID string
	Model string
	Prompt string
	Completions []string
	Tokens []string
}

func (c *Client) Complete(ctx context.Context, req *CompletionRequest) (*CompletionResponse, error) {
	httpReq, err := http.NewRequest("POST", c.Endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))

	// Set the request body.
	data := map[string]interface{}{
		"prompt":      req.Prompt,
		"model":       req.Model,
		"max_tokens":  req.MaxTokens,
		"temperature": req.Temperature,
	}
	body, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}
	httpReq.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	httpResp, err := c.HTTPClient.Do(httpReq.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer httpResp.Body.Close()
	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	var resp CompletionResponse
	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response body: %w", err)
	}

	return &resp, nil
}