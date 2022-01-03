# MetaClean API Version 1

This module is a [Gin](https://github.com/gin-gonic/gin) server providing a simple RESTful API to remove metadata from images. This server can be run locally on a host with [Golang](https://go.dev/) installed, via the provided [Docker](https://docker.com) environment, or accessed via the web.

## Getting Started

There are a number of Bash scripts provided for consistent setup and testing. Before attempting to develop or run, please use the `setup.sh` script to ensure you have all dependencies installed.

```
./setup.sh
```

From within the `Dockerfile.dev` image, the script is available as `setup`:

```
setup
```

## Running Locally

For development, the provided `Dockerfile` image should be built and run with this directory volume mounted.

```
./build.sh &&\
./run.sh
```

This will run the server in release mode and allow for local usage of the API.

## Running For Development

For development, the provided `Dockerfile.dev` image should be built and run with this directory volume mounted.

```
./build-dev.sh &&\
./dev.sh
```

This will set up a valid environment for Go and allow for local file changes to be available on the Docker image. Once inside the container, running `go run .` will start the server. Refer to `go run . -help` for available runtime flags and usage guides. If you are unable to access the server, please try setting the host to `0.0.0.0` or run the image with `--network host`.

### TODOs:

- Provide API reference
- Update API reference
