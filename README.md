# Provider Artifactory

`provider-artfactory` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Artifactory API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/myorg/provider-artfactory):
```
up ctp provider install myorg/provider-artfactory:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-artfactory
spec:
  package: myorg/provider-artfactory:v0.1.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/myorg/provider-artfactory).

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/myorg/provider-artfactory/issues).



# Testing

Cannot use OSS Artifactory because it does not support creating repositories through REST APIs.
Get a license for Artifactory: https://jfrog.com/start-free/#ft
Set the environment variable `ARTIFACTORY_LICENSE_KEY` in your local ~/.zshrc and restart your IDE.

in a terminal in the dev container:
```console
mage setupE2E
kubectl port-forward --namespace jfrog svc/artifactory-artifactory-nginx 8888:80
```

in new terminal:
```console
kubectl apply -f package/crds
make run
```

in new terminal:
```console
kubectl apply -f e2e/providerconfig.yaml
#kubectl apply -f examples/manifests/genericrepository.yaml
mage testE2E
```
