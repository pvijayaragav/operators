package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AnalyticsDbDaemonSetSpec defines the desired state of AnalyticsDbDaemonSet
// +k8s:openapi-gen=true
type AnalyticsDbDaemonSetSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	contrailNodes string `json:"contrailNodes"`
}

// AnalyticsDbDaemonSetStatus defines the observed state of AnalyticsDbDaemonSet
// +k8s:openapi-gen=true
type AnalyticsDbDaemonSetStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AnalyticsDbDaemonSet is the Schema for the analyticsdbdaemonsets API
// +k8s:openapi-gen=true
type AnalyticsDbDaemonSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AnalyticsDbDaemonSetSpec   `json:"spec,omitempty"`
	Status AnalyticsDbDaemonSetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AnalyticsDbDaemonSetList contains a list of AnalyticsDbDaemonSet
type AnalyticsDbDaemonSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AnalyticsDbDaemonSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AnalyticsDbDaemonSet{}, &AnalyticsDbDaemonSetList{})
}
