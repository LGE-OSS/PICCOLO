# VehDeployment

Deployment Object for Vehicle.

| Attribute              | Value Type | List  | Kind                          |
| ---------------------- | ---------- | ----- | ----------------------------- |
| devResInfo             | cr         | true  | VehDeployment.attr            |
| devName                | string     | false | VehDeployment.DevResInfo.attr |
| devDestPath            | string     | false | VehDeployment.DevResInfo.attr |
| devContainerIdx        | int        | false | VehDeployment.attr            |
| vehContainerConfigInfo | cr         | true  | VehDeployment.attr            |
| vehContainerConfigIdx  | int        | false | VehDeployment.attr            |
| appType                | string     | false | VehDeployment.attr            |
| deploymentPolicy       | string     | false | VehDeployment.attr            |
| podSpec                | string     | false | VehDeployment.attr            |

### Attribute Description

#### *VehDeployment*

------

##### devResInfo

- ######  Value Type

  cr(custom resource)

- ###### List

  true

- ###### Kind

  VehDeployment.attr

- ###### Mandatory option

  optional

- ###### Description

  The field written when you want to request device resources.

  When you want to use a device in a specific container, you can use the devices provided by custom resource.



**cr_device_camera.yaml**

```yaml
apiVersion: crd.piccolo.org/v1alpha1
kind: DevRes
metadata:
  name: cam-interior-driver
spec:
  path: /dev/video0
```

------

##### devName

- ######  Value Type

  string

- ###### List

  false

- ###### Kind

  VehDeployment.DevResInfo.attr

- ###### Mandatory option

  true (when use DevResInfo)

- ###### Description

  The name of the Custom Resource Object of the Device you are trying to access.

------
##### devDestPath

- ######  Value Type

  string

- ###### List

  false

- ###### Kind

  VehDeployment.DevResInfo.attr

- ###### Mandatory option

  true(when use DevResInfo)

- ###### Description

  device mount path.

------
##### devContainerIdx

- ######  Value Type

  int

- ###### List

  false

- ###### Kind

  VehDeployment.attr

- ###### Mandatory option

  true(when use DevResInfo)

- ###### Description

  Container index to use the device.

------

##### vehContainerConfigInfo

- ######  Value Type

  cr(customresource)

- ###### List

  false

- ###### Kind

  VehDeployment.attr

- ###### Mandatory option

  optional

- ###### Description

  It mounts the required host path or configures the environment in the Pod.

  You can set environment variables or mount paths of deployment nodes in a specific container. This will be necessary when using features such as X11 forwarding.

  You can write and use a custom resource according to your desired requirements.



**cr_function_x11.yaml**

```yaml
apiVersion: crd.piccolo.org/v1alpha1
kind: VehContainerConfig
metadata:
  name: vehicle-x11-window
spec:
  piccoloVolumes:
    - volume:
        name: x11-unix
        hostPath:
          path: /tmp/.X11-unix
      volumeMount:
        name: x11-unix
        mountPath: /tmp/.X11-unix
    - volume:
        name: x11-auth
        hostPath:
          path: /root/.Xauthority
      volumeMount:
        name: x11-auth
        mountPath: /root/.Xauthority
  piccoloEnv:
    - name: DISPLAY
      value: "localhost:11.0"
```

------

##### vehContainerConfigIdx

- ######  Value Type

  int

- ###### List

  false

- ###### Kind

  VehDeployment.attr

- ###### Mandatory option

  true(when use VehContainerConfigInfo)

- ###### Description

  Container index to use VehContainerConfig.

------

##### appType

- ######  Value Type

  string

- ###### List

  false

- ###### Kind

  VehDeployment.attr

- ###### Mandatory option

  true

- ###### Description

  It sets the permissions for the Pod based on the input of the type of system or 3rd-party.
  When this value is set, it can restrict or authorize the usage of computing resources such as CPU and memory, as well as access permissions for custom resources, according to the System Policy.

------

##### deploymentPolicy

- ######  Value Type

  string

- ###### List

  false

- ###### Kind

  VehDeployment.attr

- ###### Mandatory option

  true

- ###### Description

  Set the structure of containers included on this pod.
  Possible values include PassiveRedundant, ActiveRedundant, NModulerRedundant, and Monitoring.

------

##### podSpec

- ######  Value Type

  string

- ###### List

  false

- ###### Kind

  VehDeployment.attr

- ###### Mandatory option

  true

- ###### Description

  The contents of the pod that will be executed are described, and based on this, the pod is executed. It inherits the podspec of the existing orchestrator.

