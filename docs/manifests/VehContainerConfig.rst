VehContainerConfig
==================

set environment config in Vehicle container.

You can set environment variables or mount paths of deployment nodes in
a specific container.

This will be necessary when using features such as X11 forwarding.

You can write and use a custom resource according to your desired
requirements.

============= ========== ==== =====================================
Attribute     Value Type List Kind
============= ========== ==== =====================================
piccoloVolume object     true VehContainerConfig.attr
volume        string     true VehContainerConfig.PiccoloVolume.attr
volumeMount   string     true VehContainerConfig.PiccoloVolume.attr
piccoloEnv    string     true VehContainerConfig.attr
============= ========== ==== =====================================

Attribute Description
---------------------

.. _vehcontainerconfig-2:

*VehContainerConfig*
~~~~~~~~~~~~~~~~~~~~

--------------

piccoloVolume
^^^^^^^^^^^^^

-  .. rubric:: Value Type
      :name: value-type

   object

-  .. rubric:: List
      :name: list

   true

-  .. rubric:: Kind
      :name: kind

   VehContainerConfig.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option

   optional

-  .. rubric:: Description
      :name: description

   The information of the volume that is intended to be used in the pod.

--------------

volume
^^^^^^

-  .. rubric:: Value Type
      :name: value-type-2

   string

-  .. rubric:: List
      :name: list-2

   true(when use PiccoloVolume)

-  .. rubric:: Kind
      :name: kind-2

   VehContainerConfig.PiccoloVolume.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option-2

   true

-  .. rubric:: Description
      :name: description-2

   The path on the host to mount in the container. Inherits the volume
   from the existing orchestrator.

--------------

volumeMount
^^^^^^^^^^^

-  .. rubric:: Value Type
      :name: value-type-3

   string

-  .. rubric:: List
      :name: list-3

   true(when use PiccoloVolume)

-  .. rubric:: Kind
      :name: kind-3

   VehContainerConfig.PiccoloVolume.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option-3

   true

-  .. rubric:: Description
      :name: description-3

   The mount path of the target container. Inherits the volumeMount from
   the existing orchestrator.

--------------

piccoloEnv
^^^^^^^^^^

-  .. rubric:: Value Type
      :name: value-type-4

   string

-  .. rubric:: List
      :name: list-4

   true

-  .. rubric:: Kind
      :name: kind-4

   VehContainerConfig.attr

-  .. rubric:: Mandatory option
      :name: mandatory-option-4

   optional

-  .. rubric:: Description
      :name: description-4

   Setting container environment variables. Inherits the env from the
   existing orchestrator.
