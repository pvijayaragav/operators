package analyticsdbdaemonset

import (
	"context"

	contrailoperatorsv1alpha1 "github.com/contrail-operators/contrail-all-in-one-operator/pkg/apis/contrailoperators/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_analyticsdbdaemonset")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new AnalyticsDbDaemonSet Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileAnalyticsDbDaemonSet{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("analyticsdbdaemonset-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource AnalyticsDbDaemonSet
	err = c.Watch(&source.Kind{Type: &contrailoperatorsv1alpha1.AnalyticsDbDaemonSet{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner AnalyticsDbDaemonSet
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &contrailoperatorsv1alpha1.AnalyticsDbDaemonSet{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileAnalyticsDbDaemonSet{}

// ReconcileAnalyticsDbDaemonSet reconciles a AnalyticsDbDaemonSet object
type ReconcileAnalyticsDbDaemonSet struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

func (r *ReconcileAnalyticsDbDaemonSet) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling AnalyticsDbDaemonSet")
	instance := &contrailoperatorsv1alpha1.AnalyticsDbDaemonSet{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}

	nodemgrcm := newNodeMgrCmForDS(instance)
	agentcm := newAgentCmForDS(instance)
	ds := newDSForCR(instance)

	if err := controllerutil.SetControllerReference(instance, nodemgrcm, r.scheme); err != nil {
		return reconcile.Result{}, err
	}
	found_cm := &corev1.ConfigMap{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: nodemgrcm.Name, Namespace: nodemgrcm.Namespace}, found_cm)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new ConfigMap", "ConfigMap.Namespace", nodemgrcm.Namespace, "ConfigMap.Name", nodemgrcm.Name)
		err = r.client.Create(context.TODO(), nodemgrcm)
		if err != nil {
			return reconcile.Result{}, err
		}
	} else if err != nil {
		return reconcile.Result{}, err
	} else {
	reqLogger.Info("Skip reconcile: ConfigMap already exists", "ConfigMap.Namespace", found_cm.Namespace, "ConfigMap.Name", found_cm.Name)
	}

	if err := controllerutil.SetControllerReference(instance, agentcm, r.scheme); err != nil {
		return reconcile.Result{}, err
	}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: agentcm.Name, Namespace: agentcm.Namespace}, found_cm)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new ConfigMap", "ConfigMap.Namespace", agentcm.Namespace, "ConfigMap.Name", agentcm.Name)
		err = r.client.Create(context.TODO(), agentcm)
		if err != nil {
			return reconcile.Result{}, err
		}
	} else if err != nil {
		return reconcile.Result{}, err
	} else {
	reqLogger.Info("Skip reconcile: ConfigMap already exists", "ConfigMap.Namespace", found_cm.Namespace, "ConfigMap.Name", found_cm.Name)
	}

	if err := controllerutil.SetControllerReference(instance, ds, r.scheme); err != nil {
		return reconcile.Result{}, err
	}
	// Check if this Pod already exists
	found := &appsv1.DaemonSet{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: ds.Name, Namespace: ds.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new DS", "DS.Namespace", ds.Namespace, "DS.Name", ds.Name)
		err = r.client.Create(context.TODO(), ds)
		if err != nil {
			return reconcile.Result{}, err
		}
		// Pod created successfully - don't requeue
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}
	// Pod already exists - don't requeue
	reqLogger.Info("Skip reconcile: Pod already exists", "Pod.Namespace", found.Namespace, "Pod.Name", found.Name)

	return reconcile.Result{}, nil
}

func newNodeMgrCmForDS(cr *contrailoperatorsv1alpha1.AnalyticsDbDaemonSet) *corev1.ConfigMap{
				return &corev1.ConfigMap{
								ObjectMeta: metav1.ObjectMeta{
												Name:           cr.Name + "-nodemgr-env",
												Namespace:      cr.Namespace,
								},
								Data: map[string]string{
												"DOCKER_HOST": "unix://mnt/docker.sock",
								},
				}
}

func getContrailNodes(nodes []corev1.Node) []string{
	var nodeNames []string
	for _, node := range nodes {
		nodeLabels := node.ObjectMeta.Labels
		for _, label := range nodeLabels {
			if "opencontrail.org/controller" == label {
				nodeNames = append(nodeNames, node.ObjectMeta.Name)
			}
		}
	}
	return nodeNames
}

func newAgentCmForDS(cr *contrailoperatorsv1alpha1.AnalyticsDbDaemonSet) *corev1.ConfigMap{
	contrailNodes := cr.Spec.ContrailMasters
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:		cr.Name + "-agent-env",
			Namespace:	cr.Namespace,
		},
		Data: map[string]string{
			"AAA_MODE":	"no-auth",
			"AUTH_MODE":	"noauth",
			"CLOUD_ORCHESTRATOR": "kubernetes",
			"CONFIG_NODES": contrailNodes,
			"CONFIGDB_NODES": contrailNodes,
			"CONTROL_NODES": contrailNodes,
			"CONTROLLER_NODES": contrailNodes,
			"KAFKA_NODES": contrailNodes,
			"LOG_LEVEL": "SYS_NOTICE",
			"METADATA_PROXY_SECRET": "contrail",
			"RABBITMQ_NODES": contrailNodes,
			"RABBITMQ_NODE_PORT": "5672",
			"REDIS_NODES": contrailNodes,
			"ZOOKEEPER_NODES": contrailNodes,
		},
	}
}

