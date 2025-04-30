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
	"k8s.io/apimachinery/pkg/util/intstr"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TatOpeSpec defines the desired state of TatOpe
type TatOpeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of TatOpe. Edit tatope_types.go to remove/update
	Foo  string `json:"foo,omitempty"`
	Hoge string `json:"hoge"`
	// +listType=set
	Ports []int32            `json:"ports,omitempty"`
	Bar55 intstr.IntOrString `json:"bar,omitempty"`
}

// TatOpeStatus defines the observed state of TatOpe
type TatOpeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	State string `json:"state,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// TatOpe is the Schema for the tatopes API
type TatOpe struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TatOpeSpec   `json:"spec,omitempty"`
	Status TatOpeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TatOpeList contains a list of TatOpe
type TatOpeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TatOpe `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TatOpe{}, &TatOpeList{})
}
