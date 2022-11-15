# gRPC Lingua Greeter

A demo on how to use Google API with gRPC application with go based server and gRPC-web client

## Pre-requisites

- Docker for Desktop
- Drone CLI
- node
- Google API Key to Access Translation API

## Generate Server and Client Stubs

```shell
drone exec --include="protoc-server" --include="protoc-web"
```

### Add 