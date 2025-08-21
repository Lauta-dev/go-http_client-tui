package logic

import (
	"bytes"
	"http_client/utils"
	"io"
	"net/http"
	"net/url"
)

type Fetch struct {
	ContentType    string
	Body           string
	URL            string
	StatusCodeText string
	StatusCode     int
}

func Fetching(
	userUrl string,
	verb string,
	h map[string]string,
	qp map[string]string,
	p []string,
	body string) (Fetch, error) {

	baseURL := utils.AddPathParam(p, userUrl)
	u, err := url.Parse(baseURL)

	if err != nil {

		return Fetch{}, err
	}

	URL := utils.AddQueryParam(u, qp)
	req, err := http.NewRequest(verb, URL, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return Fetch{}, err
	}

	for k, v := range h {
		req.Header.Set(k, v)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Fetch{}, err
	}

	defer res.Body.Close()
	bodyByte, err := io.ReadAll(res.Body)

	if err != nil {
		return Fetch{}, err
	}

	bodyToString := string(bodyByte)
	contentType := res.Header.Get("Content-Type")

	return Fetch{
		Body:           bodyToString,
		ContentType:    contentType,
		URL:            URL,
		StatusCodeText: res.Status,
		StatusCode:     res.StatusCode,
	}, nil
}
