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

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/frauniki/Yumemi/api/v1alpha1"
	"github.com/frauniki/Yumemi/pkg/scope"
	"github.com/frauniki/Yumemi/pkg/services"
	"github.com/frauniki/Yumemi/pkg/services/mirakurun"
	"github.com/pkg/errors"
)

// MirakurunReconciler reconciles a Mirakurun object
type MirakurunReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	mirakurunServiceFavtory func(scope.MirakurunScope) services.MirakurunInterface
}

func (r *MirakurunReconciler) getMirakurunService(s scope.MirakurunScope) services.MirakurunInterface {
	if r.mirakurunServiceFavtory != nil {
		return r.mirakurunServiceFavtory(s)
	}
	return mirakurun.NewService(s)
}

//+kubebuilder:rbac:groups=yumemi.sinoa.jp,resources=mirakuruns,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=yumemi.sinoa.jp,resources=mirakuruns/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=yumemi.sinoa.jp,resources=mirakuruns/finalizers,verbs=update

func (r *MirakurunReconciler) Reconcile(ctx context.Context, req ctrl.Request) (_ ctrl.Result, reterr error) {
	_ = log.FromContext(ctx)

	mirakurun := &v1alpha1.Mirakurun{}
	if err := r.Get(ctx, req.NamespacedName, mirakurun); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	mirakurunScope, err := scope.NewMirakurunScope(scope.MirakurunScopeParams{
		Client:    r.Client,
		Mirakurun: mirakurun,
	})
	if err != nil {
		return ctrl.Result{}, errors.Wrap(err, "failed to create scope")
	}

	defer func() {
		if err := mirakurunScope.Close(); err != nil {
			reterr = err
		}
	}()

	if !mirakurun.DeletionTimestamp.IsZero() {
		return r.reconcileDelete(ctx, mirakurunScope)
	}

	return r.reconcileNormal(ctx, mirakurunScope)
}

func (r *MirakurunReconciler) reconcileDelete(ctx context.Context, s *scope.MirakurunScope) (ctrl.Result, error) {
	return ctrl.Result{}, nil
}

func (r *MirakurunReconciler) reconcileNormal(ctx context.Context, s *scope.MirakurunScope) (ctrl.Result, error) {
	// TODO: add finalizer

	mirakurunService := r.getMirakurunService(*s)

	if err := mirakurunService.ReconcileMirakurun(ctx); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MirakurunReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.Mirakurun{}).
		Complete(r)
}