func initContainersForDS(cr *contrailoperatorsv1alpha1.AnalyticsDbDaemonSet) []corev1.Container{

	return []corev1.Container{
		{
			Name:    		"contrail-node-init",
			Image:   		"opencontrailnightly/contrail-node-init",
			Command: 		[]string{"./entrypoint.sh"},
			SecurityContext:	&corev1.SecurityContext{
							Privileged: func(b bool) *bool { return &b }(true),
			},
			Env:			[]corev1.EnvVar{
						{
							Name: "CONTRAIL_STATUS_IMAGE",
							Value: "opencontrailnightly/contrail-status",
						},
						},
			EnvFrom:		[]corev1.EnvFromSource{
							{
								ConfigMapRef: &corev1.ConfigMapEnvSource{
										LocalObjectReference: corev1.LocalObjectReference{Name: cr.Name + "-agent-env"},
								},
							},
						},
			VolumeMounts:		[]corev1.VolumeMount{
						{
							MountPath: "/host/usr/bin",
							Name: "host-usr-bin",
						},
						},
		},
		{
			Name:    "contrail-download-contrail-status-image",
			Image:   "opencontrailnightly/contrail-status",
			ImagePullPolicy: "IfNotPresent",
			Command: []string{"/bin/true"},
		},
		{
			Name:    "contrail-download-contrail-kubernetes-cni-init-image",
			Image:   "opencontrailnightly/contrail-kubernetes-cni-init",
			ImagePullPolicy: "IfNotPresent",
			Command: []string{"/bin/true"},
		},
		{
			Name:    "contrail-download-contrail-vrouter-agent-image",
			Image:   "opencontrailnightly/contrail-vrouter-agent",
			ImagePullPolicy: "IfNotPresent",
			Command: []string{"/bin/true"},
		},
		{
			Name:    "contrail-download-contrail-nodemgr-image",
			Image:   "opencontrailnightly/contrail-nodemgr",
			ImagePullPolicy: "IfNotPresent",
			Command: []string{"/bin/true"},
		},
		{
			Name:    "contrail-vrouter-kernel-init",
			Image:   "opencontrailnightly/contrail-vrouter-kernel-init",
			ImagePullPolicy: "IfNotPresent",
			SecurityContext:	&corev1.SecurityContext{
							Privileged: func(b bool) *bool { return &b }(true),
			},
			EnvFrom:	[]corev1.EnvFromSource{
				{
					ConfigMapRef: &corev1.ConfigMapEnvSource{
							LocalObjectReference: corev1.LocalObjectReference{Name: cr.Name + "-agent-env"},
					},
				},
			},
			VolumeMounts:	[]corev1.VolumeMount{
						{
							MountPath: "/host/usr/bin",
							Name: "host-usr-bin",
						},
						{
							MountPath: "/usr/src",
							Name: "usr-src",
						},
						{
							MountPath: "/lib/modules",
							Name: "lib-modules",
						},
						{
							MountPath: "/etc/sysconfig/network-scripts",
							Name: "network-scripts",
						},
						{
							MountPath: "/host/bin",
							Name: "host-bin",
						},
			},
		},
		{
			Name:	"contrail-vrouter-cni-init",
			Image:	"opencontrailnightly/contrail-kubernetes-cni-init",
			ImagePullPolicy:	"IfNotPresent",
			SecurityContext:	&corev1.SecurityContext{
							Privileged: func(b bool) *bool { return &b }(true),
			},
			EnvFrom:	[]corev1.EnvFromSource{
				{
					ConfigMapRef: &corev1.ConfigMapEnvSource{
							LocalObjectReference: corev1.LocalObjectReference{Name: cr.Name + "-agent-env"},
					},
				},
			},
			VolumeMounts:	[]corev1.VolumeMount{
					{
						MountPath: "/var/lib/contrail",
						Name: "var-lib-contrail",
					},
					{
						MountPath: "/host/etc_cni",
						Name: "etc-cni",
					},
					{
						MountPath: "/host/opt_cni_bin",
						Name: "opt-cni-bin",
					},
					{
						MountPath: "/host/log_cni",
						Name: "var-log-contrail-cni",
					},
					{
						MountPath: "/var/log/contrail",
						Name: "agent-logs",
					},
			},
		},
	}
}

