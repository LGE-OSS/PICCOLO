VehDeployment
=============

Deployment Object for Vehicle.

====================== ========== ===== =============================
Attribute              Value Type List  Kind
====================== ========== ===== =============================
devResInfo             cr         true  VehDeployment.attr
devName                string     false VehDeployment.DevResInfo.attr
devDestPath            string     false VehDeployment.DevResInfo.attr
devContainerIdx        int        false VehDeployment.attr
vehContainerConfigInfo cr         true  VehDeployment.attr
vehContainerConfigIdx  int        false VehDeployment.attr
appType                string     false VehDeployment.attr
deploymentPolicy       string     false VehDeployment.attr
podSpec                string     false VehDeployment.attr
====================== ========== ===== =============================

Attribute Description
---------------------

.. _vehdeployment-2:

*VehDeployment*
~~~~~~~~~~~~~~~

--------------

devResInfo
^^^^^^^^^^

-  .. rubric:: Value Type
      :name: value-type

   cr(custom resource)

-  .. rubric:: List
      :name: list

   true

-  .. rubric:: Kind
      :name: kind

   VehDeployment.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option

   optional

-  .. rubric:: Description
      :name: description

   The field written when you want to request device resources.

   When you want to use a device in a specific container, you can use
   the devices provided by custom resource.

**cr_device_camera.yaml**

.. code:: yaml

   apiVersion: crd.piccolo.org/v1alpha1
   kind: DevRes
   metadata:
     name: cam-interior-driver
   spec:
     path: /dev/video0

--------------

devName
^^^^^^^

-  .. rubric:: Value Type
      :name: value-type-2

   string

-  .. rubric:: List
      :name: list-2

   false

-  .. rubric:: Kind
      :name: kind-2

   VehDeployment.DevResInfo.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option-2

   true (when use DevResInfo)

-  .. rubric:: Description
      :name: description-2

   The name of the Custom Resource Object of the Device you are trying
   to access.

--------------

devDestPath
^^^^^^^^^^^

-  .. rubric:: Value Type
      :name: value-type-3

   string

-  .. rubric:: List
      :name: list-3

   false

-  .. rubric:: Kind
      :name: kind-3

   VehDeployment.DevResInfo.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option-3

   true(when use DevResInfo)

-  .. rubric:: Description
      :name: description-3

   device mount path.

--------------

devContainerIdx
^^^^^^^^^^^^^^^

-  .. rubric:: Value Type
      :name: value-type-4

   int

-  .. rubric:: List
      :name: list-4

   false

-  .. rubric:: Kind
      :name: kind-4

   VehDeployment.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option-4

   true(when use DevResInfo)

-  .. rubric:: Description
      :name: description-4

   Container index to use the device.

--------------

vehContainerConfigInfo
^^^^^^^^^^^^^^^^^^^^^^

-  .. rubric:: Value Type
      :name: value-type-5

   cr(customresource)

-  .. rubric:: List
      :name: list-5

   false

-  .. rubric:: Kind
      :name: kind-5

   VehDeployment.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option-5

   optional

-  .. rubric:: Description
      :name: description-5

   It mounts the required host path or configures the environment in the
   Pod.

   You can set environment variables or mount paths of deployment nodes
   in a specific container. This will be necessary when using features
   such as X11 forwarding.

   You can write and use a custom resource according to your desired
   requirements.

**cr_function_x11.yaml**

.. code:: yaml

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

--------------

vehContainerConfigIdx
^^^^^^^^^^^^^^^^^^^^^

-  .. rubric:: Value Type
      :name: value-type-6

   int

-  .. rubric:: List
      :name: list-6

   false

-  .. rubric:: Kind
      :name: kind-6

   VehDeployment.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option-6

   true(when use VehContainerConfigInfo)

-  .. rubric:: Description
      :name: description-6

   Container index to use VehContainerConfig.

--------------

appType
^^^^^^^

-  .. rubric:: Value Type
      :name: value-type-7

   string

-  .. rubric:: List
      :name: list-7

   false

-  .. rubric:: Kind
      :name: kind-7

   VehDeployment.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option-7

   true

-  .. rubric:: Description
      :name: description-7

   | It sets the permissions for the Pod based on the input of the type
     of system or 3rd-party.
   | When this value is set, it can restrict or authorize the usage of
     computing resources such as CPU and memory, as well as access
     permissions for custom resources, according to the System Policy.

--------------

deploymentPolicy
^^^^^^^^^^^^^^^^

-  .. rubric:: Value Type
      :name: value-type-8

   string

-  .. rubric:: List
      :name: list-8

   false

-  .. rubric:: Kind
      :name: kind-8

   VehDeployment.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option-8

   true

-  .. rubric:: Description
      :name: description-8

   | Set the structure of containers included on this pod.
   | Possible values include PassiveRedundant, ActiveRedundant,
     NModulerRedundant, and Monitoring.

--------------

podSpec
^^^^^^^

-  .. rubric:: Value Type
      :name: value-type-9

   string

-  .. rubric:: List
      :name: list-9

   false

-  .. rubric:: Kind
      :name: kind-9

   VehDeployment.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option-9

   true

-  .. rubric:: Description
      :name: description-9

   The contents of the pod that will be executed are described, and
   based on this, the pod is executed. It inherits the podspec of the
   existing orchestrator.
