package mirakurun

import (
	mirakurun_client "github.com/frauniki/go-mirakurun"
)

type Service struct {
	MirakurunClient *mirakurun_client.Client
}

func NewService() *Service {
	return &Service{
		MirakurunClient: mirakurun_client.NewClient(),
	}
}
