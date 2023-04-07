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
	"fmt"
	"time"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/frauniki/Yumemi/api/v1alpha1"
	"github.com/frauniki/Yumemi/pkg/hash"
	"github.com/frauniki/Yumemi/pkg/logger"
	"github.com/frauniki/Yumemi/pkg/scope"
	"github.com/frauniki/Yumemi/pkg/services"
	"github.com/frauniki/Yumemi/pkg/services/mirakurun"
	"github.com/pkg/errors"
)

// MirakurunReconciler reconciles a Mirakurun object
type MirakurunReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	mirakurunServiceFactory func() services.MirakurunInterface
}

func (r *MirakurunReconciler) getMirakurunService() services.MirakurunInterface {
	if r.mirakurunServiceFactory != nil {
		return r.mirakurunServiceFactory()
	}
	return mirakurun.NewService()
}

//+kubebuilder:rbac:groups=yumemi.sinoa.jp,resources=mirakuruns,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=yumemi.sinoa.jp,resources=mirakuruns/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=yumemi.sinoa.jp,resources=mirakuruns/finalizers,verbs=update

func (r *MirakurunReconciler) Reconcile(ctx context.Context, req ctrl.Request) (_ ctrl.Result, reterr error) {
	log := logger.FromContext(ctx)

	mirakurun := &v1alpha1.Mirakurun{}
	if err := r.Get(ctx, req.NamespacedName, mirakurun); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	mirakurunScope, err := scope.NewMirakurunScope(scope.MirakurunScopeParams{
		Logger:    log,
		Client:    r.Client,
		Mirakurun: mirakurun,
	})
	if err != nil {
		return ctrl.Result{}, errors.Wrap(err, "failed to create mirakurun scope")
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

	if err := r.reconcileMirakurun(ctx, s); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *MirakurunReconciler) reconcileMirakurun(ctx context.Context, s *scope.MirakurunScope) error {
	now := time.Now()

	s.Logger.Info("Reconciling Mirakurun")

	mirakurunService := r.getMirakurunService()

	if err := mirakurunService.SetMirakurunEndpoint(s.Endpoint()); err != nil {
		return err
	}

	tuners, err := mirakurunService.FetchTuners(ctx)
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
	s.SetTunersStatus(ts)

	channels, err := mirakurunService.FetchChannels(ctx)
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
	s.SetChannelsStatus(cs)

	s.SetLastUpdatedTime(&metav1.Time{Time: now})

	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MirakurunReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.Mirakurun{}).
		Complete(r)
}

func generateHashedChannelName(name string) (string, error) {
	// hashSize = 31 - length of "ch" - length of "-" = 29
	shortName, err := hash.Base36TruncatedHash(name, 29)
	if err != nil {
		return "", errors.Wrap(err, "unable to create channel name")
	}

	return fmt.Sprintf("%s-%s", "ch", shortName), nil
}
