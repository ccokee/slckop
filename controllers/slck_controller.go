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
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	cachev1alpha1 "github.com/ccokee/slck-operator/api/v1alpha1"
)

// SlckReconciler reconciles a Slck object
type SlckReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=cache.redrvm.cloud,resources=slcks,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cache.redrvm.cloud,resources=slcks/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=cache.redrvm.cloud,resources=slcks/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Slck object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *SlckReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	// Lookup the slck instance for this reconcile request
	slck := &cachev1alpha1.Slck{}
	err := r.Get(ctx, req.NamespacedName, slck)

	return ctrl.Result{Requeue: true}, err
}

// +kubebuilder:rbac:groups=cache.redrvum.es,resources=slcks,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cache.redrvum.es,resources=slcks/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=cache.redrvum.es,resources=slcks/finalizers,verbs=update
// +kubebuilder:rbac:groups=core,resources=events,verbs=create;patch
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch
// SetupWithManager sets up the controller with the Manager.
func (r *SlckReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cachev1alpha1.Slck{}).
		Owns(&appsv1.Deployment{}).
		WithOptions(controller.Options{MaxConcurrentReconciles: 2}).
		Complete(r)
}
