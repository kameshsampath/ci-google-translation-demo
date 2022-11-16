# gRPC Lingua Greeter

A demo on how to use Google API with gRPC application with go based server and gRPC-web client

## Pre-requisites

- Docker for Desktop
- [k3d](https://k3d.io)
- Drone CLI
- node
- Google API Key to Access Translation API
- Node (if you want to try client)

## Create Cluster and Registry

```shell
k3d registry create myregistry.localhost --port 5001
k3d cluster create --name=lingua-greeter-demos --registry-use k3d-myregistry.localhost:5001
```

## Build Application

```shell
drone exec --trusted --env-file=.env
```

## Deploy Application

```shell
kubectl apply -k k8s
```

## Build Client

```shell
drone exec --trusted --env-file=.env --include="protoc-web"
```

### Build JS

```shell
cd client/web
pnpm install
pnpm build 
```

>**NOTE**: If you don't have pnpm you can use npm or yarn

Open the index.html in browser to view the Translated Greetings.