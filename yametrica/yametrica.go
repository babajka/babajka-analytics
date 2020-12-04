package yametrica

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	queryString = "https://api-metrika.yandex.net/stat/v1/data?dimensions=%s&metrics=%s&id=%s&date1=%s&limit=%s"
	dimensions  = "ym:pv:title,ym:pv:URL"
	metrics     = "ym:pv:pageviews"
	limit       = "100000"
)

// Pageviews per URL
type Pageviews map[string]int

type ymDimension struct {
	Name string `json:"name"`
}

type ymDataRow struct {
	Dimensions []ymDimension `json:"dimensions"`
	Metrics    []float64     `json:"metrics"`
}

type ymResponse struct {
	// Query interface{}
	Data []ymDataRow `json:"data"`
}

// Client is a simple client for Yandex.Metrica
type Client struct {
	projectID string
	date1     string
	authKey   string
}

// NewClient returns a new client for Yandex.Metrica
func NewClient(projectID, date1, authKey string) *Client {
	return &Client{
		projectID: projectID,
		date1:     date1,
		authKey:   authKey,
	}
}

// GetPageviews returns pageviews for all Wir.by content.
func (ym *Client) GetPageviews() (Pageviews, error) {
	resp, err := ym.makeRequest()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	var ymResp ymResponse
	if err := json.Unmarshal(body, &ymResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %v", err)
	}

	pv := make(Pageviews)
	for _, dataRow := range ymResp.Data {
		pv[dataRow.Dimensions[1].Name] += int(dataRow.Metrics[0])
	}

	return pv, nil
}

func (ym *Client) makeRequest() (*http.Response, error) {
	url := fmt.Sprintf(queryString, dimensions, metrics, ym.projectID, ym.date1, limit)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("OAuth %s", ym.authKey))

	return http.DefaultClient.Do(req)
}
