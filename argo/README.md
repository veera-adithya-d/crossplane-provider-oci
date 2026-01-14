# How to Setup and Test Argo Workflow

## Step 1: Install Argo Workflow

Install Argo workflow using the following command:

<!-- TODO: Update instructions to install argo -->
```bash
kubectl create namespace argo
kubectl apply -n argo -f https://github.com/argoproj/argo-workflows/releases/download/v3.7.6/quick-start-minimal.yaml
```

## Step 2: Port-forward Argo Server

Port-forward the Argo server using the following command:

```bash
kubectl -n argo port-forward deployment/argo-server 2746:2746
```

## Step 3: Clone the Repository

Clone the repository using the following command:
<!-- Replace with oracle repo-->
```bash
git clone https://github.com/endurthiabhilash/crossplane-provider-oci.git -b argo-workflow
```

## Step 4: Create Argo Workflow Service Account

Create the Argo workflow service account by running:

```bash
kubectl apply -f argo/setup/workflow-serviceaccount.yaml
```

## Step 5: Create PVC for Workflow

Create a PVC to store the cloned repository in the workflow by running:

```bash
kubectl apply -f argo/setup/git-repo-pvc.yaml
```

## Step 6: Create Cluster Admin Role Binding

Create the cluster admin role binding by running:

```bash
kubectl apply -f argo/setup/workflow-admin-clusterrolebinding.yaml
```

## Step 7: Submit the Workflow

To start the workflow, run:

```bash
argo submit argo/workflows/test-workflow.yaml \
-p clone_repo=true \
-p delete_crossplane=true \
-p setup_crossplane=true \
-p setup_uptest=false \
-p create_compartment=true \
-p git_ref=argo-workflow \
-p git_repo=https://github.com/endurthiabhilash/crossplane-provider-oci.git \
-p provider-image-repo-name=iad.ocir.io/iddevjmhjw0n/kousalyak \
-p family-provider-version=v0.0.1-alpha \
-p tenancy=ocid1.tenancy.oc1.xxx \
-p compartment_id=ocid1.compartment.oc1.xxx \
-p image_id=ocid1.image.oc1.xxx
```

## Notes

- Make sure to replace the placeholders (`ocid1.tenancy.oc1.xxx`, `ocid1.compartment.oc1.xxx`, `ocid1.image.oc1.xxx`) with your actual OCI tenancy, compartment, and image IDs.
- Update the `provider-image-repo-name` and `family-provider-version` parameters according to your setup.

## Available Workflow Templates

The `argo/workflows/templates` directory contains reusable workflow templates that can be referenced in your workflows. These templates provide common functionality for tasks such as cloning repositories, setting up Crossplane, and creating resources.

### Available Templates

1. **clone-repo-template**: Clones a Git repository into a Persistent Volume Claim (PVC).
   - Parameters:
     - `git_repo`: URL of the Git repository to clone.
     - `git_ref`: Branch or commit to check out.

2. **crossplane-template**: Sets up Crossplane with OCI providers.
   - Parameters:
     - `namespace`: Namespace to install Crossplane.
     - `region`: OCI region.
     - `providers`: Comma-separated list of Crossplane providers to install.
     - `provider-image-repo-name`: Image repository for OCI provider.
     - `family-provider-version`: Version of the OCI provider family.
     - `tenancy`: OCI tenancy OCID.

3. **test-template**: Creates resources from YAML files.
   - Parameters:
     - `resourceFile`: Path to the YAML file defining the resource.
     - `envVars`: Environment variables to substitute in the YAML file.

## Using Templates in Workflows

To use these templates in your workflows, first apply them to your Argo instance using:

```bash
kubectl apply -f argo/workflows/templates/clone-repo.yaml
kubectl apply -f argo/workflows/templates/crossplane.yaml
kubectl apply -f argo/workflows/templates/test.yaml
```

Then, reference them in your workflow definition using `templateRef`. For example:

```yaml
- name: clone-repo
  templateRef:
    name: clone-repo-template
    template: clone-repo
  arguments:
    parameters:
      - name: git_repo
        value: "https://github.com/your/repo.git"
      - name: git_ref
        value: "main"
```
