package main

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
)

func AddPathParam(params []string, baseURL string) string {
	pathParam := ""
	for _, v := range params {
		pathParam += "/" + v
	}

	return baseURL + pathParam
}

func addQueryParam(u *url.URL, qp map[string]string) *url.URL {
	q := u.Query()

	for k, v := range qp {
		if q.Get(k) == "" {
			q.Add(k, v)

		} else {
			q.Set(k, v)
		}
	}

	u.RawQuery = q.Encode()
	return u
}

func Fetching(
	userUrl string,
	verb string,
	h map[string]string,
	qp map[string]string,
	p []string,
	body string) (string, error) {
	baseURL := AddPathParam(p, userUrl)
	u, err := url.Parse(baseURL)

	if err != nil {

		return "", err
	}

	finalURL := addQueryParam(u, qp)

	userUrl = finalURL.String()

	req, err := http.NewRequest(verb, userUrl, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return "", err
	}

	for k, v := range h {
		req.Header.Set(k, v)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	status = res.Status
	contentType = res.Header.Get("Content-Type")
	completeUrl = userUrl

	return string(bytes), nil
}
