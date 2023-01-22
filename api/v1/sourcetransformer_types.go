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

package v1

import (
	kustomize "github.com/fluxcd/kustomize-controller/api/v1beta2"
	"github.com/fluxcd/pkg/apis/meta"
	source "github.com/fluxcd/source-controller/api/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SourceTransformerSpec defines the desired state of SourceTransformer
type SourceTransformerSpec struct {
	// The FluxCD Source that we want to transform
	SourceRef   kustomize.CrossNamespaceSourceReference `json:"sourceRef"`
	Transformer Transformer                             `json:"transformer"`
}

type Transformer struct {
	Command string `json:"command"`
	Output  string `json:"output"`
}

// SourceTransformerStatus defines the observed state of SourceTransformer
type SourceTransformerStatus struct {
	// ObservedGeneration is the last observed generation.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// URL is the download link for the artifact output of the last OCI Repository sync.
	// +optional
	URL string `json:"url,omitempty"`

	// Artifact represents the output of the last successful OCI Repository sync.
	// +optional
	Artifact *source.Artifact `json:"artifact,omitempty"`

	meta.ReconcileRequestStatus `json:",inline"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// SourceTransformer is the Schema for the sourcetransformers API
type SourceTransformer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SourceTransformerSpec   `json:"spec,omitempty"`
	Status SourceTransformerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SourceTransformerList contains a list of SourceTransformer
type SourceTransformerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SourceTransformer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SourceTransformer{}, &SourceTransformerList{})
}
