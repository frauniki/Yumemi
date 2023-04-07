package scope

import (
	"context"

	"github.com/frauniki/Yumemi/api/v1alpha1"
	"github.com/frauniki/Yumemi/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
	"sigs.k8s.io/cluster-api/util/patch"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type MirakurunScopeParams struct {
	Client    client.Client `validate:"required"`
	Logger    *logger.Logger
	Mirakurun *v1alpha1.Mirakurun `validate:"required"`
}

type MirakurunScope struct {
	logger.Logger
	client      client.Client
	patchHelper *patch.Helper

	Mirakurun *v1alpha1.Mirakurun
}

var _ MirakurunScoper = &MirakurunScope{}

func NewMirakurunScope(params MirakurunScopeParams) (*MirakurunScope, error) {
	validate := validator.New()
	if err := validate.Struct(params); err != nil {
		return nil, errors.Wrap(err, "mirakurun scope params validation failed")
	}

	if params.Logger == nil {
		l := klog.Background()
		params.Logger = logger.NewLogger(l)
	}

	scope := &MirakurunScope{
		Logger:    *params.Logger,
		client:    params.Client,
		Mirakurun: params.Mirakurun,
	}

	helper, err := patch.NewHelper(params.Mirakurun, params.Client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init patch helper")
	}
	scope.patchHelper = helper

	return scope, nil
}

func (s *MirakurunScope) Name() string {
	return s.Mirakurun.Name
}

func (s *MirakurunScope) Endpoint() string {
	return s.Mirakurun.Spec.Endpoint
}

func (s *MirakurunScope) IsDefault() bool {
	return s.Mirakurun.Spec.IsDefault
}

func (s *MirakurunScope) SetTunersStatus(ts []v1alpha1.Tuner) {
	s.Mirakurun.Status.Tuners = ts
}

func (s *MirakurunScope) SetChannelsStatus(cs []v1alpha1.Channel) {
	s.Mirakurun.Status.Channels = cs
}

func (s *MirakurunScope) SetLastUpdatedTime(t *metav1.Time) {
	s.Mirakurun.Status.LastUpdatedTime = t
}

func (s *MirakurunScope) PatchObject() error {
	return s.patchHelper.Patch(
		context.TODO(),
		s.Mirakurun,
	)
}

func (s *MirakurunScope) Close() error {
	return s.PatchObject()
}
