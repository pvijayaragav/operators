package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ConfigDbDaemonSetSpec defines the desired state of ConfigDbDaemonSet
// +k8s:openapi-gen=true
type ConfigDbDaemonSetSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	contrailNodes string `json:"contrailNodes"`
}

// ConfigDbDaemonSetStatus defines the observed state of ConfigDbDaemonSet
// +k8s:openapi-gen=true
type ConfigDbDaemonSetStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConfigDbDaemonSet is the Schema for the configdbdaemonsets API
// +k8s:openapi-gen=true
type ConfigDbDaemonSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigDbDaemonSetSpec   `json:"spec,omitempty"`
	Status ConfigDbDaemonSetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConfigDbDaemonSetList contains a list of ConfigDbDaemonSet
type ConfigDbDaemonSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigDbDaemonSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConfigDbDaemonSet{}, &ConfigDbDaemonSetList{})
}
