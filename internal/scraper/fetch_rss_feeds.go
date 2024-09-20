package scraper

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

func FetchFeedData(client http.Client, url string) (RSS, error) {
	resp, err := client.Get(url)
	if err != nil {
		return RSS{}, fmt.Errorf("failed to fetch feed data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return RSS{}, fmt.Errorf("failed to fetch feed data: %s", resp.Status)
	}

	var rss RSS
	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&rss)
	if err != nil {
		return RSS{}, fmt.Errorf("failed to decode response body: %w", err)
	}

	return rss, nil
}
