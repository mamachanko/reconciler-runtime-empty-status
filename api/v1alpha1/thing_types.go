/*
Copyright 2022.
*/

package v1alpha1

import (
	"github.com/vmware-labs/reconciler-runtime/apis"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ThingSpec defines the desired state of Thing
type ThingSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Thing. Edit thing_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// ThingStatus defines the observed state of Thing
type ThingStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	apis.Status `json:",inline"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Thing is the Schema for the things API
type Thing struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ThingSpec   `json:"spec,omitempty"`
	Status ThingStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ThingList contains a list of Thing
type ThingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Thing `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Thing{}, &ThingList{})
}
