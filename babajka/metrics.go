package babajka

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

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
	pv, err := ymClient.GetPageViews()
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

	if ok := checkMetricsHealthy(articles); !ok {
		message := fmt.Sprintf("💥 ALERT 💥  *CORRUPTED DATA* (probably)  ENV %s", cl.env)
		message += fmt.Sprintf("\n%20sbe, ru, en interface localization\n", "")

		rows := make([]string, 0, len(articles))
		for slug, article := range articles {
			row := fmt.Sprintf("`%-30s`", slug)
			for _, locale := range []string{"be", "ru", "en"} {
				row += fmt.Sprintf("%d  ", article[locale])
			}
			rows = append(rows, row)
		}
		sort.Strings(rows)
		message += strings.Join(rows, "\n")

		slackConfig := cl.config.Services.SlackAnalyticsApp
		if err := pushSlackNotification(slackConfig.APIToken, slackConfig.ChannelName, message); err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("data from yandex looks corrupted. slack report sent")
	}

	return articles, nil
}

// checkMetricsHealthy checks for metrics being valid. This is to produce alerts and debug occurring issues
// with broken data on babajka admin dashboard.
func checkMetricsHealthy(articles Metrics) bool {
	// Issue type: way too many zeroes in output.
	countEmpty := 0
	for _, article := range articles {
		for _, localeMetric := range article {
			if localeMetric == 0 {
				countEmpty++
			}
		}
	}
	if countEmpty > len(articles)/4 {
		// At least 25% data is corrupted.
		return false
	}

	// Issue type: data gets rounded.
	countRounded := 0
	for _, article := range articles {
		for _, localeMetric := range article {
			if localeMetric%100 == 0 && localeMetric > 100 {
				countRounded++
			}
		}
	}
	if countRounded > 10 {
		return false
	}

	return true
}
