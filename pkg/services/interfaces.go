package services

import (
	"context"

	mirakurun_client "github.com/frauniki/go-mirakurun"
)

type MirakurunInterface interface {
	SetMirakurunEndpoint(string) error
	FetchTuners(context.Context) ([]*mirakurun_client.TunerDevice, error)
	FetchChannels(context.Context) ([]*mirakurun_client.Channel, error)
}
