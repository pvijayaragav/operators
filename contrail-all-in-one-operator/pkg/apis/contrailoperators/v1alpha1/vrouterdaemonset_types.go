package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// VrouterDaemonSetSpec defines the desired state of VrouterDaemonSet
// +k8s:openapi-gen=true
type VrouterDaemonSetSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	contrailNodes string `json:"contrailNodes"`
}

// VrouterDaemonSetStatus defines the observed state of VrouterDaemonSet
// +k8s:openapi-gen=true
type VrouterDaemonSetStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VrouterDaemonSet is the Schema for the vrouterdaemonsets API
// +k8s:openapi-gen=true
type VrouterDaemonSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VrouterDaemonSetSpec   `json:"spec,omitempty"`
	Status VrouterDaemonSetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VrouterDaemonSetList contains a list of VrouterDaemonSet
type VrouterDaemonSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VrouterDaemonSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VrouterDaemonSet{}, &VrouterDaemonSetList{})
}
