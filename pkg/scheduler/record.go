package scheduler

import (
	"context"
	"time"

	"github.com/frauniki/Yumemi/api/v1alpha1"
	"github.com/frauniki/Yumemi/pkg/logger"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type RecordScheduler struct {
	logger     *logger.Logger
	interval   time.Duration
	client     client.Client
	namespaces []string
}

func NewRecordScheduler(
	log *logger.Logger,
	interval time.Duration,
	client client.Client,
	namespaces []string,
) *RecordScheduler {
	return &RecordScheduler{
		logger:     log,
		interval:   interval,
		client:     client,
		namespaces: namespaces,
	}
}

func (s *RecordScheduler) Start(ctx context.Context) {
	ticker := time.NewTicker(s.interval)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if err := s.recordSchedule(ctx); err != nil {
					s.logger.Error(err, "failed to record schedule")
				}
			}
		}
	}()
}

func (s *RecordScheduler) recordSchedule(ctx context.Context) error {
	now := time.Now()

	s.logger.Debug("Exec record schedule", "time", now.String())

	records := []v1alpha1.Record{}
	if len(s.namespaces) == 0 {
		s.namespaces = []string{""}
	}
	for _, n := range s.namespaces {
		r := &v1alpha1.RecordList{}
		if err := s.client.List(ctx, r, &client.ListOptions{
			//FieldSelector: , TODO: filter by spec.suspend
			Namespace: n,
		}); err != nil {
			return err
		}
		records = append(records, r.Items...)
	}

	for _, record := range records {
		s.logger.Debug(record.Name)
	}

	return nil
}
