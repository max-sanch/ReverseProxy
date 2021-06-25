package handler

import (
	"fmt"
	"github.com/max-sanch/ReverseProxy/pkg/service"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) WriteResponse(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	urlStr := fmt.Sprintf("%s%s", viper.GetString("targetHost"), r.URL.Path)
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating request to %s", urlStr), http.StatusInternalServerError)
		return
	}

	u, err := url.Parse(urlStr)
	if err != nil {
		http.Error(w, "Error URL parse", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Host", u.Host)
	req.Header.Set("URL", urlStr)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error sending request to %s", urlStr), http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(res.StatusCode)
	_, err = w.Write(body)
	if err != nil {
		http.Error(w, "Error write response body", http.StatusInternalServerError)
		return
	}
}