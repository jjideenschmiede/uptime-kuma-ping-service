# Uptime Kuma Ping Service

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/jjideenschmiede/uptime-kuma-push-service.svg)](https://golang.org/) [![Go](https://github.com/jjideenschmiede/uptime-kuma-ping-service/actions/workflows/go.yml/badge.svg)](https://github.com/jjideenschmiede/uptime-kuma-ping-service/actions/workflows/go.yml) [![Docker Image CI](https://github.com/jjideenschmiede/uptime-kuma-ping-service/actions/workflows/docker-image.yml/badge.svg)](https://github.com/jjideenschmiede/uptime-kuma-ping-service/actions/workflows/docker-image.yml) [![Docker Hub](https://img.shields.io/docker/pulls/jjdevelopment/uptime-kuma-ping-service.svg)](https://hub.docker.com/r/jjdevelopment/uptime-kuma-ping-service)

This Docker is to get an ping from an [Uptime Kuma](https://github.com/louislam/uptime-kuma) service. Here you will find a little introduction on how to use it.

The whole thing is brought to run in a Docker container. For this, a few variables from the Dockerfile are needed, if they are not stored, then they have a default value stored. You can find the variables here.

| Variable              | Default value |
|-----------------------|:-------------:|
| ENABLE_SSL            | false         |
| CERTIFICATE_CRT_NAME  | default       |
| CERTIFICATE_KEY_NAME  | default       |

If the ENABLE_SSL is set to true, then the certificate names must be stored. These must be stored in the volume mapping. Below are the instructions for SSL.

## Launch Docker Container

To start the container properly, here is a small template. A volume mapping is only required if the application is to be operated with SSL.

### With SSL

```console
docker run -d --restart always --name uptime-kuma-ping-service -e ENABLE_SSL='true' -e CERTIFICATE_CRT_NAME='...fullchain.pem' -e CERTIFICATE_KEY_NAME='...key.pem' -v /var/lib/certificates:/go/src/app/files/certificates jjdevelopment/uptime-kuma-ping-service:latest
```

### Without SSL

```console
docker run -d --restart always --name uptime-kuma-ping-service jjdevelopment/uptime-kuma-ping-service:latest
```

Now the container can be started, so that your Uptime Kuma can connect to your server and check the service directly. So you know directly if Docker is still running on the server.

Click [here](https://hub.docker.com/r/jjdevelopment/uptime-kuma-ping-service) to go directly to the Docker HUB.

## Contribute

If you want to help with development, or have found a bug, open a [new issue](https://github.com/jjideenschmiede/uptime-kuma-ping-service/issues).
