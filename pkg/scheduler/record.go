package scheduler

import (
	"context"
	"time"

	"github.com/frauniki/Yumemi/api/v1alpha1"
	"github.com/frauniki/Yumemi/pkg/logger"
	"github.com/frauniki/Yumemi/pkg/scope"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	recordPreparationStartOffsetTime = -5 * time.Minute
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

		s, err := scope.NewRecordScope(scope.RecordScopeParams{
			Client: s.client,
			Logger: s.logger,
			Record: &record,
		})
		if err != nil {
			return err
		}
		if err := ReconcileRecord(ctx, s); err != nil {
			return err
		}
	}

	return nil
}

func ReconcileRecord(ctx context.Context, s *scope.RecordScope) (err error) {
	now := time.Now()

	s.Logger.Info("Reconciling Record")

	changed := false
	defer func() {
		if changed {
			err = s.PatchObject()
		}
	}()

	startTime, err := s.StartTime()
	if err != nil {
		return err
	}
	if startTime == (time.Time{}) {
		startTime = now
	}
	offsetStartTime := startTime.Add(recordPreparationStartOffsetTime)
	endTime, err := s.EndTime()
	if err != nil {
		return err
	}

	switch s.Phase() {
	case v1alpha1.RecordPreparation, v1alpha1.RecordFinished, v1alpha1.RecordCanceled, v1alpha1.RecordFailed:
		return nil

	case v1alpha1.RecordScheduled:
		if offsetStartTime.Unix() >= now.Unix() {
			// TODO: create record job
			s.Record.Status.Phase = v1alpha1.RecordPreparation
			changed = true
		}

	case v1alpha1.RecordRecording:
		if endTime.Unix() > now.Unix() {
			// TODO: check record job status
			s.Record.Status.Phase = v1alpha1.RecordFinished
			changed = true
			return nil
		}

	default:
		if offsetStartTime.Unix() < now.Unix() {
			s.Record.Status.Phase = v1alpha1.RecordScheduled
			changed = true
		}
	}

	return nil
}
