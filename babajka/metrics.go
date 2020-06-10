package babajka

import (
	"fmt"
	"log"
	"net/url"
	"regexp"

	"github.com/babajka/babajka-analytics/yametrica"
)

const (
	contentPathRgxp = `/(?P<Locale>ru|be|en)/article/(?P<Slug>[a-zA-Z0-9-_]*)/`
)

// LocalizedMetric maps locale to number of views
type LocalizedMetric map[string]int

// Metrics maps content content slug to localized numbers
type Metrics map[string]LocalizedMetric

// GetContentMetrics returns pageviews for all Wir.by content.
func (cl *Client) GetContentMetrics() (Metrics, error) {
	yandexConfig := cl.config.Services.Yandex
	ymClient := yametrica.NewClient(yandexConfig.ProjectID, yandexConfig.LaunchDate, yandexConfig.AuthKey)
	pv, err := ymClient.GetPageviews()
	if err != nil {
		return nil, fmt.Errorf("failed to get data from Yandex: %v", err)
	}
	r := regexp.MustCompile(contentPathRgxp)
	articles := make(Metrics)
	for rawurl, views := range *pv {
		url, err := url.Parse(rawurl)
		if err != nil {
			log.Printf("failed to parse URL '%s': %v\n", rawurl, err)
			continue
		}
		path := url.RequestURI()
		match := r.FindStringSubmatch(path)
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
