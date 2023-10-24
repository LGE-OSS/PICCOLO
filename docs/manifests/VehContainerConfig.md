# VehContainerConfig

set environment config in Vehicle container.

You can set environment variables or mount paths of deployment nodes in a specific container. 

This will be necessary when using features such as X11 forwarding.

You can write and use a custom resource according to your desired requirements.

| Attribute     | Value Type | List | Kind                                  |
| ------------- | ---------- | ---- | ------------------------------------- |
| piccoloVolume | object     | true | VehContainerConfig.attr               |
| volume        | string     | true | VehContainerConfig.PiccoloVolume.attr |
| volumeMount   | string     | true | VehContainerConfig.PiccoloVolume.attr |
| piccoloEnv    | string     | true | VehContainerConfig.attr               |

### Attribute Description

#### *VehContainerConfig*

------

##### piccoloVolume

- ######  Value Type

  object

- ###### List

  true

- ###### Kind

  VehContainerConfig.attr

- ###### Mandatory option

  optional

- ###### Description

  The information of the volume that is intended to be used in the pod.

------

##### volume

- ######  Value Type

  string

- ###### List

  true(when use PiccoloVolume)

- ###### Kind

  VehContainerConfig.PiccoloVolume.attr

- ###### Mandatory option

  true

- ###### Description

  The path on the host to mount in the container. Inherits the volume from the existing orchestrator.

------

##### volumeMount

- ######  Value Type

  string

- ###### List

  true(when use PiccoloVolume)

- ###### Kind

  VehContainerConfig.PiccoloVolume.attr

- ###### Mandatory option

  true

- ###### Description

  The mount path of the target container. Inherits the volumeMount from the existing orchestrator.

------

##### piccoloEnv

- ######  Value Type

  string

- ###### List

  true

- ###### Kind

  VehContainerConfig.attr

- ###### Mandatory option

  optional

- ###### Description

  Setting container environment variables. Inherits the env from the existing orchestrator.
