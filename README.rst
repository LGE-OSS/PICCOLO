PICCOLO
=======

| Piccolo is a project providing MCO (Mixed Critical Orchestration) for
  SDV(Software Defined Vehicle).
| PICCOLO is a project specifically designed to consider the unique
  characteristics of vehicles in the vehicle environment and provide
  orchestration technology based on containerization for vehicle
  scenarios.

**[Korean Version]**

| PICCOLO는 자동차를 위한 Mixed Critical Orchestration을 제공하기 위한
  프로젝트이다.
| PICCOLO는 차량 환경에서 차량만의 특성을 고려한 차량의 시나리오를
  컨테이너 기반으로 운용하는 오케스트레이션 기술을 제공하기 위한
  프로젝트이다.

Description
-----------

| PICCOLO configures the operating environment required for
  applications/services installed in vehicles, restricts deployment
  locations, and supports operation on an application/service level.
| For Vehicle Services/Applications composed of a collection of Micro
  Services, PICCOLO operates as follows:

1. During the packaging phase, it supports the configuration of Vehicle
   Service/Application units.

2. During the deployment phase, it sets up the operating environment on
   a Micro Service level.

3. During the operation phase, it manages and operates on a Vehicle
   Service/Application level.

4. During the update phase, it updates on a Micro Service level.

PICCOLO provides the `PICCOLO Specification <docs/manifests/>`__ for
configuring the runtime environment, defining deployment policies,
managing device access permissions, and setting up network
configurations for applications installed in vehicles. The development
of PICCOLO follows the following steps to provide Dynamic Mixed Critical
Workloads Orchestration for SDV(Software Defined Vehicle):

1. Define manifests/specifications for the management elements and
   operational methods required for Vehicle Service Orchestration in
   SDV(Software Defined Vehicle).

2. Utilize existing container ecosystems such as K8S, K3S, Docker Swarm,
   Docker, Containerd, Podman for quick and easy validation.

3. Develop core orchestrator functionality (including API-Server and
   Scheduler) based on an Embedded Edge Platform (TBD).

4. Develop the Vehicle Orchestrator Platform (TBD).

**[Korean Version]**

PICCOLO는 차량에 탑재되는 Application / Service에 필요한 구동 환경을
구성하고 배포 위치를 제한하며 Application / Service 단위 운용을
지원한다.

Micro Service 단위의 집합으로 구성된 Vehicle Service / Application에
대하여 PICCOLO는 아래와 같이 동작한다.

1. 패키지 단계에서는 Vehicle Service / Application 단위 구성을 지원한다.

2. 배포 단계에서는 Micro Service 단위로 구동 환경을 구성한다.

3. 운용 단계에서는 Vehicle Service / Application 단위로 관리 운용한다.

4. 업데이트 단계에서는 Micro Service 단위로 업데이트한다.

PICCOLO는 차량에 탑재되는 Application을 위한 구동 환경 설정, 운용 및
배포 정책, 디바이스 접근 권한, 네트워크 설정 등을 위한 `PICCOLO
Specification <docs/manifests/>`__\ 을 제공한다. PICCOLO는 SDV를 위한
Dynamic Mixed Critical Workloads Orchestration을 제공하기 위한
프로젝트로 아래와 같은 순서에 따라서 개발을 진행한다.

1. SDV를 위한 Vehicle Service Orchestration에 필요한 관리 요소 및 운용
   방법에 대한 Manifests / Specification 정의

2. K8S, K3S, Docker Swarm, Docker, Containerd, Podman 등 기존 Container
   Ecosystem을 활용하여 빠르고 쉽게 검증

3. Embedded Edge Platform을 위한 Core Orchestrator(API-Server, Scheduler
   포함) 기반으로 기능 개발 (TBD)

4. Vehicle Orchestrator Platform 개발 (TBD)

Installation
------------

See `Installation <docs/installation.rst>`__.

Quick Start
-----------

See `Quick Start <docs/quick-start.rst>`__.

Contributing
------------

How it makes
~~~~~~~~~~~~

| This folder structure was created by
  `kubebuilder <https://github.com/kubernetes-sigs/kubebuilder>`__.
| However, you don't need ``kubebuilder`` to work.

.. code:: 

   kubebuilder init --domain piccolo.org --repo crd.piccolo.org/piccolo
   kubebuilder create api --group crd --version v1alpha1 --kind VehPod --resource --controller
   kubebuilder create api --group crd --version v1alpha1 --kind DevRes --resource --controller=false
   kubebuilder create api --group crd --version v1alpha1 --kind VehContainerConfig --resource --controller=false

======== ======== ==========
CRD name resource controller
======== ======== ==========
Pod      Y        Y
Device   Y        **N**
Function Y        **N**
======== ======== ==========

How it works
~~~~~~~~~~~~

PICCOLO 프로젝트는 현재 Kubernetes 혹은 K3S 기반의 `Operator
Pattern <https://kubernetes.io/docs/concepts/extend-kubernetes/operator/>`__\ 을
따른다.
(`Controllers <https://kubernetes.io/docs/concepts/architecture/controller/>`__
참조)

This project aims to follow the Kubernetes `Operator
pattern <https://kubernetes.io/docs/concepts/extend-kubernetes/operator/>`__.

| It uses
  `Controllers <https://kubernetes.io/docs/concepts/architecture/controller/>`__,
| which provide a reconcile function responsible for synchronizing
  resources until the desired state is reached on the cluster.

The initial version of Piccolo is based on
`Kubebuilder <https://book.kubebuilder.io/>`__. So you can get the
information you need for Piccolo development from kubebuilder's
documentation.

Piccolo also used `Helmify <https://github.com/arttor/helmify>`__ to
create helm chart.

Test It Out
~~~~~~~~~~~

1. Install the CRDs into the cluster:

.. code:: 

   make install

1. Run your controller (this will run in the foreground, so switch to a
   new terminal if you want to leave it running):

.. code:: 

   make run

1. (optional) The controller can apply by pod.

.. code:: 

   make docker-build docker-push IMG=<some-registry>/piccolo:tag

Modifying the API definitions
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

If you are editing the API definitions, generate the manifests such as
CRs or CRDs using:

.. code:: bash

   make manifests

**NOTE:** Run ``make help`` for more information on all potential
``make`` targets

More information can be found via the `Kubebuilder
Documentation <https://book.kubebuilder.io/introduction.html>`__
