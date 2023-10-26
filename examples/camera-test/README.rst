Camera test
===========

| This example makes an application using a USB camera.
| The key point is that the application developer does not need to know
  any detailed information such as camera mount path or environment
  variables for X11 output.
| When the app developer selects only the predefined Device/Function CR,
  the Piccolo controller automatically injects the configuration
  information into the application Pod.

|image1|

Running on the cluster
----------------------

In this example, we are using Ubuntu and the camera path is
``/dev/video0``.

Build the camera test application:
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

.. code:: 

   cd examples/camera-test
   docker build -t <some-registry>/<APP_IMG_NAME>:tag .
   docker push <some-registry>/<APP_IMG_NAME>:tag

Modify your camera device path:
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

In ``examples/camera-test/yaml/cr_devres_camera.yaml``,

.. code:: 

   spec:
     path: /dev/video0

Build and push your controller image:
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

.. code:: 

   make docker-build docker-push IMG=<some-registry>/piccolo:tag

**Note:** Default value of ``IMG`` is ``controller:latest``

Deploy CRDs and the controller to the cluster:
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

.. code:: 

   make install
   make deploy IMG=<some-registry>/piccolo:tag

Or you can use dry-run

.. code:: 

   make dry-run IMG=<some-registry>/piccolo:tag
   kubectl apply -f dry-run/

Modify your test application image name:
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

In ``examples/camera-test/yaml/cr_vehpod_cameratest.yaml``,

.. code:: 

     podSpec:
       hostNetwork: true
       containers:
       - name: camera-app-pod
         image: <some-registry>/<APP_IMG_NAME>:tag

Add node label or delete nodeSelector in yaml
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

In ``examples/camera-test/yaml/cr_vehpod_cameratest.yaml``,

.. code:: 

       nodeSelector:
         name: "master"

You need to add label or delete nodeSelector in yaml.

|image2|

Optional: Modify your X11 setting
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

In ``examples/camera-test/yaml/cr_vehcontainerconfig_x11.yaml``,

.. code:: 

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
         # TODO: ${DISPLAY} is not converted in this file.
         value: "localhost:11.0"

**Note:** This step is for display output, and you can check the result
in CLI even without this.

Install instances of Custom Resources:
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

.. code:: 

   kubectl apply -f examples/camera-test/yaml/

Check the result:
~~~~~~~~~~~~~~~~~

.. code:: 

   kubectl get po -A

You can see the result like below.

|image3|

Undeploy CRDs and controller:
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

UnDeploy the controller from the cluster:

.. code:: bash

   make undeploy
   make uninstall

If you use dry-run command,

.. code:: 

   kubectl delete -f dry-run/

.. |image1| image:: /examples/camera-test/img/diagram-camera-test.png
.. |image2| image:: /examples/camera-test/img/label.png
.. |image3| image:: /examples/camera-test/img/pod-by-controller.png
