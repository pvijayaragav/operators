// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.AnalyticsDaemonSet":               schema_pkg_apis_contrailoperators_v1alpha1_AnalyticsDaemonSet(ref),
		"github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.AnalyticsDaemonSetSpec":           schema_pkg_apis_contrailoperators_v1alpha1_AnalyticsDaemonSetSpec(ref),
		"github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.AnalyticsDaemonSetStatus":         schema_pkg_apis_contrailoperators_v1alpha1_AnalyticsDaemonSetStatus(ref),
		"github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.AnalyticsDbDaemonSet":             schema_pkg_apis_contrailoperators_v1alpha1_AnalyticsDbDaemonSet(ref),
		"github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.AnalyticsDbDaemonSetSpec":         schema_pkg_apis_contrailoperators_v1alpha1_AnalyticsDbDaemonSetSpec(ref),
		"github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.AnalyticsDbDaemonSetStatus":       schema_pkg_apis_contrailoperators_v1alpha1_AnalyticsDbDaemonSetStatus(ref),
		"github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.ConfigDbDaemonSet":                schema_pkg_apis_contrailoperators_v1alpha1_ConfigDbDaemonSet(ref),
		"github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.ConfigDbDaemonSetSpec":            schema_pkg_apis_contrailoperators_v1alpha1_ConfigDbDaemonSetSpec(ref),
		"github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.ConfigDbDaemonSetStatus":          schema_pkg_apis_contrailoperators_v1alpha1_ConfigDbDaemonSetStatus(ref),
		"github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.ControllerConfigDaemonSet":        schema_pkg_apis_contrailoperators_v1alpha1_ControllerConfigDaemonSet(ref),
		"github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.ControllerConfigDaemonSetSpec":    schema_pkg_apis_contrailoperators_v1alpha1_ControllerConfigDaemonSetSpec(ref),
		"github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.ControllerConfigDaemonSetStatus":  schema_pkg_apis_contrailoperators_v1alpha1_ControllerConfigDaemonSetStatus(ref),
		"github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.ControllerControlDaemonSet":       schema_pkg_apis_contrailoperators_v1alpha1_ControllerControlDaemonSet(ref),
		"github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.ControllerControlDaemonSetSpec":   schema_pkg_apis_contrailoperators_v1alpha1_ControllerControlDaemonSetSpec(ref),
		"github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.ControllerControlDaemonSetStatus": schema_pkg_apis_contrailoperators_v1alpha1_ControllerControlDaemonSetStatus(ref),
		"github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.VrouterDaemonSet":                 schema_pkg_apis_contrailoperators_v1alpha1_VrouterDaemonSet(ref),
		"github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.VrouterDaemonSetSpec":             schema_pkg_apis_contrailoperators_v1alpha1_VrouterDaemonSetSpec(ref),
		"github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.VrouterDaemonSetStatus":           schema_pkg_apis_contrailoperators_v1alpha1_VrouterDaemonSetStatus(ref),
	}
}

func schema_pkg_apis_contrailoperators_v1alpha1_AnalyticsDaemonSet(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AnalyticsDaemonSet is the Schema for the analyticsdaemonsets API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.AnalyticsDaemonSetSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.AnalyticsDaemonSetStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.AnalyticsDaemonSetSpec", "github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.AnalyticsDaemonSetStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_contrailoperators_v1alpha1_AnalyticsDaemonSetSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AnalyticsDaemonSetSpec defines the desired state of AnalyticsDaemonSet",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_contrailoperators_v1alpha1_AnalyticsDaemonSetStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AnalyticsDaemonSetStatus defines the observed state of AnalyticsDaemonSet",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_contrailoperators_v1alpha1_AnalyticsDbDaemonSet(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AnalyticsDbDaemonSet is the Schema for the analyticsdbdaemonsets API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.AnalyticsDbDaemonSetSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.AnalyticsDbDaemonSetStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.AnalyticsDbDaemonSetSpec", "github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.AnalyticsDbDaemonSetStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_contrailoperators_v1alpha1_AnalyticsDbDaemonSetSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AnalyticsDbDaemonSetSpec defines the desired state of AnalyticsDbDaemonSet",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_contrailoperators_v1alpha1_AnalyticsDbDaemonSetStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "AnalyticsDbDaemonSetStatus defines the observed state of AnalyticsDbDaemonSet",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_contrailoperators_v1alpha1_ConfigDbDaemonSet(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ConfigDbDaemonSet is the Schema for the configdbdaemonsets API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.ConfigDbDaemonSetSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.ConfigDbDaemonSetStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.ConfigDbDaemonSetSpec", "github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.ConfigDbDaemonSetStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_contrailoperators_v1alpha1_ConfigDbDaemonSetSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ConfigDbDaemonSetSpec defines the desired state of ConfigDbDaemonSet",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_contrailoperators_v1alpha1_ConfigDbDaemonSetStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ConfigDbDaemonSetStatus defines the observed state of ConfigDbDaemonSet",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_contrailoperators_v1alpha1_ControllerConfigDaemonSet(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ControllerConfigDaemonSet is the Schema for the controllerconfigdaemonsets API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.ControllerConfigDaemonSetSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.ControllerConfigDaemonSetStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.ControllerConfigDaemonSetSpec", "github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.ControllerConfigDaemonSetStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_contrailoperators_v1alpha1_ControllerConfigDaemonSetSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ControllerConfigDaemonSetSpec defines the desired state of ControllerConfigDaemonSet",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_contrailoperators_v1alpha1_ControllerConfigDaemonSetStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ControllerConfigDaemonSetStatus defines the observed state of ControllerConfigDaemonSet",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_contrailoperators_v1alpha1_ControllerControlDaemonSet(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ControllerControlDaemonSet is the Schema for the controllercontroldaemonsets API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.ControllerControlDaemonSetSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.ControllerControlDaemonSetStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.ControllerControlDaemonSetSpec", "github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.ControllerControlDaemonSetStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_contrailoperators_v1alpha1_ControllerControlDaemonSetSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ControllerControlDaemonSetSpec defines the desired state of ControllerControlDaemonSet",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_contrailoperators_v1alpha1_ControllerControlDaemonSetStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ControllerControlDaemonSetStatus defines the observed state of ControllerControlDaemonSet",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_contrailoperators_v1alpha1_VrouterDaemonSet(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "VrouterDaemonSet is the Schema for the vrouterdaemonsets API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.VrouterDaemonSetSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.VrouterDaemonSetStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.VrouterDaemonSetSpec", "github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1.VrouterDaemonSetStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_contrailoperators_v1alpha1_VrouterDaemonSetSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "VrouterDaemonSetSpec defines the desired state of VrouterDaemonSet",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_contrailoperators_v1alpha1_VrouterDaemonSetStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "VrouterDaemonSetStatus defines the observed state of VrouterDaemonSet",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}