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
	config      *SecretConfig
	env         string
	enableSlack bool
}

// NewClient ..
func NewClient(secretConfigPath, env string, enableSlack bool) (*Client, error) {
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
	return &Client{config, env, enableSlack}, nil
}

// UpdateAnalytics updates analytics in Babajka DB.
func (cl *Client) UpdateAnalytics() (Metrics, error) {
	metrics, err := cl.GetContentMetrics()
	if err != nil {
		if cl.enableSlack {
			slackConfig := cl.config.Services.SlackAnalyticsApp
			text := fmt.Sprintf("Failed to perform update, error occurred: %v", err)
			slackErr := pushSlackNotification(slackConfig.APIToken, slackConfig.ChannelName, text)
			if slackErr != nil {
				log.Printf("failed to send error to slack: %v", slackErr)
			}
		}
		return nil, err
	}

	countDocuments, totalMetrics, err := cl.pushMetricsToDB(metrics)
	if err != nil {
		return nil, err
	}

	result := fmt.Sprintf(`ANALYTICS REPORT
    Env: %v
    Records (unique article slugs) pushed: %v
    Total metrics pageviews: %v`, cl.env, countDocuments, totalMetrics)
	log.Printf(result)

	if cl.enableSlack {
		slackConfig := cl.config.Services.SlackAnalyticsApp
		err = pushSlackNotification(slackConfig.APIToken, slackConfig.ChannelName, result)
		if err != nil {
			return nil, err
		}
	}

	return metrics, nil
}
