package babajka

import (
	"fmt"
	"regexp"

	"github.com/babajka/babajka-analytics/yametrica"
)

var contentPathRE = regexp.MustCompile(`/(?P<Locale>ru|be|en)/article/(?P<Slug>[a-zA-Z0-9-_]*)/`)

// LocalizedMetric maps locale to number of views.
type LocalizedMetric map[string]int

// Metrics maps content content slug to localized numbers.
type Metrics map[string]LocalizedMetric

// GetContentMetrics returns pageviews for all Wir.by content.
func (cl *Client) GetContentMetrics() (Metrics, error) {
	yandexConfig := cl.config.Services.Yandex
	ymClient := yametrica.NewClient(yandexConfig.ProjectID, yandexConfig.LaunchDate, yandexConfig.AuthKey)
	pv, err := ymClient.GetPageviews()
	if err != nil {
		return nil, fmt.Errorf("failed to get data from Yandex: %v", err)
	}
	articles := make(Metrics)
	for rawURL, views := range pv {
		match := contentPathRE.FindStringSubmatch(rawURL)
		if match == nil {
			continue
		}
		locale, slug := match[1], match[2]
		if _, ok := articles[slug]; !ok {
			articles[slug] = make(LocalizedMetric)
		}
		articles[slug][locale] += views
	}
	return articles, nil
}
