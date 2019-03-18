package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ControllerConfigDaemonSetSpec defines the desired state of ControllerConfigDaemonSet
// +k8s:openapi-gen=true
type ControllerConfigDaemonSetSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	contrailNodes string `json:"contrailNodes"`
}

// ControllerConfigDaemonSetStatus defines the observed state of ControllerConfigDaemonSet
// +k8s:openapi-gen=true
type ControllerConfigDaemonSetStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ControllerConfigDaemonSet is the Schema for the controllerconfigdaemonsets API
// +k8s:openapi-gen=true
type ControllerConfigDaemonSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ControllerConfigDaemonSetSpec   `json:"spec,omitempty"`
	Status ControllerConfigDaemonSetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ControllerConfigDaemonSetList contains a list of ControllerConfigDaemonSet
type ControllerConfigDaemonSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ControllerConfigDaemonSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ControllerConfigDaemonSet{}, &ControllerConfigDaemonSetList{})
}
