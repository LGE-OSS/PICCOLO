Piccolo for EWAOL
=================

| `README <README.rst>`__ is a document for x86 architecture.
| Hence, the guide cannot apply in EWAOL architecture.

Test it out
-----------

| EWAOL does not contain commands like ``git, apt-get, make``.
| Therefore, after configuring it through Raspberry Pi, you use the
  method of copying it through ``scp``.

Configuration in Raspberry Pi
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

1. Build and push controller container image:

.. code:: 

   make docker-build docker-push IMG=<some-registry>/piccolo:tag

1. Configure ``yaml`` file:

.. code:: 

   make dry-run IMG=<some-registry>/piccolo:tag

1. Copy ``piccolo`` folder to EWAOL:

.. code:: 

   scp -r piccolo/ root@IP:/($destination)

Running on EWAOL
~~~~~~~~~~~~~~~~

1. Now, we can see following structure in EWAOL folder.

.. code:: 

   root@ava:~/test/piccolo# ls -al
   total 156
   drwxr-xr-x 12 root root  4096 Oct 18 08:48 .
   drwxr-xr-x  3 root root  4096 Oct 18 08:47 ..
   -rw-r--r--  1 root root   129 Oct 18 08:47 .dockerignore
   drwxr-xr-x  8 root root  4096 Oct 18 08:47 .git
   -rw-r--r--  1 root root   393 Oct 18 08:47 .gitignore
   -rw-r--r--  1 root root  1489 Oct 18 08:47 CONTRIBUTING.md
   -rw-r--r--  1 root root  1278 Oct 18 08:47 Dockerfile
   -rw-r--r--  1 root root  7922 Oct 18 08:47 Makefile
   -rw-r--r--  1 root root   860 Oct 18 08:47 PROJECT
   -rw-r--r--  1 root root  3834 Oct 18 08:47 README.md
   drwxr-xr-x  3 root root  4096 Oct 18 08:47 api
   drwxr-xr-x  3 root root  4096 Oct 18 08:47 bin
   drwxr-xr-x  2 root root  4096 Oct 18 08:48 cmd
   drwxr-xr-x  8 root root  4096 Oct 18 08:47 config
   -rw-r--r--  1 root root  2221 Oct 18 08:47 cover.out
   drwxr-xr-x  3 root root  4096 Oct 18 08:48 docs
   drwxr-xr-x  2 root root  4096 Oct 18 08:47 dry-run
   drwxr-xr-x  4 root root  4096 Oct 18 08:48 examples
   -rw-r--r--  1 root root  2983 Oct 18 08:48 go.mod
   -rw-r--r--  1 root root 60178 Oct 18 08:48 go.sum
   -rw-r--r--  1 root root    23 Oct 18 08:48 go.work
   drwxr-xr-x  2 root root  4096 Oct 18 08:47 hack
   drwxr-xr-x  3 root root  4096 Oct 18 08:47 internal
   -rw-r--r--  1 root root  1563 Oct 18 08:47 license.rst

1. At first, apply controller and CRDs.

.. code:: 

   root@ava:~/test/piccolo# kubectl apply -f dry-run/
   namespace/piccolo-system created
   customresourcedefinition.apiextensions.k8s.io/devres.crd.piccolo.org created
   customresourcedefinition.apiextensions.k8s.io/vehcontainerconfigs.crd.piccolo.org created
   customresourcedefinition.apiextensions.k8s.io/vehpods.crd.piccolo.org created
   serviceaccount/piccolo-controller-manager created
   role.rbac.authorization.k8s.io/piccolo-leader-election-role created
   clusterrole.rbac.authorization.k8s.io/piccolo-manager-role created
   clusterrole.rbac.authorization.k8s.io/piccolo-metrics-reader created
   clusterrole.rbac.authorization.k8s.io/piccolo-proxy-role created
   rolebinding.rbac.authorization.k8s.io/piccolo-leader-election-rolebinding created
   clusterrolebinding.rbac.authorization.k8s.io/piccolo-manager-rolebinding created
   clusterrolebinding.rbac.authorization.k8s.io/piccolo-proxy-rolebinding created
   service/piccolo-controller-manager-metrics-service created
   deployment.apps/piccolo-controller-manager created

1. We can check CRDs and controller

.. code:: 

   root@ava:~/test/piccolo# kubectl get crd
   NAME                                    CREATED AT
   devres.crd.piccolo.org                  2023-10-23T03:05:29Z
   vehcontainerconfigs.crd.piccolo.org     2023-10-23T03:05:29Z
   vehpods.crd.piccolo.org                 2023-10-23T03:05:29Z

   root@ava:~/test/piccolo# kubectl get po -A | grep controller
   piccolo-system  piccolo-controller-manager-6d49f94594-c9j2z 2/2  Running  0    91s

1. | We have two examples, but only ``dds`` test is available in EWAOL.
   | See `DDS test </examples/dds-test/README.rst>`__ for details.

.. code:: 

   root@ava:~/test/piccolo/examples/dds-test# kubectl apply -f yaml/
   vehcontainerconfig.crd.piccolo.org/dds-domain-id created
   vehpod.crd.piccolo.org/ddstest-ping created
   vehpod.crd.piccolo.org/ddstest-pong created
   namespace/system-application created
   namespace/3rd-party-application created
   limitrange/limit-range created
   resourcequota/object-count-quota created

   root@ava:~/test/piccolo/examples/dds-test# kubectl get po -n system-application
   NAME           READY   STATUS    RESTARTS   AGE
   ddstest-ping   1/1     Running   0          20s
   ddstest-pong   1/1     Running   0          20s

1. Check the logs.

.. code:: 

   root@ava:~/test/piccolo/examples/dds-test# kubectl logs -n system-application ddstest-ping
   === [Publisher] begin
   === [Publisher] Waiting for subscriber. [domain_id] 5
   === [Publisher] success finding subscriber.
   receive, i : 1, success : 1
   450996
   receive, i : 2, success : 2
   386836
   receive, i : 3, success : 3
   269157
   ......

   root@ava:~/test/piccolo/examples/dds-test# kubectl logs -n system-application ddstest-pong
   === [Subscriber] Waiting for publisher. [doamin_id] 5
   === [Subscriber] success finding publisher.

License
-------

Unless otherwise specified, all content, including all source code files
and documentation files in this repository are:

Copyright (c) 2023 LG Electronics, Inc.

Licensed under the Apache License, Version 2.0 (the "License"); you may
not use this file except in compliance with the License. You may obtain
a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

SPDX-License-Identifier: Apache-2.0
