# SessionMapper ![GitHub branch checks state](https://img.shields.io/github/checks-status/carnage-sh/sessionmapper/main?color=deeppink)

**SessionMapper** adds properties to your requests based on custom logic.
To perform its task, it connects to a Mapping Server that matches sessions to
users, groups, roles, experiments, preferences, tracing properties or whatever
is needed to provide advanced server-side routing capabilities.

## Overview

The SessionMapper middleware for Traefik queries an API and set additional
headers before it runs upstream requests. The middleware can be used in front
of others. It allows advanced behaviors and applications are limitless... The
principle is shown below:

![overview](./img/architecture.png)

SessionMapper can easily be used and adapted for A/B testing, canary
deployment, role-based security, advanced tracing, test collection,
personalization...

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
[carnage-sh/sessionmapper](http://github.com/carnage-sh/sessionmapper) provides
a simple service that implements the current request/response protocol. Next
release will improve the protocol to reduce the latency and make it more
reliable, including blocking on failure. Do not hesitate to open an issue if
you repository find this plugin useful and want to support more advanced
scenarios.
