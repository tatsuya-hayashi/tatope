/*
Copyright 2025.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TatOpeExSpec defines the desired state of TatOpeEx
type TatOpeExSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of TatOpeEx. Edit tatopeex_types.go to remove/update
	Foo string `json:"foo,omitempty"`
	Bar string `json:"bar,omitempty"`
}

// TatOpeExStatus defines the observed state of TatOpeEx
type TatOpeExStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	State string `json:"state,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// TatOpeEx is the Schema for the tatopeexes API
type TatOpeEx struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TatOpeExSpec   `json:"spec,omitempty"`
	Status TatOpeExStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TatOpeExList contains a list of TatOpeEx
type TatOpeExList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TatOpeEx `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TatOpeEx{}, &TatOpeExList{})
}
