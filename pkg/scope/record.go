package scope

import (
	"context"

	"github.com/frauniki/Yumemi/api/v1alpha1"
	"github.com/frauniki/Yumemi/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"k8s.io/klog/v2"
	"sigs.k8s.io/cluster-api/util/patch"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type RecordScopeParams struct {
	Client client.Client `validate:"required"`
	Logger *logger.Logger
	Record *v1alpha1.Record `validate:"required"`
}

type RecordScope struct {
	logger.Logger
	client      client.Client
	patchHelper *patch.Helper

	Record *v1alpha1.Record
}

var _ RecordScoper = &RecordScope{}

func NewRecordScope(params RecordScopeParams) (*RecordScope, error) {
	validate := validator.New()
	if err := validate.Struct(params); err != nil {
		return nil, errors.Wrap(err, "record scope params validation failed")
	}

	if params.Logger == nil {
		l := klog.Background()
		params.Logger = logger.NewLogger(l)
	}

	scope := &RecordScope{
		Logger: *params.Logger,
		client: params.Client,
		Record: params.Record,
	}

	helper, err := patch.NewHelper(params.Record, params.Client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init patch helper")
	}
	scope.patchHelper = helper

	return scope, nil
}

func (s *RecordScope) Name() string {
	return s.Record.Name
}

func (s *RecordScope) PatchObject() error {
	return s.patchHelper.Patch(
		context.TODO(),
		s.Record,
	)
}

func (s *RecordScope) Close() error {
	return s.PatchObject()
}
