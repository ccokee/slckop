# SLCK Operator

SLCK Operator is a Kubernetes Operator designed to provide a secure execution environment for running commands using the SLCK system on your Kubernetes cluster. It enables you to deploy and manage SLCK instances with ease, allowing users to interact with command-line applications using frontend components for React, Vue.js, or Angular. SLCK Operator follows the Kubernetes Operator pattern and leverages custom resources to define and manage SLCK instances.

## Description

SLCK Operator automates the deployment, scaling, and management of SLCK instances on your Kubernetes cluster. It allows you to easily expose binaries for remote execution using WebSocket connections and provides a browser-based terminal for running commands securely. The operator is built using the Kubebuilder framework and includes custom resources, controllers, and reconciliation logic to ensure the desired state of your SLCK instances is maintained.

## Getting Started

You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing or run against a remote cluster. **Note:** Your controller will automatically use the current context in your kubeconfig file (i.e., whatever cluster `kubectl cluster-info` shows).

### Running on the cluster

1. Install Instances of Custom Resources:

```sh
kubectl apply -f config/samples/
```

2. Build and push your image to the location specified by `IMG`:

```sh
make docker-build docker-push IMG=<some-registry>/slck-operator:tag
```

3. Deploy the controller to the cluster with the image specified by `IMG`:

```sh
make deploy IMG=<some-registry>/slck-operator:tag
```

### Uninstall CRDs

To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller

UnDeploy the controller from the cluster:

```sh
make undeploy
```

## Contributing

We welcome contributions to the SLCK Operator project! If you'd like to contribute, please follow the standard GitHub workflow: fork the repository, create a new branch, make your changes, and submit a pull request. Make sure to follow the existing code style and provide clear, concise commit messages.

### How it works

This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/).

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/),
which provide a reconcile function responsible for synchronizing resources until the desired state is reached on the cluster.

### Test It Out

1. Install the CRDs into the cluster:

```sh
make install
```

2. Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

```sh
make run
```

**NOTE:** You can also run this in one step by running: `make install run`

### Modifying the API definitions

If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2023.

Licensed under the MIT License (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://opensource.org/licenses/MIT

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.