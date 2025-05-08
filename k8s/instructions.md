# Kubernetes Deployment Instructions

This document provides step-by-step instructions for deploying the Guest Check microservice on a Kubernetes cluster.

## Prerequisites

- A running Kubernetes cluster (local or cloud)
- [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl) installed and configured
- [Helm](https://helm.sh/) (optional, for advanced deployments)

## Create Secrets

1. Fill in the `k8s/.env` file using `k8s/.env.example` as a template.
2. Run the following command to create a Kubernetes secret from the `.env` file:

   ```sh
   kubectl create secret generic guest-check-secret --from-env-file k8s/.env
   ```

## Create a Secret for Docker Registry

If your deployment requires pulling images from a private Docker registry, create a registry secret:

```sh
kubectl create secret docker-registry regsecret \
  --docker-server=$DOCKER_REGISTRY_SERVER \
  --docker-username=$DOCKER_USER \
  --docker-password=$DOCKER_PASSWORD \
  --docker-email=$DOCKER_EMAIL
```

Where:
- `$DOCKER_REGISTRY_SERVER`: URL for the registry
- `$DOCKER_USER`: Registry username
- `$DOCKER_PASSWORD`: Registry password
- `$DOCKER_EMAIL`: Optional, any email

## Deploy All Resources

Run the following command to deploy all Kubernetes resources:

```sh
kubectl apply -f ./k8s
```

## Troubleshooting

- **Secret Issues**: Ensure your `.env` file is correctly formatted and all required variables are set.
- **Registry Access**: Verify that your Docker registry credentials are correct and the registry is accessible from your cluster.
- **Resource Limits**: Check if your cluster has sufficient resources (CPU, memory) to run the deployment.

## Cleanup

To remove all deployed resources, run:

```sh
kubectl delete -f ./k8s
```

---

For more detailed information, refer to the [main README.md](../README.md).
