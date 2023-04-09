package scope

import (
	"time"

	"github.com/frauniki/Yumemi/api/v1alpha1"
	"github.com/frauniki/Yumemi/pkg/logger"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MirakurunScoper interface {
	logger.Wrapper

	Name() string
	Endpoint() string
	IsDefault() bool

	SetChannelsStatus([]v1alpha1.Channel)
	SetTunersStatus([]v1alpha1.Tuner)
	SetLastUpdatedTime(*metav1.Time)

	PatchObject() error
	Close() error
}

type RecordScoper interface {
	logger.Wrapper

	Name() string
	StartTime() (time.Time, error)
	EndTime() (time.Time, error)
	Phase() v1alpha1.RecordStatusPhase

	PatchObject() error
	Close() error
}
