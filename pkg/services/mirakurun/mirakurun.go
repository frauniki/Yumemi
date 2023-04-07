package mirakurun

import (
	"context"
	"net/url"

	mirakurun_client "github.com/frauniki/go-mirakurun"
)

func (s *Service) SetMirakurunEndpoint(endpoint string) error {
	var err error
	s.MirakurunClient.BaseURL, err = url.Parse(endpoint)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) FetchTuners(ctx context.Context) ([]*mirakurun_client.TunerDevice, error) {
	tuners, _, err := s.MirakurunClient.GetTuners(ctx)
	if err != nil {
		return nil, err
	}
	return tuners, nil
}

func (s *Service) FetchChannels(ctx context.Context) ([]*mirakurun_client.Channel, error) {
	channels, _, err := s.MirakurunClient.GetChannels(ctx, &mirakurun_client.ChannelsListOptions{})
	if err != nil {
		return nil, err
	}
	return channels, nil
}
