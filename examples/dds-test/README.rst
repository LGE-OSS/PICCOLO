DDS test
========

| This example makes two application pod using the DDS communication.
| The key point is that the application developer does not need to know
  any detailed information about environment variables for DDS
  configuration.

|image1|

Running on the cluster
----------------------

Build the dds test application:
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

.. code:: 

   cd examples/dds-test
   docker build -t <some-registry>/<APP_IMG_NAME>:tag -f Dockerfile_<ping or pong>
   docker push <some-registry>/<APP_IMG_NAME>:tag

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

In ``examples/dds-test/yaml/cr_vehpod_ddstest_ping.yaml``,

.. code:: 

     podSpec:
       containers:
       - name: ping-container
         image: <some-registry>/<APP_IMG_NAME>:tag

In ``examples/dds-test/yaml/cr_vehpod_ddstest_pong.yaml``,

.. code:: 

     podSpec:
       containers:
       - name: pong-container
         image: <some-registry>/<APP_IMG_NAME>:tag

Add node label or delete nodeSelector in yaml
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

In two VehPod yaml,

.. code:: 

       nodeSelector:
         name: "master"

You need to add label or delete nodeSelector in yaml.

|image2|

Modify your dds setting
~~~~~~~~~~~~~~~~~~~~~~~

In ``examples/dds-test/yaml/cr_vehcontainerconfig_dds.yaml``,

.. code:: 

   spec:
     piccoloEnv:
       - name: CYCLONEDDS_URI
         value: '<CycloneDDS><Domain Id="100"></Domain></CycloneDDS>'

**Note:** domain id value can be changed.

Install instances of Custom Resources:
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

.. code:: 

   kubectl apply -f examples/dds-test/yaml/

Check the result:
~~~~~~~~~~~~~~~~~

.. code:: 

   kubectl get po -A

You can see the result like below.

|image3|

.. code:: 

   kubectl logs -f <POD_NAME> --namespace=<APP_NAMESPACE>

You can see the result like below.

|image4|

Undeploy CRDs and controller:
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

UnDeploy the controller from the cluster:

.. code:: bash

   make undeploy
   make uninstall

If you use dry-run command,

.. code:: 

   kubectl delete -f dry-run/

.. |image1| image:: /examples/dds-test/img/diagram-dds-test.png
.. |image2| image:: /examples/camera-test/img/label.png
.. |image3| image:: /examples/dds-test/img/pod-by-controller.png
.. |image4| image:: /examples/dds-test/img/domain-id-log.png
