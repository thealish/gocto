package gocto

import (
	"errors"
	"net/http"
	"time"
)

const (
	baseURL           = "https://secure.octo.uz/"
	contentTypeHeader = "application/json"
)

var (
	ErrInvalidOrEmptyShopID     = errors.New("invalid or empty shop id")
	ErrInvalidOrEmptySecret     = errors.New("invalid or empty secret")
	ErrEmptyBindNotificationURL = errors.New("empty bind notification url")
)

type Gocto struct {
	octoShopID    int
	octoSecret    string
	returnURL     string
	notifyURL     string
	bindNotifyURL string
	hcl           *http.Client
}

type Options struct {
	OctoShopID    int
	OctoSecret    string
	ReturnURL     string
	NotifyURL     string
	BindNotifyURL string
	HTTPClient    *http.Client
}

func (o Options) validate() error {
	if o.OctoSecret == "" {
		return ErrInvalidOrEmptySecret
	}

	if o.OctoShopID == 0 {
		return ErrInvalidOrEmptyShopID
	}

	if o.BindNotifyURL == "" {
		return ErrEmptyBindNotificationURL
	}

	return nil
}

func New(opt Options) Gocto {
	gocto := Gocto{
		octoShopID:    opt.OctoShopID,
		octoSecret:    opt.OctoSecret,
		returnURL:     opt.ReturnURL,
		notifyURL:     opt.NotifyURL,
		bindNotifyURL: opt.BindNotifyURL,
	}

	if opt.HTTPClient == nil {
		gocto.hcl = &http.Client{
			Timeout: 10 * time.Second,
		}
	}

	return gocto
}
