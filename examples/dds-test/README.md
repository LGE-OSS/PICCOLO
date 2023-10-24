# DDS test
This example makes two application pod using the DDS communication.  
The key point is that the application developer does not need to know any detailed information about environment variables for DDS configuration.  

![Diagram](/examples/dds-test/img/diagram-dds-test.png)

## Running on the cluster

### Build the dds test application:
```
cd examples/dds-test
docker build -t <some-registry>/<APP_IMG_NAME>:tag -f Dockerfile_<ping or pong>
docker push <some-registry>/<APP_IMG_NAME>:tag
```

### Build and push your controller image:
```
make docker-build docker-push IMG=<some-registry>/piccolo:tag
```
**Note:** Default value of `IMG` is `controller:latest`

### Deploy CRDs and the controller to the cluster:
```
make install
make deploy IMG=<some-registry>/piccolo:tag
```
Or you can use dry-run
```
make dry-run IMG=<some-registry>/piccolo:tag
kubectl apply -f dry-run/
```

### Modify your test application image name:
In `examples/dds-test/yaml/cr_vehpod_ddstest_ping.yaml`,
```
  podSpec:
    containers:
    - name: ping-container
      image: <some-registry>/<APP_IMG_NAME>:tag
```
In `examples/dds-test/yaml/cr_vehpod_ddstest_pong.yaml`,
```
  podSpec:
    containers:
    - name: pong-container
      image: <some-registry>/<APP_IMG_NAME>:tag
```

### Add node label or delete nodeSelector in yaml
In two VehPod yaml,
```
    nodeSelector:
      name: "master"
```
You need to add label or delete nodeSelector in yaml.

![Screenshot - node label](/examples/camera-test/img/label.png)

### Modify your dds setting
In `examples/dds-test/yaml/cr_vehcontainerconfig_dds.yaml`,
```
spec:
  piccoloEnv:
    - name: CYCLONEDDS_URI
      value: '<CycloneDDS><Domain Id="100"></Domain></CycloneDDS>'
```
**Note:** domain id value can be changed.

### Install instances of Custom Resources:
```
kubectl apply -f examples/dds-test/yaml/
```

### Check the result:
```
kubectl get po -A
```
You can see the result like below.

![Screenshot - kubectl get po](/examples/dds-test/img/pod-by-controller.png)

```
kubectl logs -f <POD_NAME> --namespace=<APP_NAMESPACE>
```
You can see the result like below.

![Screenshot - kubectl get po](/examples/dds-test/img/domain-id-log.png)

### Undeploy CRDs and controller:
UnDeploy the controller from the cluster:
```sh
make undeploy
make uninstall
```
If you use dry-run command,
```
kubectl delete -f dry-run/
```
