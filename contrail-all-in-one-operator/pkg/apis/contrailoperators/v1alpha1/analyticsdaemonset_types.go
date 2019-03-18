package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AnalyticsDaemonSetSpec defines the desired state of AnalyticsDaemonSet
// +k8s:openapi-gen=true
type AnalyticsDaemonSetSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	contrailNodes string `json:"contrailNodes"`
}

// AnalyticsDaemonSetStatus defines the observed state of AnalyticsDaemonSet
// +k8s:openapi-gen=true
type AnalyticsDaemonSetStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AnalyticsDaemonSet is the Schema for the analyticsdaemonsets API
// +k8s:openapi-gen=true
type AnalyticsDaemonSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AnalyticsDaemonSetSpec   `json:"spec,omitempty"`
	Status AnalyticsDaemonSetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AnalyticsDaemonSetList contains a list of AnalyticsDaemonSet
type AnalyticsDaemonSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AnalyticsDaemonSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AnalyticsDaemonSet{}, &AnalyticsDaemonSetList{})
}
