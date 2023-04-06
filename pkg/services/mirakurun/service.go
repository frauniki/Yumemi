package mirakurun

import (
	"github.com/frauniki/Yumemi/pkg/scope"
	mirakurun_client "github.com/frauniki/go-mirakurun"
)

type Service struct {
	scope           scope.MirakurunScope
	MirakurunClient *mirakurun_client.Client
}

func NewService(scope scope.MirakurunScope) *Service {
	return &Service{
		scope:           scope,
		MirakurunClient: mirakurun_client.NewClient(),
	}
}
