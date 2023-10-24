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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// VehPodSpec defines the desired state of VehPod
type VehPodSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// The field written when you want to request device resources.
	// When you want to use a device in a specific container, you can use the devices provided by custom resource.
	DevResInfo []DevResInfo `json:"devResInfo,omitempty"`

	// Container index to use the device.
	DevContainerIdx int `json:"devContainerIdx,omitempty"`

	// It mounts the required host path or configures the environment in the Pod.
	// You can set environment variables or mount paths of deployment nodes in a specific container. This will be necessary when using features such as X11 forwarding.
	// You can write and use a custom resource according to your desired requirements.
	VehContainerConfigInfo []string `json:"vehContainerConfigInfo,omitempty"`

	// Container index to use VehContainerConfig.
	VehContainerConfigIdx int `json:"vehContainerConfigIdx,omitempty"`

	// It sets the permissions for the Pod based on the input of the type of system or 3rd-party.
	// When this value is set, it can restrict or authorize the usage of computing resources such as CPU and memory, as well as access permissions for custom resources, according to the System Policy.
	AppType string `json:"appType,omitempty"`

	// Set the structure of containers included on this pod.
	// Possible values include PassiveRedundant, ActiveRedundant, NModulerRedundant, and Monitoring.
	DeploymentPolicy string `json:"deploymentPolicy,omitempty"`

	// The contents of the pod that will be executed are described, and based on this, the pod is executed. It inherits the podspec of the existing orchestrator.
	PodSpec corev1.PodSpec `json:"podSpec"`
}

// VehPodStatus defines the observed state of VehPod
type VehPodStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// VehPod is the Schema for the vehpods API
type VehPod struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VehPodSpec   `json:"spec,omitempty"`
	Status VehPodStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// VehPodList contains a list of VehPod
type VehPodList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VehPod `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VehPod{}, &VehPodList{})
}
