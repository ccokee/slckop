package controllers

import (
	"context"
	"fmt"
	helm "helm.sh/helm/v3/pkg/action"
	helmlibloader "helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/getter"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"os"
	cachev1alpha1 "github.com/ccokee/slck-operator/api/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	// +kubebuilder:scaffold:imports
)

// SlckReconciler reconciles a Slck object
type SlckReconciler struct {
    client.Client
    Clientset kubernetes.Interface
    Scheme *runtime.Scheme
}


//+kubebuilder:rbac:groups=cache.example.com,resources=slcks,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cache.example.com,resources=slcks/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=cache.example.com,resources=slcks/finalizers,verbs=update
//+kubebuilder:rbac:groups=core,resources=events,verbs=create;patch
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch
//+kubebuilder:rbac:groups=apps,resources=services,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=statefulsets,verbs=get;list;watch
//+kubebuilder:rbac:groups=apps,resources=daemonsets,verbs=get;list;watch;create;update;patch;delete
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
    r.Clientset = mgr.GetClientset()
    return ctrl.NewControllerManagedBy(mgr).
        For(&cachev1alpha1.Slck{}).
        Complete(r)
}



