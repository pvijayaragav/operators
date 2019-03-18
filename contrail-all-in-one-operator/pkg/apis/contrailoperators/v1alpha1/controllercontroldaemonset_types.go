package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ControllerControlDaemonSetSpec defines the desired state of ControllerControlDaemonSet
// +k8s:openapi-gen=true
type ControllerControlDaemonSetSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	contrailNodes string `json:"contrailNodes"`
}

// ControllerControlDaemonSetStatus defines the observed state of ControllerControlDaemonSet
// +k8s:openapi-gen=true
type ControllerControlDaemonSetStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ControllerControlDaemonSet is the Schema for the controllercontroldaemonsets API
// +k8s:openapi-gen=true
type ControllerControlDaemonSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ControllerControlDaemonSetSpec   `json:"spec,omitempty"`
	Status ControllerControlDaemonSetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ControllerControlDaemonSetList contains a list of ControllerControlDaemonSet
type ControllerControlDaemonSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ControllerControlDaemonSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ControllerControlDaemonSet{}, &ControllerControlDaemonSetList{})
}
