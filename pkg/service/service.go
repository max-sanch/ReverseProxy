package service

import (
	"github.com/go-redis/redis/v8"
	"net/http"
)

type TargetRequest interface {
	GetTargetResponse(urlStr string) (*http.Response, error)
	GetTargetResponseBody(res *http.Response) ([]byte, error)
}

type Hashing interface {
	GetBodyFromHash(urlStr string) (bool, []byte)
	SetBodyInHash(urlStr string, body []byte)
}

type Service struct {
	TargetRequest
	Hashing
}

func NewService(rdb *redis.Client) *Service {
	return &Service{
		TargetRequest: NewTargetResponse(),
		Hashing:       NewHashing(rdb),
	}
}

