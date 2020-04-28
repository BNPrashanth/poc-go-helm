### POC-HELM
A poc project to test how helm charts work in golang project with configmap

## To run the docker image with helm
- Install helm
- Start a kubernetes cluster (example minikube start)
- To deploy the helm charts,
- helm install poc-helm ./.helm

## To check if it started properly
- kubectl get all
The above command should show a new deployment, pod and a service
