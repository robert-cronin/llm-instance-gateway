/*
Copyright 2024.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// LLMServerPoolSpec defines the desired state of LLMServerPool
type LLMServerPoolSpec struct {

	// ModelServerSelector uses a map of label to watch model server pods
	// that should be included in the LLMServerPool. ModelServers should not
	// be with any other Service or LLMServerPool, that behavior is not supported
	// and will result in sub-optimal utilization.
	// Due to this selector being translated to a service a simple map is used instead
	// of: https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#LabelSelector
	// To avoid footshoot errors when the https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#LabelSelectorAsMap would be used.
	ModelServerSelector map[string]string `json:"modelServerSelector,omitempty"`

	// TargetPort is the port number that the model servers within the pool expect
	// to recieve traffic from.
	// This maps to the TargetPort in: https://pkg.go.dev/k8s.io/api/core/v1#ServicePort
	TargetPort int32 `json:"targetPort,omitempty"`
}

// LLMServerPoolStatus defines the observed state of LLMServerPool
type LLMServerPoolStatus struct {

	// Conditions track the state of the LLMServerPool.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +genclient

// LLMServerPool is the Schema for the llmserverpools API
type LLMServerPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LLMServerPoolSpec   `json:"spec,omitempty"`
	Status LLMServerPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LLMServerPoolList contains a list of LLMServerPool
type LLMServerPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LLMServerPool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LLMServerPool{}, &LLMServerPoolList{})
}
