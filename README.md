**SessionMapper** is a Traefik plugin. It provides an advanced middleware to
easily map sessions with users, groups or any properties. It enables A/B
testing, canary deployments, security, data collection and more...

![GitHub branch checks state](https://img.shields.io/github/checks-status/blaqkube/sessionmapper/main?color=deeppink)
![GitHub](https://img.shields.io/github/license/blaqkube/sessionmapper?color=lime)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/blaqkube/sessionmapper/main?color=blue&label=Go)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/blaqkube/sessionmapper?color=black)

## Overview

The SessionMapper middleware for Traefik queries an API and set additional
headers before it runs upstream requests. The middleware can be used in front
of others. It allows advanced behaviors like setting a headers for a session, a
user, a group of users and an experiment. Applications are endless... The
principle is shown below:

![overview](./img/architecture.png)

## Configuration

You only need a few parameters to configure the middleware:

- `headers` lists the headers to capture from the request and send to the
  server (default: `me`)
- `server` is the server URL (default: `http://localhost:7777/`)
- `timeout` is the delay in milliseconds before the request to the server
  fails and no headers are set to the upstream request.

## Implementing a SessionMapper Server

For now, the session Mapper is a simple HTTP server that should be
colocated to traefik, for instance as a sidecar. The middleware performs
a `POST` to the `server` URL. The return message would look like below:

```json
{"upstream": {
	"key1": "value1",
	"key2": "value2"
}}
```

The `server` directory that is part of the
[blaqkube/sessionmapper](http://github.com/blaqkube/sessionmapper) provides
a simple service that implements the current request/response protocol. Next
release will improve the protocol to reduce the latency and make it more
reliable, including blocking on failure. Do not hesitate to open an issue if
you repository find this plugin useful and want to support more advanced
scenarios.
