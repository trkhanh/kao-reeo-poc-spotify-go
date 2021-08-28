package spotify

import (
	"io"
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

func (a *API) post(apiVersion, endpoint string, query url.Values, body io.Reader, res interface{}) error {
	return a.call(http.MethodPost, apiVersion, endpoint, query, body, res)
}

func (a *API) put(apiVersion, endpoint string, query url.Values, body io.Reader) error {
	return a.call(http.MethodPut, apiVersion, endpoint, query, body, nil)
}

func (a *API) delete(apiVersion, endpoint string, query url.Values) error {
	return a.call(http.MethodDelete, apiVersion, endpoint, query, nil, nil)
}

func 