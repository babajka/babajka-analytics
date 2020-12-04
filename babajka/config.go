package babajka

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type dbConfigOptions struct {
	DBName string `json:"dbName"`
}

type dbConfig struct {
	URL     string          `json:"url"`
	Options dbConfigOptions `json:"options"`
}

type yandexConfig struct {
	AuthKey    string `json:"authKey"`
	ProjectID  string `json:"projectID"`
	LaunchDate string `json:"launchDate"`
}

type slackConfig struct {
	APIToken    string `json:"apiToken"`
	ChannelName string `json:"channelName"`
}

type servicesConfig struct {
	Yandex            yandexConfig `json:"yandex"`
	SlackAnalyticsApp slackConfig  `json:"slackAnalyticsApp"`
}

// SecretConfig ..
type SecretConfig struct {
	Mongodb  dbConfig       `json:"mongodb"`
	Services servicesConfig `json:"services"`
}

// readSecretConfig ..
func readSecretConfig(filePath string) (*SecretConfig, error) {
	jsonConfig, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read secret file: %v", err)
	}
	defer jsonConfig.Close()

	buf, err := ioutil.ReadAll(jsonConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to read from json config: %v", err)
	}

	config := &SecretConfig{}
	err = json.Unmarshal(buf, config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal secret file: %v", err)
	}

	return config, nil
}
