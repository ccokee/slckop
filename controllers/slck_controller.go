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
	cachev1alpha1 "github.com/ccokee/slckop/api/v1alpha1"
	helm "helm.sh/helm/v3/pkg/action"
	helmlibloader "helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/getter"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"os"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// SlckReconciler reconciles a Slck object
type SlckReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=cli.redrvm.cloud,resources=slcks,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cli.redrvm.cloud,resources=slcks/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=cli.redrvm.cloud,resources=slcks/finalizers,verbs=update

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
	log := log.FromContext(ctx)

	// Lookup the Slck instance for this reconcile request
	slck := &cachev1alpha1.Slck{}
	err := r.Get(ctx, req.NamespacedName, slck)
	if err != nil {
		if errors.IsNotFound(err) {
			log.Info("Slck resource not found. Ignoring since object must be deleted.")
			return ctrl.Result{}, nil
		}
		log.Error(err, "Failed to get Slck")
		return ctrl.Result{}, err
	}

	// Install the Helm chart using values from SlckSpec
	err = r.installHelmChart(slck)
	if err != nil {
		log.Error(err, "Failed to install Helm chart")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *SlckReconciler) installHelmChart(slck *cachev1alpha1.Slck) error {
	actionConfig := new(helm.Configuration)
	settings := cli.New()

	// Initialize the Helm action configuration
	err := actionConfig.Init(settings.RESTClientGetter(), slck.Namespace, os.Getenv("HELM_DRIVER"), log.Log)
	if err != nil {
		return fmt.Errorf("failed to initialize Helm action configuration: %w", err)
	}

	installClient := helm.NewInstall(actionConfig)
	installClient.Namespace = slck.Namespace
	installClient.ReleaseName = slck.Name

	chartPath, err := installClient.ChartPathOptions.LocateChart(slck.Spec.ChartRepo, settings)
	if err != nil {
		return fmt.Errorf("failed to locate Helm chart: %w", err)
	}

	chart, err := helmlibloader.Load(chartPath)
	if err != nil {
		return fmt.Errorf("failed to load Helm chart: %w", err)
	}

	providers := getter.All(settings)
	values, err := installClient.MergeValues(providers, slck.Spec.Values)
	if err != nil {
		return fmt.Errorf("failed to merge Helm values: %w", err)
	}

	_, err = installClient.Run(chart, values)
	if err != nil {
		return fmt.Errorf("failed to run Helm installation: %w", err)
	}

	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *SlckReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cliv1alpha1.Slck{}).
		Complete(r)
}
