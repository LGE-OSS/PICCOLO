/*
Copyright 2023.

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

package controller

import (
	"context"
	"strconv"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	crdv1alpha1 "crd.piccolo.org/piccolo/api/v1alpha1"
)

// VehPodReconciler reconciles a VehPod object
type VehPodReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=crd.piccolo.org,resources=vehpods,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=crd.piccolo.org,resources=vehpods/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=crd.piccolo.org,resources=vehpods/finalizers,verbs=update
//+kubebuilder:rbac:groups=crd.piccolo.org,resources=devres,verbs=get;list;watch;update
//+kubebuilder:rbac:groups=crd.piccolo.org,resources=vehcontainerconfigs,verbs=get;list;watch;update
//+kubebuilder:rbac:groups="",resources=pods,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the VehPod object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.4/pkg/reconcile
func (r *VehPodReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	controllerLog := log.FromContext(ctx)
	var TAG string = "Custom Controller for Piccolo Pod"

	var vehPod crdv1alpha1.VehPod
	if err := r.Get(ctx, req.NamespacedName, &vehPod); err != nil {
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	vapFinalizerName := vehPod.Name + "/finalizer"

	// Check if the pod already exists, if not create a new one
	found := &corev1.Pod{}
	var podNsn types.NamespacedName = types.NamespacedName{
		Name:      vehPod.Name,
		Namespace: vehPod.Spec.AppType + "-application",
	}

	err := r.Get(ctx, podNsn, found)

	if err != nil && errors.IsNotFound(err) {
		ls := map[string]string{"application-type": vehPod.Spec.AppType}
		po := &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      podNsn.Name,
				Namespace: podNsn.Namespace,
				Labels:    ls,
			},
			Spec: vehPod.Spec.PodSpec,
		}

		for index, deviceInfo := range vehPod.Spec.DevResInfo {
			nsn := req.NamespacedName
			nsn.Name = deviceInfo.DevName

			var piccoloDevice crdv1alpha1.DevRes
			if err := r.Get(ctx, nsn, &piccoloDevice); err != nil {
				return ctrl.Result{}, client.IgnoreNotFound(err)
			}
			pathName := strings.Replace(deviceInfo.DevName, ".", "-", -1) + "-mount-path-"

			po.Spec.Volumes = append(po.Spec.Volumes,
				corev1.Volume{
					Name: pathName + strconv.Itoa(index),
					VolumeSource: corev1.VolumeSource{
						HostPath: &corev1.HostPathVolumeSource{
							Path: piccoloDevice.Spec.Path,
						},
					},
				},
			)

			po.Spec.Containers[0].VolumeMounts = append(po.Spec.Containers[0].VolumeMounts,
				corev1.VolumeMount{
					Name:      pathName + strconv.Itoa(index),
					MountPath: deviceInfo.DevDestPath,
				},
			)
		}

		for _, functionInfo := range vehPod.Spec.VehContainerConfigInfo {
			nsn := req.NamespacedName
			nsn.Name = functionInfo

			var piccoloFunction crdv1alpha1.VehContainerConfig
			if err := r.Get(ctx, nsn, &piccoloFunction); err != nil {
				return ctrl.Result{}, client.IgnoreNotFound(err)
			}

			for _, piccoloVolume := range piccoloFunction.Spec.PiccoloVolumes {
				po.Spec.Volumes = append(po.Spec.Volumes, piccoloVolume.Volume)
				po.Spec.Containers[0].VolumeMounts = append(po.Spec.Containers[0].VolumeMounts, piccoloVolume.VolumeMount)
			}

			po.Spec.Containers[0].Env = append(po.Spec.Containers[0].Env, piccoloFunction.Spec.PiccoloEnv...)
		}

		err = r.Create(ctx, po)
		if err != nil {
			return ctrl.Result{}, err
		}

		if vehPod.ObjectMeta.DeletionTimestamp.IsZero() &&
			!controllerutil.ContainsFinalizer(&vehPod, vapFinalizerName) {
			controllerutil.AddFinalizer(&vehPod, vapFinalizerName)
			if err := r.Update(ctx, &vehPod); err != nil {
				return ctrl.Result{}, err
			}
			controllerLog.Info(TAG, "create pod and add finalizer to vap CR", podNsn)
		}
		return ctrl.Result{Requeue: true}, nil
	} else if err == nil {
		if !vehPod.ObjectMeta.DeletionTimestamp.IsZero() &&
			controllerutil.ContainsFinalizer(&vehPod, vapFinalizerName) {
			if err := r.Delete(ctx, found); err == nil {
				controllerLog.Info(TAG, "delete pod and remove finalizer", podNsn)
				controllerutil.RemoveFinalizer(&vehPod, vapFinalizerName)
				if err := r.Update(ctx, &vehPod); err != nil {
					return ctrl.Result{}, err
				}
			}
		}
	}
	return ctrl.Result{}, err
}

// SetupWithManager sets up the controller with the Manager.
func (r *VehPodReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&crdv1alpha1.VehPod{}).
		Complete(r)
}
