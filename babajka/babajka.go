package babajka

import (
	"fmt"
	"log"
)

var (
	validEnvs = []string{"dev", "staging", "production"}
)

// Client is an entry point for Babajka package.
type Client struct {
	config *SecretConfig
	env    string
}

// NewClient ..
func NewClient(secretConfigPath, env string) (*Client, error) {
	config, err := readSecretConfig(secretConfigPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read secret config: %v", err)
	}
	validEnv := false
	for _, e := range validEnvs {
		if e == env {
			validEnv = true
			break
		}
	}
	if !validEnv {
		return nil, fmt.Errorf("bad env provided: '%v'", env)
	}
	return &Client{config, env}, nil
}

// UpdateAnalytics ...
func (cl *Client) UpdateAnalytics() error {
	metrics, err := cl.GetContentMetrics()
	if err != nil {
		return err
	}
	countDocuments, totalMetrics, err := cl.pushMetricsToDB(metrics)
	if err != nil {
		return err
	}
	result := fmt.Sprintf(`ANALYTICS REPORT
    Env: %v
    Records (unique article slugs) pushed: %v
    Total metrics pageviews: %v`, cl.env, countDocuments, totalMetrics)
	log.Printf(result)
	slackConfig := cl.config.Services.SlackAnalyticsApp
	err = pushSlackNotification(slackConfig.APIToken, slackConfig.ChannelName, result)
	if err != nil {
		return err
	}
	return nil
}
