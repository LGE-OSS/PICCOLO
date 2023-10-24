// SPDX-FileCopyrightText: Copyright 2023 LG Electronics Inc.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
)

type DevResInfo struct {
	// The name of the Custom Resource Object of the Device you are trying to access.
	DevName string `json:"devName"`

	// device mount path.
	DevDestPath string `json:"devDestPath"`
}

type PiccoloVolume struct {
	// The path on the host to mount in the container. Inherits the volume from the existing orchestrator.
	Volume corev1.Volume `json:"volume"`

	// The mount path of the target container. Inherits the volumeMount from the existing orchestrator.
	VolumeMount corev1.VolumeMount `json:"volumeMount"`
}
