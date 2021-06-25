package service

type TargetRequest interface {
}

type Hashing interface {
}

type Service struct {
	TargetRequest
	Hashing
}

func NewService() *Service {
	return &Service{}
}

