# Mono Repository for all the microservices

## Set Up with Kubernetes
- `kubectl`, `make` and `minikube` should already be installed;
- make sure the `docker daemon`, `minikube` are running and `ingress`, `ingress-dns` addons are available;
```
    sudo systemctl start docker
    minikube start
    minikube addons enable ingress
    minikube addons enable ingress-dns
```
- `80` port of the host machine should not be in use; 
- add hosts to the `/etc/hosts`:
```
    127.0.0.1       ui
    127.0.0.1       gateway
```
- run the following command from the **project root**:
```
    make k8s_start
```
- the app will be available under the:
```
    http://ui
    http://gateway
```

## Set Up with Docker
- `docker engine`, `docker compose` and `make` should already be installed, `docker engine` should run;
- `8080` and `8888` ports of the host machine should not be in use; 
- create `.env` file using `.env.dist` as template;
- run the following command from the **project root**:
```
    make docker_start
```
- to stop containers run the following command from the **project root**:
```
    make docker_stop
```
- the app will be available under the following origins:
```
    http://localhost:8080 - ui
    http://localhost:8888 - gateway
```