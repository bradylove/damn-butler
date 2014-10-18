package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type Host struct {
	CCUrl string
}

func NewHost(url string) (Host, error) {
	_, err := FetchResults(url)
	if err != nil {
		return Host{}, err
	}

	return Host{CCUrl: url}, nil
}

func FetchResults(url string) (*Result, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	results := Result{}
	err = xml.Unmarshal(data, &results)
	if err != nil {
		return nil, err
	}

	return &results, nil
}

func FetchAllResults(hosts []Host) (*Result, error) {
	allResults := Result{}

	for _, h := range hosts {
		results, err := FetchResults(h.CCUrl)
		if err != nil {
			continue
		}

		allResults.Projects = append(allResults.Projects, results.Projects...)
	}

	return &allResults, nil
}
