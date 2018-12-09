## Overview

Firebase currently supports many forms of authentication outside of google, but not all of them.

This project provides an API written in GO that exposes an endpoint which receives a externally provided token and creates a firebase custom token based on that.

For it to work, a firebase project needs to be created and a service account file named "service-acc.json" needs to be generated from the project options and placed in the root folder.

This url currently requires an API Key to work, this can be easily changed by modifying the handler.go file.

## Endpoints

```
POST /api/token
Generates custom token
```
```
GET /ping
Health
```

## To call
```
curl -H 'X-API-Key: your-api-key' -d '{"token":"token_value"}' -X POST http://your-url/api/token
```

## Docker and Appengine
- Dockerfile generates a docker image.
- app.yaml pushes the project to appengine
