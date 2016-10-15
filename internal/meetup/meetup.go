package meetup

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	GroupURLName string
}

func (c *Client) FetchEvents() ([]Event, error) {
	params := url.Values{}
	params.Set("group_urlname", c.GroupURLName)
	params.Set("status", "upcoming")
	params.Set("page", "15")
	params.Set("only", "name,time,utc_offset,timezone,venue.name,venue.address_1")

	var body struct {
		Events []Event `json:"results"`
	}
	err := c.apiRequest("/2/events", params, &body)
	if err != nil {
		return []Event{}, err
	}

	return body.Events, nil
}

func (c *Client) apiRequest(method string, params url.Values, v interface{}) error {
	u := url.URL{
		Scheme:   "https",
		Host:     "api.meetup.com",
		Path:     "/2/events",
		RawQuery: params.Encode(),
	}

	res, err := http.Get(u.String())
	if err != nil {
		return err
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	return nil
}

type Event struct {
	Name            string `json:"name"`
	Timestamp       int64  `json:"time"`
	TimestampOffset int    `json:"utc_offset"`
	Venue           Venue  `json:"venue""`
}

func (e *Event) Time() time.Time {
	zone := time.FixedZone("name", e.TimestampOffset/1000)
	return time.Unix(e.Timestamp/1000, 0).In(zone)
}

type Venue struct {
	Name         string `json:"name"`
	AddressLine1 string `json:"address_1"`
}
