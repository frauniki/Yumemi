/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"github.com/frauniki/Yumemi/api/v1alpha1"
	"github.com/frauniki/Yumemi/pkg/logger"
	"github.com/frauniki/Yumemi/pkg/scheduler"
	"github.com/frauniki/Yumemi/pkg/scope"
	"github.com/frauniki/Yumemi/pkg/services"
	"github.com/frauniki/Yumemi/pkg/services/mirakurun"
	"github.com/pkg/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// RecordReconciler reconciles a Record object
type RecordReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	mirakurunServiceFactory func() services.MirakurunInterface
}

func (r *RecordReconciler) getMirakurunService() services.MirakurunInterface {
	if r.mirakurunServiceFactory != nil {
		return r.mirakurunServiceFactory()
	}
	return mirakurun.NewService()
}

//+kubebuilder:rbac:groups=yumemi.sinoa.jp,resources=records,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=yumemi.sinoa.jp,resources=records/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=yumemi.sinoa.jp,resources=records/finalizers,verbs=update

func (r *RecordReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logger.FromContext(ctx)

	record := &v1alpha1.Record{}
	if err := r.Get(ctx, req.NamespacedName, record); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	recordScope, err := scope.NewRecordScope(scope.RecordScopeParams{
		Logger: log,
		Client: r.Client,
		Record: record,
	})
	if err != nil {
		return ctrl.Result{}, errors.Wrap(err, "failed to create record scope")
	}

	defer func() {
		if err := recordScope.Close(); err != nil {
			log.Error(err, "failed to close record scope")
		}
	}()

	if !record.DeletionTimestamp.IsZero() {
		return r.reconcileDelete(ctx, recordScope)
	}

	return r.reconcileNormal(ctx, recordScope)
}

func (r *RecordReconciler) reconcileDelete(ctx context.Context, s *scope.RecordScope) (ctrl.Result, error) {
	return ctrl.Result{}, nil
}

func (r *RecordReconciler) reconcileNormal(ctx context.Context, s *scope.RecordScope) (ctrl.Result, error) {
	// TODO: add finalizer

	if err := scheduler.ReconcileRecord(ctx, s); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RecordReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.Record{}).
		Complete(r)
}
