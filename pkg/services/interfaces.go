package services

import "context"

type MirakurunInterface interface {
	ReconcileMirakurun(context.Context) error
}
