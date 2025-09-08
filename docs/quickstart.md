# Quick Start

This guide walks through the process to install the OCI official provider-family.

To use OCI official provider-family, install Crossplane into your Kubernetes cluster, install the `Provider`, apply a `ProviderConfig`, and create a managed resource in OCI via Kubernetes.

## Install the official OCI provider-family
Install the official provider-family into the Kubernetes cluster with a Kubernetes configuration file. For instance, let's install the `provider-oci-objectstorage`

Note: The first provider installed of a family also installs an extra provider-family Provider. The provider-family provider manages the `ProviderConfig` for all other providers in the same family.

```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-oci-family
spec:
  package: ghcr.io/oracle-samples/provider-family-oci:v0.0.1-alpha.1-amd64
---
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-oci-objectstorage
spec:
  package: ghcr.io/oracle-samples/provider-oci-objectstorage:v0.0.1-alpha.1-amd64
EOF
```

Apply this configuration with `kubectl apply -f`.

After installing the provider, verify the install with `kubectl get providers`.

```
$ kubectl get providers         

NAME                         INSTALLED   HEALTHY   PACKAGE                                                                  AGE
provider-oci-family          True        True      ghcr.io/oracle-samples/provider-family-oci:v0.0.1-alpha.1-amd64          3m3s
provider-oci-objectstorage   True        True      ghcr.io/oracle-samples/provider-oci-objectstorage:v0.0.1-alpha.1-amd64   3m2s
```

It may take up to 5 minutes to report `HEALTHY`.

## Create a managed resource

Create a managed resource to verify the provider-oci-objectstorage is functioning.

Use this command to instruct Crossplane to create the bucket in the OCI tenancy.

```shell
# Edit examples/objectstorage/bucket.yaml with your compartment and storage name space as documented.
# Apply the example that creates an Object Storage bucket

$ kubectl apply -f examples/objectstorage/bucket.yaml
```

Verify the status of the resource by running this command (example output is shown).
```shell
$ kubectl get managed

NAME                                                                             READY   SYNCED   EXTERNAL-NAME                                 AGE
bucket.objectstorage.oci.upbound.io/bucket-via-crossplane4   True    True     n/idimd1fghobk/b/bucket-via-crossplane   10m
```
Upbound created the bucket when the values `READY` and `SYNCED` are True. This may take up to 5 minutes.

If the `READY` or `SYNCED` are blank or False use kubectl describe to understand why.

Here is an example of a failure because the `spec.providerConfigRef.name` value in the `Bucket` doesn't match the `ProviderConfig` `metadata.name`.

```
$kubectl describe bucket

Name:         bucket-via-crossplane4
Namespace:    
Labels:       provider=oci
Annotations:  <none>
API Version:  objectstorage.oci.upbound.io/v1alpha1
Kind:         Bucket
# Output truncated
Spec:
  Deletion Policy:  Delete
  For Provider:
    Compartment Id:  ocid1.compartment.oc1..xxx
    Name:            bucket-via-crossplane
    Namespace:       id8pypxcqtqs
  Management Policies:
    *
  Provider Config Ref:
    Name:  default
Status:
  At Provider:
  Conditions:
    Last Transition Time:  2025-09-05T20:07:20Z
    Message:               connect failed: cannot get terraform setup: cannot get referenced ProviderConfig: ProviderConfig.oci.upbound.io "default" not found
    Observed Generation:   1
    Reason:                ReconcileError
    Status:                False
    Type:                  Synced
Events:
  Type     Reason                   Age               From                                                        Message
  ----     ------                   ----              ----                                                        -------
  Warning  CannotConnectToProvider  3s (x4 over 10s)  managed/objectstorage.oci.upbound.io/v1alpha1, kind=bucket  cannot get terraform setup: cannot get referenced ProviderConfig: ProviderConfig.oci.upbound.io "default" not found
```

The output indicates the `Bucket` is using a `ProviderConfig` named `default`. The applied `ProviderConfig` is `oci-provider`.

```
$ kubectl get providerconfig
NAME           AGE
oci-provider   7s
```

## Delete the managed resource

Remove the managed resource by using `kubectl delete -f examples/objectstorage/bucket.yaml`. Verify removal of the bucket with kubectl get buckets

```
$ kubectl delete -f examples/objectstorage/bucket.yaml

bucket.objectstorage.oci.upbound.io "bucket-via-crossplane4" deleted
```

```
$ kubectl get buckets

No resources found
```