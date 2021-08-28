package spotify

import (
	"net/http"
	"net/url"
)

const APIhost = "api.spotify"

type API struct {
	token string
}

func NewApi(token string) *API {
	return &API{token}
}

func (a *API) get(apiVersion, endpoint string, query url.Values, res interface{}) error {
	return a.call(http.MethodGet, apiVersion, endpoint, query, nil, res)
}
