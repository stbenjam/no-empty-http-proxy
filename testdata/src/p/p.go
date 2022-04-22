package p

import (
	"net/http"
	"net/url"
)

func _() {
	_ = http.Client{
		Transport: &http.Transport{}, // want "http.Transport should set Proxy"
	}

	_ = http.Transport{} // want "http.Transport should set Proxy"

	_ = http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}

	_ = http.Transport{
		Proxy: func(*http.Request) (*url.URL, error) {
			return nil, nil
		},
	}

	_ = http.Transport{
		Proxy: nil,
	}
}
