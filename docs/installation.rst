Installation
============

| There is nothing to install to use Piccolo, but you need to apply CRD
| (custom resource definition) and controller to k8s cluster.
| Piccolo supports two options for this.

-  Kustomize

-  Helm

| If you are new to Piccolo, recommend that you read the
| `Quick Start </docs/quick-start.rst>`__ first, before proceeding
  further.

Kustomize
---------

.. code:: bash

   make deploy IMG=<some-registry>/piccolo:tag

Or

.. code:: 

   make dry-run IMG=<some-registry>/piccolo:tag
   kubectl apply -f dry-run/

Alternatively, if you are familiar with kustomize, you can also use the
following command depending on your situation.

.. code:: 

   kubectl apply -k config/default/

Helm
----

.. code:: bash

   make helm IMG=<some-registry>/piccolo:tag
   helm install <RELEASE-NAME> ./chart
   kubectl apply -f examples/camera-test/yaml

OR

.. code:: 

   kubectl apply -f examples/dds-test/yaml
