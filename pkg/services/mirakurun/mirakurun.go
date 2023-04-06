package mirakurun

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/frauniki/Yumemi/api/v1alpha1"
	"github.com/frauniki/Yumemi/pkg/hash"
	mirakurun_client "github.com/frauniki/go-mirakurun"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Service) ReconcileMirakurun(ctx context.Context) error {
	now := time.Now()

	var err error
	s.MirakurunClient.BaseURL, err = url.Parse(s.scope.Endpoint())
	if err != nil {
		return err
	}

	tuners, err := s.fetchTuners(ctx)
	if err != nil {
		return err
	}
	ts := []v1alpha1.Tuner{}
	for _, t := range tuners {
		ts = append(ts, v1alpha1.Tuner{
			Name:    t.Name,
			Types:   t.Types,
			IsReady: t.IsAvailable,
		})
	}
	s.scope.SetTunersStatus(ts)

	channels, err := s.fetchChannels(ctx)
	if err != nil {
		return err
	}
	cs := []v1alpha1.Channel{}
	for _, c := range channels {
		hashedName, err := generateHashedChannelName(c.Name)
		if err != nil {
			return err
		}
		cs = append(cs, v1alpha1.Channel{
			Name:        hashedName,
			DisplayName: c.Name,
			Type:        c.Type,
			Channel:     c.Channel,
		})
	}
	s.scope.SetChannelsStatus(cs)

	s.scope.SetLastUpdatedTime(&metav1.Time{Time: now})

	return nil
}

func (s *Service) fetchTuners(ctx context.Context) ([]*mirakurun_client.TunerDevice, error) {
	tuners, _, err := s.MirakurunClient.GetTuners(ctx)
	if err != nil {
		return nil, err
	}
	return tuners, nil
}

func (s *Service) fetchChannels(ctx context.Context) ([]*mirakurun_client.Channel, error) {
	channels, _, err := s.MirakurunClient.GetChannels(ctx, &mirakurun_client.ChannelsListOptions{})
	if err != nil {
		return nil, err
	}
	return channels, nil
}

func generateHashedChannelName(name string) (string, error) {
	// hashSize = 32 - length of "ch" - length of "-" = 29
	shortName, err := hash.Base36TruncatedHash(name, 29)
	if err != nil {
		return "", errors.Wrap(err, "unable to create channel name")
	}

	return fmt.Sprintf("%s-%s", "ch", shortName), nil
}
