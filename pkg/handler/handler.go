package handler

import (
	"fmt"
	"github.com/max-sanch/ReverseProxy/pkg/service"
	"github.com/spf13/viper"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) WriteResponse(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	urlStr := fmt.Sprintf("%s%s", viper.GetString("targetHost"), r.URL.Path)

	body, err := getBody(w, *h.services, urlStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(body)
	if err != nil {
		http.Error(w, "error write response body", http.StatusInternalServerError)
		return
	}
}

func getBody(w http.ResponseWriter, services service.Service, urlStr string) ([]byte, error) {
	isReceived, body := services.GetBodyFromHash(urlStr)

	if !isReceived {
		res, err := services.GetTargetResponse(urlStr)
		if err != nil {
			return nil, err
		}

		defer res.Body.Close()

		body, err = services.GetTargetResponseBody(res)
		if err != nil {
			return nil, err
		}

		services.SetBodyInHash(urlStr, body)
		w.WriteHeader(res.StatusCode)
	}

	return body, nil
}