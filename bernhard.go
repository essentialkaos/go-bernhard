package bernhard

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"sync"

	"pkg.re/essentialkaos/ek.v10/req"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	STATE_OK       = "ok"
	STATE_WARN     = "warn"
	STATE_PROBLEM  = "problem"
	STATE_CRITICAL = "critical"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Alert contains alert data
type Alert struct {
	Service string `json:"service"`
	Host    string `json:"host"`
	TTL     int    `json:"ttl"`
	State   string `json:"state"`
	Desc    string `json:"description"`
}

// Client basic Bernhard client
type Client struct {
	url   string
	data  *alertsData
	mutex *sync.RWMutex
}

// ////////////////////////////////////////////////////////////////////////////////// //

type alertsData struct {
	Items []Alert `json:"items"`
}

// ////////////////////////////////////////////////////////////////////////////////// //

var ErrServiceNotSet = errors.New("Service field is mandatory and must be set")
var ErrHostNotSet = errors.New("Host field is mandatory and must be set")
var ErrStateNotSet = errors.New("State field is mandatory and must be set")
var ErrBadServiceName = errors.New("Service name contains disallowed symbols")

// ////////////////////////////////////////////////////////////////////////////////// //

var serviceValidationRegex = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)

// ////////////////////////////////////////////////////////////////////////////////// //

// NewClient creates new client instance
func NewClient(bernhardURL string) (*Client, error) {
	_, err := url.Parse(bernhardURL)

	if err != nil {
		return nil, err
	}

	return &Client{
		url: bernhardURL,
	}, nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Add adds new alert to stack
func (c *Client) Add(a Alert) error {
	err := validateAlert(a)

	if err != nil {
		return err
	}

	if c.mutex == nil {
		c.mutex = &sync.RWMutex{}
	}

	c.mutex.Lock()

	if c.data == nil {
		c.data = &alertsData{}
	}

	c.data.Items = append(c.data.Items, a)

	c.mutex.Unlock()

	return nil
}

// Send sends alerts data to Bernhard service
func (c *Client) Send() error {
	if c.mutex == nil {
		c.mutex = &sync.RWMutex{}
	}

	if c.data == nil || len(c.data.Items) == 0 {
		return nil
	}

	c.mutex.RLock()

	defer c.mutex.RUnlock()

	resp, err := req.Request{
		URL:         c.url + "/events",
		Body:        c.data,
		ContentType: req.CONTENT_TYPE_JSON,
		AutoDiscard: true,
	}.Post()

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Bernhard return status code %d", resp.StatusCode)
	}

	c.data.Items = nil

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// validateAlert validates mandatory alert fields
func validateAlert(a Alert) error {
	switch {
	case a.Service == "":
		return ErrServiceNotSet
	case a.Host == "":
		return ErrHostNotSet
	case a.State == "":
		return ErrStateNotSet
	case !serviceValidationRegex.MatchString(a.Service):
		return ErrBadServiceName
	}

	return nil
}
