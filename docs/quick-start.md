# Quick Start

You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.  
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

## Prerequisites
Same as prerequisities of kubebuilder.
- go 1.20.0+
- docker 17.03+.
- kubectl 1.11.3+.
- Access to a Kubernetes 1.11.3+ cluster.  

## Running Examples
1. [Camera test](/examples/camera-test/README.md) - For this example, a camera connection is also required.
2. [DDS test](/examples/dds-test/README.md)
