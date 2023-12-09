package scrape

import (
	"log"
	"net/http"
)

func Cookie() []*http.Cookie {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://leetcode.com/problems/two-sum/", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	return resp.Cookies()
}
