## Overview

Firebase currently supports many forms of authentication outside of google, but not all of them.

This project provides an API written in GO that exposes an endpoint which receives a externally provided token and creates a firebase custom token based on that.

For it to work, a firebase project needs to be created and a service account file named "service-acc.json" needs to be generated from the project options and placed in the root folder.

Since this project generates a JWT token based on a string value, it's assumed that this string value will also be a JWT token, so /api/token requires a Bearer token in its header. This can be easily changed in handler.go.

## Endpoints

```
GET /api/token
Generates custom token
```
```
GET /ping
Health
```

## To call
```
curl -H 'Authorizationy: Bearer _your-api-key_' -X GET http://your-url/api/token
```

## Docker and Appengine
- Dockerfile generates a docker image.
- app.yaml pushes the project to appengine
