package service

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type TargetResponseService struct {
}

func NewTargetResponse() *TargetResponseService {
	return &TargetResponseService{}
}

func (t *TargetResponseService) GetTargetResponse(urlStr string) (*http.Response, error) {
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error creating request to %s", urlStr))
	}

	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, errors.New("error URL parse")
	}

	req.Header.Set("Host", u.Host)
	req.Header.Set("URL", urlStr)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error sending request to %s", urlStr))
	}

	return res, nil
}

func (t *TargetResponseService) GetTargetResponseBody(res *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("error reading response body")
	}
	return body, nil
}
