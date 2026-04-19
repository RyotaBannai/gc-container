# gc-container

## Deploy To GKE

This repository builds and publishes the image below to Artifact Registry.

`asia-northeast1-docker.pkg.dev/singular-rope-268807/my-repo/my-go-app:good`

### 1. Authenticate kubectl to your cluster

Replace the placeholders with your cluster name and zone or region.

```bash
gcloud container clusters get-credentials CLUSTER_NAME --zone CLUSTER_ZONE --project singular-rope-268807
```

If the cluster is regional, use `--region` instead of `--zone`.

### 2. Ensure GKE can pull from Artifact Registry

For a cluster in the same project, the node service account usually needs the Artifact Registry Reader role.

```bash
gcloud projects add-iam-policy-binding singular-rope-268807 \
	--member="serviceAccount:NODE_SERVICE_ACCOUNT" \
	--role="roles/artifactregistry.reader"
```

You can confirm the node service account with:

```bash
gcloud container clusters describe CLUSTER_NAME \
	--zone CLUSTER_ZONE \
	--project singular-rope-268807 \
	--format="value(nodeConfig.serviceAccount)"
```

### 3. Deploy the workload

```bash
kubectl apply -f k8s-sample.yaml
```

### 4. Check rollout status

```bash
kubectl rollout status deployment/gc-container
kubectl get pods
kubectl get svc gc-container
```

When the external IP is assigned, access the app with:

```bash
kubectl get svc gc-container
curl http://EXTERNAL_IP/
```

### 5. Update the image later

If you push a new image tag, update the Deployment image and roll it out.

```bash
kubectl set image deployment/gc-container \
	gc-container=asia-northeast1-docker.pkg.dev/singular-rope-268807/my-repo/my-go-app:good
```
