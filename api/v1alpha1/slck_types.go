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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type AutoscalingConfig struct {
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// +optional
	MinReplicas int32 `json:"minReplicas,omitempty"`
	// +optional
	MaxReplicas int32 `json:"maxReplicas,omitempty"`
	// +optional
	TargetCPUUtilizationPercentage int32 `json:"targetCPUUtilizationPercentage,omitempty"`
}

type ImageConfig struct {
	// +optional
	Repository string `json:"repository,omitempty"`
	// +optional
	PullPolicy string `json:"pullPolicy,omitempty"`
	Tag        string `json:"tag,omitempty"`
}

type ServiceConfig struct {
	// +optional
	Type string `json:"type,omitempty"`
	// +optional
	Port int32 `json:"port,omitempty"`
}

type RedisConfig struct {
	// +optional
	Replicas int32 `json:"replicas,omitempty"`
	// +optional
	Service ServiceConfig `json:"service,omitempty"`
	// +optional
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`
}

type ResourceUsage struct {
	//+optional
	CPU string `json:"cpu,omitempty"`
	//+optional
	Mem string `json:"mem,omitempty"`
}

// SlckSpec defines the desired state of Slck
type SlckSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Helm chart related configurations
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	ChartRepo string `json:"chartRepo,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	ChartName string `json:"chartName,omitempty"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	ChartVersion string `json:"chartVersion,omitempty"`
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Namespace string `json:"namespace,omitempty"`
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Values map[string]string `json:"values,omitempty"`

	// Autoscaling configuration
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Autoscaling AutoscalingConfig `json:"autoscaling"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// Image configuration
	Image ImageConfig `json:"image"`

	// ImagePullSecrets configuration
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty"`

	// Number of replicas
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Replicas int32 `json:"replicas"`

	// Resource requests and limits
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Resources corev1.ResourceRequirements `json:"resources"`

	// NameOverride and FullnameOverride configurations
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	NameOverride string `json:"nameOverride,omitempty"`
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	FullnameOverride string `json:"fullnameOverride,omitempty"`

	// Service configuration
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Service ServiceConfig `json:"service"`

	// NodeSelector configuration
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// Tolerations configuration
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Tolerations []corev1.Toleration `json:"tolerations,omitempty"`

	// Affinity configuration
	// +optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Affinity corev1.Affinity `json:"affinity,omitempty"`

	// Redis configuration
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Redis RedisConfig `json:"redis"`
}

// SlckStatus defines the observed state of Slck
type SlckStatus struct {
	// Existing Conditions field
	// +operator-sdk:csv:customresourcedefinitions:type=status
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`

	// Additional fields to enrich status data
	// +operator-sdk:csv:customresourcedefinitions:type=status
	Replicas int32 `json:"replicas,omitempty"`

	// +operator-sdk:csv:customresourcedefinitions:type=status
	Resources ResourceUsage `json:"resources,omitempty"`

	// +operator-sdk:csv:customresourcedefinitions:type=status
	LastError string `json:"lastError,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Slck is the Schema for the slcks API
type Slck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SlckSpec   `json:"spec,omitempty"`
	Status SlckStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SlckList contains a list of Slck
type SlckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Slck `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Slck{}, &SlckList{})
}
