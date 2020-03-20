package babajka

import (
  "fmt"
  "log"
)

// Client is an entry point for Babajka package.
type Client struct {
  config *SecretConfig
}

// NewClient ..
func NewClient(secretConfigPath string) (*Client, error) {
  config, err := readSecretConfig(secretConfigPath)
  if err != nil {
    return nil, fmt.Errorf("failed to read secret config: %v", err)
  }
  return &Client{config}, nil
}

// UpdateAnalytics ...
func (cl *Client) UpdateAnalytics() error {
  metrics, err := cl.GetContentMetrics()
  if err != nil {
    return err
  }
  count, err := cl.pushMetricsToDB(metrics)
  if err != nil {
    return err
  }
  log.Printf("Analytics: %v records successfully pushed to DB\n", count)
  return nil
}
