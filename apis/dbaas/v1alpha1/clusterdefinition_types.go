/*
Copyright 2022.

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

// ClusterDefinitionSpec defines the desired state of ClusterDefinition
type ClusterDefinitionSpec struct {
	// Type define well known cluster types. Valid values are in-list of
	// [state.redis, mq.mqtt, mq.kafka, state.mysql-8, state.mysql-5.7, state.mysql-5.6, state-mongodb]
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MaxLength=24
	Type string `json:"type"`

	// +kubebuilder:validation:MinItems=1
	// +optional
	Components []ClusterDefinitionComponent `json:"components,omitempty"`

	// +kubebuilder:validation:Enum={DoNotTerminate,Halt,Delete,WipeOut}
	DefaultTerminatingPolicy string `json:"defaultTerminationPolicy,omitempty"`

	// +optional
	ConnectionCredential ClusterDefinitionConnectionCredential `json:"connectionCredential,omitempty"`
}

// ClusterDefinitionStatus defines the observed state of ClusterDefinition
type ClusterDefinitionStatus struct {
	// phase - in list of [Available,Deleting]
	// +kubebuilder:validation:Enum={Available,Deleting}
	Phase Phase `json:"phase,omitempty"`
	// +optional
	Message string `json:"message,omitempty"`
	// observedGeneration is the most recent generation observed for this
	// ClusterDefinition. It corresponds to the ClusterDefinition's generation, which is
	// updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:categories={dbaas},scope=Cluster,shortName=cd
//+kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
//+kubebuilder:printcolumn:name="PHASE",type="string",JSONPath=".status.phase",description="status phase"

// ClusterDefinition is the Schema for the clusterdefinitions API
type ClusterDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterDefinitionSpec   `json:"spec,omitempty"`
	Status ClusterDefinitionStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ClusterDefinitionList contains a list of ClusterDefinition
type ClusterDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterDefinition `json:"items"`
}

type ClusterDefinitionComponent struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MaxLength=12
	TypeName string `json:"typeName,omitempty"`

	// +kubebuilder:default=0
	// +kubebuilder:validation:Minimum=0
	MinAvailable int `json:"minAvailable,omitempty"`

	// +kubebuilder:validation:Minimum=0
	MaxAvailable int `json:"maxAvailable,omitempty"`

	// +kubebuilder:default=0
	// +kubebuilder:validation:Minimum=0
	DefaultReplicas int `json:"defaultReplicas,omitempty"`

	// antiAffinity defines components should have anti-affinity constraint to same component type
	// +kubebuilder:default=false
	AntiAffinity bool `json:"antiAffinity,omitempty"`

	// podSpec of final workload
	// +optional
	PodSpec *corev1.PodSpec `json:"podSpec,omitempty"`

	// Service defines the behavior of a service spec.
	// provide readWrite service when ComponentType is Consensus
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// +optional
	Service corev1.ServiceSpec `json:"service,omitempty"`

	// ReadonlyService defines the behavior of a service spec.
	// provide readonly service when ComponentType is Consensus
	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// +optional
	ReadonlyService corev1.ServiceSpec `json:"readonlyService,omitempty"`

	// script exec order：component.pre => roleGroup.pre => component.exec => roleGroup.exec => roleGroup.post => component.post
	// builtin ENV variables:
	// self: OPENDBAAS_SELF_{builtin_properties}
	// rule: OPENDBAAS_{conponent_name}[n]-{roleGroup_name}[n]-{builtin_properties}
	// builtin_properties:
	// - ID # which shows in Cluster.status
	// - HOST # e.g. example-mongodb2-0.example-mongodb2-svc.default.svc.cluster.local
	// - PORT
	// - N # number of current component/roleGroup
	// +optional
	Scripts ClusterDefinitionScripts `json:"scripts,omitempty"`

	// ComponentType defines type of the component
	// +kubebuilder:validation:Required
	// +kubebuilder:default=Stateless
	// +kubebuilder:validation:Enum={Stateless,Stateful,Consensus}
	ComponentType ComponentType `json:"componentType,omitempty"`

	// ConsensusSpec defines consensus related spec if componentType is Consensus
	// CAN'T be empty if componentType is Consensus
	// +optional
	ConsensusSpec ConsensusSetSpec `json:"consensusSpec,omitempty"`
}

type ComponentType string

const (
	Stateless ComponentType = "Stateless"
	Stateful  ComponentType = "Stateful"
	Consensus ComponentType = "Consensus"
)

type ClusterDefinitionScripts struct {
	Default         ClusterDefinitionScript `json:"default,omitempty"`
	Create          ClusterDefinitionScript `json:"create,omitempty"`
	Upgrade         ClusterDefinitionScript `json:"upgrade,omitempty"`
	VerticalScale   ClusterDefinitionScript `json:"verticalScale,omitempty"`
	HorizontalScale ClusterDefinitionScript `json:"horizontalScale,omitempty"`
	Delete          ClusterDefinitionScript `json:"delete,omitempty"`
}

type ClusterDefinitionScript struct {
	Pre  []ClusterDefinitionContainerCMD `json:"pre,omitempty"`
	Post []ClusterDefinitionContainerCMD `json:"post,omitempty"`
}

type ClusterDefinitionContainerCMD struct {
	Container string   `json:"container,omitempty"`
	Command   []string `json:"command,omitempty"`
	Args      []string `json:"args,omitempty"`
}

type ClusterDefinitionUpdateStrategy struct {
	// +kubebuilder:default=1
	// +kubebuilder:validation:Minimum=0
	MaxUnavailable int `json:"maxUnavailable,omitempty"`
	// +kubebuilder:default=0
	// +kubebuilder:validation:Minimum=0
	MaxSurge int `json:"maxSurge,omitempty"`
}

type ClusterDefinitionConnectionCredential struct {
	// +kubebuilder:default=root
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
}

type ClusterDefinitionStatusGeneration struct {

	// ClusterDefinition generation number
	// +optional
	ClusterDefGeneration int64 `json:"clusterDefGeneration,omitempty"`

	// ClusterDefinition sync. status
	// +kubebuilder:validation:Enum={InSync,OutOfSync}
	// +optional
	ClusterDefSyncStatus Status `json:"clusterDefSyncStatus,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ClusterDefinition{}, &ClusterDefinitionList{})
}