func containersForDS(cr *contrailoperatorsv1alpha1.AnalyticsDbDaemonSet) []corev1.Container{
	return []corev1.Container{
		{
		Name:			"contrail-vrouter-agent",
		Image:   		"opencontrailnightly/contrail-vrouter-agent",
		Command: 		[]string{"./entrypoint.sh"},
		ImagePullPolicy: "IfNotPresent",
		SecurityContext:	&corev1.SecurityContext{
						Privileged: func(b bool) *bool { return &b }(true),
		},
		EnvFrom:		[]corev1.EnvFromSource{
			{
				ConfigMapRef: &corev1.ConfigMapEnvSource{
						LocalObjectReference: corev1.LocalObjectReference{Name: cr.Name + "-agent-env"},
				},
			},
		},
		VolumeMounts:		[]corev1.VolumeMount{
				{
					MountPath: "/dev",
					Name: "dev",
				},
				{
					MountPath: "/etc/sysconfig/network-scripts",
					Name: "network-scripts",
				},
				{
					MountPath: "/host/bin",
					Name: "host-bin",
				},
				{
					MountPath: "/var/log/contrail",
					Name: "agent-logs",
				},
				{
					MountPath: "/usr/src",
					Name: "usr-src",
				},
				{
					MountPath: "/lib/modules",
					Name: "lib-modules",
				},
				{
					MountPath: "/var/lib/contrail",
					Name: "var-lib-contrail",
				},
				{
					MountPath: "/var/crashes",
					Name: "var-crashes",
				},
				{
					MountPath: "/etc/localtime",
					Name: "localtime",
				},
		},
	},
	{
		Name:    		"contrail-agent-nodemgr",
		Image:   		"opencontrailnightly/contrail-nodemgr",
		ImagePullPolicy: "IfNotPresent",
		SecurityContext:	&corev1.SecurityContext{
						Privileged: func(b bool) *bool { return &b }(true),
		},
		Env:			[]corev1.EnvVar{
					{
						Name: "NODE_TYPE",
						Value: "vrouter",
					},
					},
		EnvFrom:		[]corev1.EnvFromSource{
			{
				ConfigMapRef: &corev1.ConfigMapEnvSource{
						LocalObjectReference: corev1.LocalObjectReference{Name: cr.Name + "-agent-env"},
				},
			},
			{
				ConfigMapRef: &corev1.ConfigMapEnvSource{
						LocalObjectReference: corev1.LocalObjectReference{Name: cr.Name + "-nodemgr-env"},
				},
			},
		},
		VolumeMounts:		[]corev1.VolumeMount{
					{
						MountPath: "/var/log/contrail",
		        Name: "agent-logs",
					},
					{
						MountPath: "/mnt",
		        Name: "docker-unix-socket",
		      },
					{
						MountPath: "/etc/localtime",
		        Name: "localtime",
					},
					},
	},
}
}

func volumesForDS() []corev1.Volume{
	return []corev1.Volume{
		{
			Name: "dev",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/dev",
				},
			},
		},
		{
			Name: "network-scripts",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/etc/sysconfig/network-scripts",
				},
			},
		},
		{
			Name: "host-bin",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/bin",
				},
			},
		},
		{
			Name: "docker-unix-socket",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/var/run",
				},
			},
		},
		{
			Name: "usr-src",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/usr/src",
				},
			},
		},
		{
			Name: "lib-modules",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/lib/modules",
				},
			},
		},
		{
			Name: "var-lib-contrail",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/var/lib/contrail",
				},
			},
		},
		{
			Name: "var-crashes",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/var/contrail/crashes",
				},
			},
		},
		{
			Name: "etc-cni",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/etc/cni",
				},
			},
		},
		{
			Name: "opt-cni-bin",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/opt/cni/bin",
				},
			},
		},
		{
			Name: "var-log-contrail-cni",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/var/log/contrail/cni",
				},
			},
		},
		{
			Name: "agent-logs",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/var/log/contrail/agent",
				},
			},
		},
		{
			Name: "host-usr-bin",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/usr/bin",
				},
			},
		},
		{
			Name: "localtime",
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: "/etc/localtime",
				},
			},
		},
	}
}

func newDSForCR(cr *contrailoperatorsv1alpha1.AnalyticsDbDaemonSet) *appsv1.DaemonSet{
    labels := map[string]string{
                "app": cr.Name,
        }
		return &appsv1.DaemonSet{
						ObjectMeta: metav1.ObjectMeta{
							Name:      cr.Name + "-ds",
							Namespace: cr.Namespace,
							Labels:    labels,
						},
						Spec: appsv1.DaemonSetSpec{
							Selector: &metav1.LabelSelector{MatchLabels: labels},
							Template: corev1.PodTemplateSpec{
								ObjectMeta: metav1.ObjectMeta{
									Name:      cr.Name + "-pod-template",
									Namespace: cr.Namespace,
									Labels:    labels,
								},
								Spec: corev1.PodSpec{
									HostNetwork: true,
									Tolerations: []corev1.Toleration{
										{
											Key: "node.kubernetes.io/not-ready",
											Operator: "Exists",
										},
									},
									InitContainers: initContainersForDS(cr),
									Containers: containersForDS(cr),
									Volumes: volumesForDS(),
								},
							},
						},
		}
}
